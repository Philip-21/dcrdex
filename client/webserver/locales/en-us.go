package locales

var EnUS = map[string]string{
	"Language":                       "en-US", // the bcp47 lang tag
	"Markets":                        "Markets",
	"Wallets":                        "Wallets",
	"Notifications":                  "Notifications",
	"Recent Activity":                "Recent Activity",
	"Sign Out":                       "Sign Out",
	"Order History":                  "Order History",
	"load from file":                 "load from file",
	"loaded from file":               "loaded from file",
	"defaults":                       "defaults",
	"Wallet Password":                "Wallet Password",
	"w_password_helper":              "This is the password you have configured with your wallet software.",
	"w_password_tooltip":             "Leave the password empty if there is no password required for the wallet.",
	"App Password":                   "App Password",
	"Add":                            "Add",
	"Unlock":                         "Unlock",
	"Rescan":                         "Rescan",
	"Wallet":                         "Wallet",
	"app_password_reminder":          "Your app password is always required when performing sensitive wallet operations.",
	"DEX Address":                    "DEX Address",
	"TLS Certificate":                "TLS Certificate",
	"remove":                         "remove",
	"add a file":                     "add a file",
	"Submit":                         "Submit",
	"Skip Registration":              "No account (view-only mode)",
	"Confirm Registration":           "Confirm Registration and Bonding",
	"app_pw_reg":                     "Enter your app password to confirm DEX registration and bond creation.",
	"reg_confirm_submit":             `When you submit this form, funds will be spent from your wallet to post a fidelity bond, which is redeemable by you in the future.`,
	"bond_strength":                  "Bond Strength",
	"update_bond_options":            "Update Bond Options",
	"bond_options_update_success":    "Bond Options have been updated successfully",
	"target_tier":                    "Target Tier",
	"target_tier_tooltip":            "This is the target account tier you wish to maintain. Set to zero if you wish to disable tier maintenance (do not post new bonds).",
	"provided_markets":               "This DEX provides the following markets:",
	"accepted_fee_assets":            "This DEX recognizes the bond assets:",
	"base_header":                    "Base",
	"quote_header":                   "Quote",
	"lot_size_headsup":               "All trades are in multiples of the lot size.",
	"Password":                       "Password",
	"Register":                       "Register",
	"Authorize Export":               "Authorize Export",
	"export_app_pw_msg":              "Enter your app password to confirm account export for",
	"Disable Account":                "Disable Account",
	"disable_app_pw_msg":             "Enter your app password to disable account",
	"disable_dex_server":             "This DEX server may be re-enabled at any time in the future by adding it again.",
	"Authorize Import":               "Authorize Import",
	"app_pw_import_msg":              "Enter your app password to confirm account import",
	"Account File":                   "Account File",
	"Change Application Password":    "Change Application Password",
	"Current Password":               "Current Password",
	"New Password":                   "New Password",
	"Confirm New Password":           "Confirm New Password",
	"cancel_pw":                      "Enter your password to submit a cancel order for the remaining",
	"cancel_no_pw":                   "Submit a cancel order for the remaining",
	"cancel_remain":                  "The remaining amount may change before the cancel order is matched.",
	"Log In":                         "Log In",
	"epoch":                          "epoch",
	"price":                          "price",
	"volume":                         "volume",
	"buys":                           "buys",
	"sells":                          "sells",
	"Buy Orders":                     "Buy Orders",
	"Quantity":                       "Quantity",
	"Rate":                           "Rate",
	"Limit Order":                    "Limit Order",
	"Market Order":                   "Market Order",
	"create_account_to_trade":        "Create an account to trade",
	"need_to_register_msg":           `You need to create an account on <span id="unregisteredDex"></span> to trade.`,
	"Create Account":                 "Create Account",
	"reg_status_msg":                 `In order to trade at <span id="regStatusDex" class="text-break"></span>, your pending bond(s) need to be confirmed.`,
	"action_required_to_trade":       "ACTION REQUIRED TO TRADE",
	"acct_tier_post_bond":            `You account tier is <span id="acctTier"></span>. You need to post new bonds to trade.`,
	"enable_bond_maintenance":        "Enable bond maintenance from DEX Settings page.",
	"Buy":                            "Buy",
	"Sell":                           "Sell",
	"lot_size":                       "Lot Size",
	"Rate Step":                      "Rate Step",
	"Max":                            "Max",
	"lot":                            "lot",
	"min trade is about":             "min trade is about",
	"immediate_explanation":          "If the order doesn't fully match during the next match cycle, any unmatched quantity will not be booked or matched again. Taker-only order.",
	"Immediate or cancel":            "Immediate or cancel",
	"Balances":                       "Balances",
	"outdated_tooltip":               "Balance may be outdated. Connect to the wallet to refresh.",
	"available":                      "available",
	"connect_refresh_tooltip":        "Click to connect and refresh",
	"add_a_wallet":                   `Add a <span data-tmpl="addWalletSymbol"></span> wallet`,
	"locked":                         "locked",
	"immature":                       "immature",
	"fee balance":                    "fee balance",
	"Sell Orders":                    "Sell Orders",
	"Your Orders":                    "Your Orders",
	"Recent Matches":                 "Recent Matches",
	"Type":                           "Type",
	"Side":                           "Side",
	"Age":                            "Age",
	"Filled":                         "Filled",
	"Settled":                        "Settled",
	"Status":                         "Status",
	"view order history":             "view order history",
	"cancel_order":                   "cancel order",
	"order details":                  "order details",
	"verify_order":                   `Verify <span id="vSideHeader"></span>  Order`,
	"prevent_temporary_overlocking":  "Prevent temporary over-locking of funds",
	"You are submitting an order to": "You are submitting an order to",
	"at a rate of":                   "at a rate of",
	"for a total of":                 "for a total of",
	"verify_market":                  "This is a market order and will match the best available order(s) on the book. Based on the current market mid-gap rate, you might receive about",
	"auth_order_app_pw":              "Authorize this order with your app password.",
	"lots":                           "lots",
	"order_disclaimer": `<span class="red">IMPORTANT</span>: Trades take time to settle, and you cannot turn off the DEX client software,
		or the <span data-unit="quote"></span> or <span data-unit="base"></span> blockchain and/or wallet software, until
		settlement is complete. Settlement can complete as quickly as a few minutes or take as long as a few hours.`,
	"acknowledge_and_hide":      "acknowledge and hide",
	"show_disclaimer":           "view warnings",
	"Order":                     "Order",
	"see all orders":            "see all orders",
	"Exchange":                  "Exchange",
	"Market":                    "Market",
	"Offering":                  "Offering",
	"Asking":                    "Asking",
	"Fees":                      "Fees",
	"order_fees_tooltip":        "On-chain transaction fees, typically collected by the miner. Decred DEX collects no trading fees.",
	"Matches":                   "Matches",
	"Match ID":                  "Match ID",
	"Time":                      "Time",
	"ago":                       "ago",
	"Cancellation":              "Cancellation",
	"Order Portion":             "Order Portion",
	"you":                       "you",
	"them":                      "them",
	"Redemption":                "Redemption",
	"Refund":                    "Refund",
	"Funding Coins":             "Funding Coins",
	"Exchanges":                 "Exchanges",
	"apply":                     "apply",
	"Assets":                    "Assets",
	"Trade":                     "Trade",
	"Set App Password":          "Set App Password",
	"reg_set_app_pw_msg":        "Set your app password. This password will protect your DEX account keys and connected wallets.",
	"Password Again":            "Password Again",
	"Add a DEX":                 "Add a DEX",
	"Pick a server":             "Pick a server",
	"reg_ssl_needed":            "Looks like we don't have an SSL certificate for this DEX. Add the server's certificate to continue.",
	"Dark Mode":                 "Dark Mode",
	"Show pop-up notifications": "Show pop-up notifications",
	"Account ID":                "Account ID",
	"Export Account":            "Export Account",
	"simultaneous_servers_msg":  "The Decred DEX Client supports simultaneous use of any number of DEX servers.",
	"Change App Password":       "Change App Password",
	"Build ID":                  "Build ID",
	"Connect":                   "Connect",
	"Send":                      "Send",
	"Deposit":                   "Deposit",
	"Lock":                      "Lock",
	"New Deposit Address":       "New Deposit Address",
	"Address":                   "Address",
	"Amount":                    "Amount",
	"Authorize the transfer with your app password.": "Authorize the transfer with your app password.",
	"Reconfigure":                 "Reconfigure",
	"pw_change_instructions":      "Changing the password below does not change the password for your wallet software. Use this form to update the DEX client after you have changed your password with the wallet software directly.",
	"New Wallet Password":         "New Wallet Password",
	"pw_change_warn":              "Note: Changing to a different wallet while having active trades might cause funds to be lost.",
	"Show more options":           "Show more options",
	"seed_implore_msg":            "You should carefully write down your application seed and save a copy. Should you lose access to this machine or the critical application files, the seed can be used to restore your DEX accounts and native wallets. Some older accounts cannot be restored from seed, and whether old or new, it's good practice to backup the account keys separately from the seed.",
	"View Application Seed":       "View Application Seed",
	"Remember my password":        "Remember my password",
	"pw_for_seed":                 "Enter your app password to show your seed. Make sure nobody else can see your screen.",
	"Asset":                       "Asset",
	"Balance":                     "Balance",
	"Actions":                     "Actions",
	"Restoration Seed":            "Restoration Seed",
	"Restore from seed":           "Restore from seed",
	"Import Account":              "Import Account",
	"no_wallet":                   "no wallet",
	"create_a_x_wallet":           "Create a <span data-asset-name=1></span> Wallet",
	"dont_share":                  "Don't share it. Don't lose it.",
	"Show Me":                     "Show Me",
	"Wallet Settings":             "Wallet Settings",
	"add_a_x_wallet":              `Add a <img data-tmpl="assetLogo" class="asset-logo mx-1"> <span data-tmpl="assetName"></span> Wallet`,
	"ready":                       "ready",
	"off":                         "off",
	"Export Trades":               "Export Trades",
	"change the wallet type":      "change the wallet type",
	"confirmations":               "confirmations",
	"how_reg":                     "How will you post bond?",
	"All markets at":              "All markets at",
	"pick a different asset":      "pick a different asset",
	"Create":                      "Create",
	"Register_loudly":             "Register!",
	"1 Sync the Blockchain":       "1: Sync the Blockchain",
	"Progress":                    "Progress",
	"remaining":                   "remaining",
	"2 Fund your Wallet":          "2: Fund your Wallet",
	"whatsabond":                  "Fidelity bonds are time-locked funds redeemable only by you, but in the future. This is meant to combat disruptive behavior like backing out on swaps.",
	"Bond amount":                 "Bond amount",
	"Reserves for tx fees":        "Funds to reserve for transaction fees to maintain your bonds",
	"Tx Fee Balance":              "Transaction Fee Balance:",
	"Your Deposit Address":        "Your Wallet's Deposit Address",
	"Send enough for bonds":       `Make sure you send enough to also cover network fees. You may deposit as much as you like to your wallet, since only the bond amount will be used in the next step. The deposit must confirm to proceed.`,
	"Send enough with estimate":   `Deposit at least <span data-tmpl="totalForBond"></span> <span class="unit">XYZ</span> to cover network fees and overlap periods when a bond is expired (but waiting for refund) and another must be posted. You may deposit as much as you like to your wallet, since only the required amount will be used in the next step. The deposit must confirm to proceed.`,
	"Send funds for token":        `Deposit at least <span data-tmpl="tokenFees"></span> <span class="unit">XYZ</span> and <span data-tmpl="parentFees"></span> <span data-tmpl="parentUnit">XYZ</span> to also cover fees. You may deposit as much as you like to your wallet, since only the required amount will be used in the next step. The deposit must confirm to proceed.`,
	"add a different server":      "add a different server",
	"Add a custom server":         "Add a custom server",
	"plus tx fees":                "+ tx fees",
	"Export Seed":                 "Export Seed",
	"Total":                       "Total",
	"Trading":                     "Trading",
	"Receiving Approximately":     "Receiving Approximately",
	"Fee Projection":              "Fee Projection",
	"details":                     "details",
	"to":                          "to",
	"Options":                     "Options",
	"fee_projection_tooltip":      "If network conditions don't change before your order matches, total realized fees should fall within this range.",
	"unlock_for_details":          "Unlock your wallets to retrieve order details and additional options.",
	"estimate_unavailable":        "Order estimates and options unavailable",
	"Fee Details":                 "Fee Details",
	"estimate_market_conditions":  "Best- and worst-case estimates are based on current network conditions, and might change by the time the order matches.",
	"Best Case Fees":              "Best Case Fees",
	"best_case_conditions":        "The best case fees occur when the entire order is consumed in a single match.",
	"Swap":                        "Swap",
	"Redeem":                      "Redeem",
	"Worst Case Fees":             "Worst Case Fees",
	"worst_case_conditions":       "The worst case can occur if the order matches one lot at a time over the course of many epochs.",
	"Maximum Possible Swap Fees":  "Maximum Possible Swap Fees",
	"max_fee_conditions":          "This is the most you would ever pay in fees on your swap. Fees are normally assessed at a fraction of this rate. The maximum is not subject to changes once your order is placed.",
	"wallet_logs":                 "Wallet Logs",
	"accelerate_order":            "accelerate order",
	"acceleration_text":           "If your swap transactions are stuck, you may attempt to accelerate them with an additional transaction. This is helpful when the the fee rate on an existing unconfirmed transaction has become too low to be mined in the next block, but not if blocks are just being mined slowly. When you submit this form, a transaction will be created that sends the change from your own swap initiation transaction to yourself with a higher fee. The effective fee rate of your swap transactions will become the rate you select below. Select a rate that is enough to be included in the next block. Consult a block explorer to be sure.",
	"effective_swap_tx_rate":      "Effective swap tx fee rate",
	"current_fee":                 "Current suggested fee rate",
	"accelerate_success":          `Successfully submitted transaction: <span id="accelerateTxID"></span>`,
	"accelerate":                  "Accelerate",
	"acceleration_transactions":   "Acceleration Transactions",
	"acceleration_cost_msg":       `Increasing the effective fee rate to <span id="feeRateEstimate"></span> will cost <span id="feeEstimate"></span>`,
	"recent_acceleration_msg":     `Your latest acceleration was only <span id="recentAccelerationTime"></span> minutes ago! Are you sure you want to accelerate?`,
	"recent_swap_msg":             `Your oldest unconfirmed swap transaction was submitted only <span id="recentSwapTime"></span> minutes ago! Are you sure you want to accelerate?`,
	"early_acceleration_help_msg": `It will cause no harm to your order, but you might be wasting money. Acceleration is only helpful if the fee rate on an existing unconfirmed transaction has become too low to be mined in the next block, but not if blocks are just being mined slowly. You can confirm this in the block explorer by closing this popup and clicking on your previous transactions.`,
	"recover":                     "Recover",
	"recover_wallet":              "Recover Wallet",
	"recover_warning":             "Recovering your wallet will move all wallet data to a backup folder. You will have to wait until the wallet syncs with the network, which could potentially take a long time, before you can use the wallet again.",
	"wallet_actively_used":        "Wallet actively used!",
	"confirm_force_message":       "This wallet is actively managing orders. After taking this action, it will take a long time to resync your wallet, potentially causing orders to fail. Only take this action if absolutely necessary!",
	"confirm":                     "Confirm",
	"cancel":                      "Cancel",
	"Update TLS Certificate":      "Update TLS Certificate",
	"registered dexes":            "Registered Dexes:",
	"successful_cert_update":      "Successfully updated certificate!",
	"update dex host":             "Update DEX Host",
	"copied":                      "Copied!",
	"export_wallet":               "Export Wallet",
	"pw_for_wallet_seed":          "Enter your app password to show the wallet seed. Make sure nobody else can see your screen. If anyone gets access to the wallet seed, they will be able to steal all of your funds.",
	"export_wallet_disclaimer":    `<span class="warning-text">Using an externally restored wallet while you have active trades running in the DEX could result in failed trades and LOST FUNDS.</span> It is recommended that you do not export your wallet unless you are an experienced user and you know what are doing.`,
	"export_wallet_msg":           "Below are the seeds needed to restore your wallet in some popular external wallets. DO NOT make transactions with your external wallet while you have active trades running on the DEX.",
	"clipboard_warning":           "Copy/Pasting a wallet seed is a potential security risk. Do this at your own risk.",
	"fiat_exchange_rate_sources":  "Fiat Exchange Rate Sources",
	"Synchronizing":               "Synchronizing",
	"wallet_wait_synced":          "wallet will be created after sync",
	"Create a Wallet":             "Create a Wallet",
	"Receive":                     "Receive",
	"Wallet Type":                 "Wallet Type",
	"Peer Count":                  "Peer Count",
	"Sync Progress":               "Sync Progress",
	"Settings":                    "Settings",
	"asset_name Markets":          "<span data-asset-name=1></span> Markets",
	"Host":                        "Host",
	"No Recent Activity":          "No Recent Activity",
	"Recent asset_name Activity":  "Recent <span data-asset-name=1></span> Activity",
	"other_actions":               "Other Actions",
	"estimated_fee":               "Estimated Fee",
	"estimated_total_spend":       "Estimated Total Spend",
	"estimated_balance":           "Estimated Balance After",
	"max_estimated_send":          "Max Estimated Send",
	"max_estimated_send_fee":      "Max Estimated Send Fee",
	"sending":                     "Sending",
	"transfer":                    "Transfer",
	"max_estimated_send_tooltip":  "This is the estimated amount that will be received if you withdraw your current balance with 'Subtract fees from amount sent' checked. If there is no subtract fee checkbox, this is the maximum estimated amount you can send.",
	"Maker Swap":                  "Maker Swap",
	"Taker Swap":                  "Taker Swap",
	"Maker Redemption":            "Maker Redemption",
	"Taker Redemption":            "Taker Redemption",
	"Pending":                     "Pending",
	"disable_wallet":              "Disable Wallet",
	"enable_wallet":               "Enable Wallet",
	"disable_wallet_warning":      "Note: This wallet will not be connected to when you start the DEX client software and cannot be used until it is enabled. This will also disable all token wallets that depend on this wallet.",
	"enable_wallet_message":       "This wallet will resume operation and might take some time to sync. If this is a token wallet, the wallet for the chain's primary asset will also be enabled.",
	"Disabled":                    "Disabled",
	"txfee_not_available":         "Transaction fee currently unavailable",
	"fiat_exchange_rate_msg":      "Sources may provide fiat exchange rates for different subsets of assets. You should select all sources that are acceptable to get fiat exchange rates for the most assets. Assets with fiat rates from multiple sources will use the average from all of those sources. Note: dcrdate provides fiat exchange rates for only BTC and DCR.",
	"delete_archived_records":     "Delete Archived Records",
	"date_time":                   "Date and Time",
	"delete_all_archived_records": "Leave unchecked to delete all archived records.",
	"show_archived_date_msg":      "Specify date of latest records to keep",
	"archived_date_tooltip":       "Archived orders and matches created before your specified and date and time will be deleted from the database.",
	"save_matches_to_file":        "Save matches to CSV file",
	"save_orders_to_file":         "Save orders to CSV file",
	"save_orders_to_file_msg":     "Optional: Whether to save deleted orders to CSV file on dexc data directory. Default is false.",
	"save_matches_to_file_msg":    "Optional: Whether to save deleted matches to CSV file on dexc data directory. Default is false.",
	// Market maker bot
	"Market Making":          "Market Making",
	"Make a Market":          "Make a Market!",
	"Editing Program":        "Editing Program",
	"exit edit mode":         "exit edit mode",
	"per side":               "per side",
	"Start_loudly":           "Start!",
	"Update":                 "Update",
	"Show other settings":    "Show other settings",
	"Hide settings":          "Hide settings",
	"lot_commit_bullet":      "Total commitment is 2 x lots",
	"funds_split_bullet":     "Starting funds can be on one side or split up",
	"target_maint_bullet":    "Program maintains lots currently on order and unfilled",
	"no_limit_bullet":        "There is no limit on the number of settling lots",
	"Your Programs":          "Your Programs",
	"No programs to display": "No programs to display",
	"Running":                "Running",
	"Pause":                  "Pause",
	"Paused":                 "Paused",
	"Start":                  "Start",
	"Drift tolerance":        "Drift tolerance",
	"Oracle bias":            "Oracle bias",
	"Multiplier":             "Multiplier",
	"Oracle weight":          "Oracle weight",
	"Configure":              "Configure",
	"Retire":                 "Retire",
	"Lots":                   "Lots",
	"show advanced options":  "show advanced options",
	"hide advanced options":  "hide advanced options",
	"manage_peers":           "Manage Peers",
	"add_peer":               "Add Peer",
	"address":                "Address",
	"source":                 "Source",
	"connected":              "Connected",
	"Remove":                 "Remove",
	"unready_wallets_msg":    "Your wallets must be connected and unlocked before trades can be processed. Resolve this ASAP!",
	"Error":                  "Error",
	"configuration guide":    "configuration guide",
	// Approve/Unapprove Tokens
	"Approve":                 "Approve",
	"approve_token_text":      `In order to trade tokens, you must approve the swap contract to handle tokens on your behalf. This is a one-time action for each token. <br/><br/> The estimated fee for the approval transaction is <b><span id="approvalFeeEstimate"></span></b>`,
	"token_approval_tx_msg":   "Your token approval has been submitted with transaction ID:",
	"approval_required_buy":   "Token approval is required before buying",
	"approval_required_sell":  "Token approval is required before selling",
	"approval_required_both":  "Token approval is required before trading",
	"disallow_token":          `Remove Token Allowances`,
	"unapprove_token_for":     `Remove Token Allowances for <span id="tokenVersionTableAssetSymbol"></span>`,
	"unapprove_token_version": `Remove Token Allowances for <span id="tokenAllowanceRemoveSymbol"></span>, Version <span id="tokenAllowanceRemoveVersion"></span>`,
	"unapprove_token_text":    `If you remove the allowance for this version of the swap contract, you will no longer be able to trade until you re-allow it. <br/><br/> The estimated transaction fee for removing the approval is <b><span id="unapprovalFeeEstimate"></span></b>`,
	"version":                 "Version",
	"used_by_dex":             "Used by DEXes",
	"no_token_allowances":     "You have not granted allowances for any swap contracts for this token.",
	"token_unapproval_tx_msg": `Your token approval has been removed with transaction ID:`,
	"approval_change_pending": "Approval change pending",
}
