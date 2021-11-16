// This code is available on the terms of the project LICENSE.md file,
// also available online at https://blueoakcouncil.org/license/1.0.0.

//go:build lgpl
// +build lgpl

package eth

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"

	dexeth "decred.org/dcrdex/dex/networks/eth"
	"decred.org/dcrdex/server/asset"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

type swapCoinType uint8

const (
	sctInit swapCoinType = iota
	sctRedeem
)

var _ asset.Coin = (*swapCoin)(nil)

type swapCoin struct {
	backend                    *Backend
	contractAddr, counterParty common.Address
	secret, secretHash         [32]byte
	value                      uint64
	gasPrice                   uint64
	txid                       string
	locktime                   int64
	sct                        swapCoinType
}

// newSwapCoin creates a new swapCoin that stores and retrieves info about a
// swap. It requires a coinID that is a txid type of the initial transaction
// initializing or redeeming the swap. A txid type and not a swap type is
// required because the contract will give us no useful information about the
// swap before it is mined. Having the initial transaction allows us to track
// it in the mempool. It also tells us all the data we need to confirm a tx
// will do what we expect if mined and satisfies contract constraints. These
// fields are verified the first time the Confirmations method is called, and
// an error is returned then if something is different than expected. As such,
// the swapCoin expects Confirmations to be called with confirmations
// available at least once before the swap be trusted for swap initializations.
func (backend *Backend) newSwapCoin(coinID []byte, sct swapCoinType) (*swapCoin, error) {
	switch sct {
	case sctInit, sctRedeem:
	default:
		return nil, fmt.Errorf("unknown swapCoin type: %d", sct)
	}

	cID, err := DecodeCoinID(coinID)
	if err != nil {
		return nil, err
	}
	txCoinID, ok := cID.(*TxCoinID)
	if !ok {
		return nil, errors.New("coin ID not a txid")
	}
	txid := txCoinID.TxID
	tx, _, err := backend.node.transaction(backend.rpcCtx, txid)
	if err != nil {
		if errors.Is(err, ethereum.NotFound) {
			return nil, asset.CoinNotFoundError
		}
		return nil, fmt.Errorf("unable to fetch transaction: %v", err)
	}

	txdata := tx.Data()
	// Transactions that call contract functions must have extra data to do
	// so.
	if len(txdata) == 0 {
		return nil, errors.New("tx calling contract function has no extra data")
	}

	var (
		counterParty       = new(common.Address)
		secret, secretHash [32]byte
		locktime           int64
	)

	switch sct {
	case sctInit:
		initiations, err := dexeth.ParseInitiateData(txdata)
		if err != nil {
			return nil, fmt.Errorf("unable to parse initiate data: %v", err)
		}
		if txCoinID.Index >= uint32(len(initiations)) {
			return nil, fmt.Errorf("tx %v does not have initiation with index %d",
				txCoinID.TxID, txCoinID.Index)
		}
		initiation := initiations[txCoinID.Index]
		counterParty = &initiation.Participant
		secretHash = initiation.SecretHash
		locktime = initiation.RefundTimestamp.Int64()
	case sctRedeem:
		secret, secretHash, err = dexeth.ParseRedeemData(txdata)
		if err != nil {
			return nil, fmt.Errorf("unable to parse redeem data: %v", err)
		}
	}

	contractAddr := tx.To()
	if *contractAddr != backend.contractAddr {
		return nil, errors.New("contract address is not supported")
	}

	// Gas price is not stored in the swap, and is used to determine if the
	// initialization transaction could take a long time to be mined. A
	// transaction with a very low gas price may need to be resent with a
	// higher price.
	gasPrice, err := ToGwei(tx.GasPrice())
	if err != nil {
		return nil, fmt.Errorf("unable to convert gas price: %v", err)
	}

	// Value is stored in the swap with the initialization transaction.
	value, err := ToGwei(tx.Value())
	if err != nil {
		return nil, fmt.Errorf("unable to convert value: %v", err)
	}

	// For redemptions, the transaction should move no value.
	if sct == sctRedeem && value != 0 {
		return nil, fmt.Errorf("expected swapCoin value of zero for redeem but got: %d", value)
	}

	return &swapCoin{
		backend:      backend,
		contractAddr: *contractAddr,
		secret:       secret,
		secretHash:   secretHash,
		value:        value,
		gasPrice:     gasPrice,
		txid:         hex.EncodeToString(txid[:]),
		counterParty: *counterParty,
		locktime:     locktime,
		sct:          sct,
	}, nil
}

