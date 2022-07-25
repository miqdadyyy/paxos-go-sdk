package constant

const (
	ScopeFundingReadProfile = "funding:read_profile"
)

const (
	ScopeExchangeReadQuote           = "exchange:read_quote"
	ScopeExchangeWriteQuoteExecution = "exchange:write_quote_execution"
	ScopeExchangeReadQuoteExecution  = "exchange:read_quote_execution"
	ScopeExchangeWriteOrder          = "exchange:write_order"
	ScopeExchangeReadOrder           = "exchange:read_order"
	ScopeExchangeHistoricalPrices    = "exchange:historical_prices"
)

const (
	ScopeTransferReadTransfer          = "transfer:read_transfer"
	ScopeTransferReadDepositAddress    = "transfer:read_deposit_address"
	ScopeTransferWriteDepositAddress   = "transfer:write_deposit_address"
	ScopeTransferWriteCryptoWithdrawal = "transfer:write_crypto_withdrawal"
	ScopeWriteCryptoWithdrawalFee      = "fee:write_crypto_withdrawal_fee"
)
