package common

const (
	QuoteIdFixed = 1
	QuoteIdFloat = 2
	QuoteIdDepth = 3
	QuoteIdCross = 4
)

const (
	OperationCodeBuy       = 0 // buy
	OperationCodeSell      = 1 // sell
	OperationCodeBuyLimit  = 2 // buy limit
	OperationCodeSellLimit = 3 // sell limit
	OperationCodeBuyStop   = 4 // buy stop
	OperationCodeSellStop  = 5 // sell stop
	OperationCodeBalance   = 6 // Read only. Used in getTradesHistory  for manager's deposit/withdrawal operations (profit>0 for deposit, profit<0 for withdrawal)
	OperationCodeCredit    = 7 // Read only
)

const (
	OperationTypeOpen    = 0
	OperationTypePending = 1
	OperationTypeClose   = 2
	OperationTypeModify  = 3
	OperationTypeDelete  = 4
)