// validateRedeem ensures that a redeem swap coin redeems a certain contract by
// comparing the contract and secret hash.
func (c *swapCoin) validateRedeem(contractID []byte) error {
	if c.sct != sctRedeem {
		return errors.New("can only validate redeem")
	}
	cID, err := DecodeCoinID(contractID)
	if err != nil {
		return err
	}
	scID, ok := cID.(*SwapCoinID)
	if !ok {
		return errors.New("contract ID not a swap")
	}
	if c.secretHash != scID.SecretHash {
		return fmt.Errorf("redeem secret hash %x does not match contract %x",
			c.secretHash, scID.SecretHash)
	}
	if c.contractAddr != scID.ContractAddress {
		return fmt.Errorf("redeem contract address %q does not match contract address %q",
			c.contractAddr, scID.ContractAddress)
	}
	return nil
}

// Confirmations returns the number of confirmations for a Coin's
// transaction.
//
// In the case of ethereum it is extra important to check confirmations before
// confirming a swap. Even though we check the initial transaction's data, if
// that transaction were in mempool at the time, it could be swapped out with
// any other values if a user sent another transaction with a higher gas fee
// and the same account and nonce, effectivly voiding the transaction we
// expected to be mined.
func (c *swapCoin) Confirmations(_ context.Context) (int64, error) {
	swap, err := c.backend.node.swap(c.backend.rpcCtx, c.secretHash)
	if err != nil {
		return -1, err
	}

	switch c.sct {
	case sctRedeem:
		// There should be no need to check the counter party, or value
		// as a swap with a specific secret hash that has been redeemed
		// wouldn't have been redeemed without ensuring the initiator
		// is the expected address and value was also as expected. Also
		// not validating the locktime, as the swap is redeemed and
		// locktime no longer relevant.
		ss := SwapState(swap.State)
		if ss == SSRedeemed {
			// While not completely accurate, we know that if the
			// swap is redeemed the redemption has at least one
			// confirmation.
			return 1, nil
		}
		// If swap is in the Initiated state, the transaction may be
		// unmined.
		if ss == SSInitiated {
			// Assume the tx still has a chance of being mined.
			return 0, nil
		}
		// If swap is in None state, then the redemption can't possibly
		// succeed as the swap must already be in the Initialized state
		// to redeem. If the swap is in the Refunded state, then the
		// redemption either failed or never happened.
		return -1, fmt.Errorf("redemption in failed state with swap at %s state", ss)

	case sctInit:
		// Uninitiated state is zero confs. It could still be in mempool.
		// It is important to only trust confirmations according to the
		// swap contract. Until there are confirmations we cannot be sure
		// that initiation happened successfuly.
		if SwapState(swap.State) == SSNone {
			// Assume the tx still has a chance of being mined.
			return 0, nil
		}
		// Any other swap state is ok. We are sure that initialization
		// happened.

		// The swap initiation transaction has some number of
		// confirmations, and we are sure the secret hash belongs to
		// this swap. Assert that the value, reciever, and locktime are
		// as expected.
		value, err := ToGwei(big.NewInt(0).Set(swap.Value))
		if err != nil {
			return -1, fmt.Errorf("unable to convert value: %v", err)
		}
		if value != c.value {
			return -1, fmt.Errorf("expected swap val (%dgwei) does not match expected (%dgwei)",
				c.value, value)
		}
		if swap.Participant != c.counterParty {
			return -1, fmt.Errorf("expected swap participant %q does not match expected %q",
				c.counterParty, swap.Participant)
		}
		if !swap.RefundBlockTimestamp.IsInt64() {
			return -1, errors.New("swap locktime is larger than expected")
		}
		locktime := swap.RefundBlockTimestamp.Int64()
		if locktime != c.locktime {
			return -1, fmt.Errorf("expected swap locktime (%d) does not match expected (%d)",
				c.locktime, locktime)
		}

		bn, err := c.backend.node.blockNumber(c.backend.rpcCtx)
		if err != nil {
			return 0, fmt.Errorf("unable to fetch block number: %v", err)
		}
		return int64(bn - swap.InitBlockNumber.Uint64()), nil
	}

	return -1, fmt.Errorf("unsupported swap type for confirmations: %d", c.sct)
}

// ID is the swap's coin ID.
func (c *swapCoin) ID() []byte {
	sc := &SwapCoinID{
		ContractAddress: c.contractAddr,
		SecretHash:      c.secretHash,
	}
	return sc.Encode()
}

// TxID is the original init transaction txid.
func (c *swapCoin) TxID() string {
	return c.txid
}

// String is a human readable representation of the swap coin.
func (c *swapCoin) String() string {
	sc := &SwapCoinID{
		ContractAddress: c.contractAddr,
		SecretHash:      c.secretHash,
	}
	return sc.String()
}

// Value is the amount paid to the swap, set in initialization. Always zero for
// redemptions.
func (c *swapCoin) Value() uint64 {
	return c.value
}

// FeeRate returns the gas rate, in gwei/gas. It is set in initialization of
// the swapCoin.
func (c *swapCoin) FeeRate() uint64 {
	return c.gasPrice
}
