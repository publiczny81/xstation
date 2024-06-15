package model

const (
	CmdGetAllSymbols          = "getAllSymbols"
	CmdGetCalendar            = "getCalendar"
	CmdGetChartLastRequest    = "getChartLastRequest"
	CmdGetChartRangeRequest   = "getChartRangeRequest"
	CmdGetCommissionDef       = "getCommissionDef"
	CmdGetCurrentUserData     = "getCurrentUserData"
	CmdGetIbsHistory          = "getIbsHistory"
	CmdLogin                  = "login"
	CmdLogout                 = "logout"
	CmdGetMarginLevel         = "getMarginLevel"
	CmdGetMarginTrade         = "getMarginTrade"
	CmdGetNews                = "getNews"
	CmdGetProfitCalculation   = "getProfitCalculation"
	CmdGetServerTime          = "getServerTime"
	CmdGetStepRules           = "getStepRules"
	CmdGetSymbol              = "getSymbol"
	CmdGetTickPrices          = "getTickPrices"
	CmdGetTradeRecords        = "getTradeRecords"
	CmdGetTrades              = "getTrades"
	CmdGetTradesHistory       = "getTradesHistory"
	CmdGetTradingHours        = "getTradingHours"
	CmdGetVersion             = "getVersion"
	CmdPing                   = "ping"
	CmdTradeTransaction       = "tradeTransaction"
	CmdTradeTransactionStatus = "tradeTransactionStatus"
)

const (
	BE000 = "BE000"
	BE001 = "BE001"
	BE002 = "BE002"
	BE003 = "BE003"
	BE004 = "BE004"
	BE005 = "BE005"
	BE006 = "BE006"
	BE007 = "BE007"
	BE008 = "BE008"
	BE009 = "BE009"
	BE010 = "BE010"
	BE011 = "BE011"
	BE012 = "BE012"
	BE013 = "BE013"
	BE014 = "BE014"
	BE016 = "BE016"
	BE017 = "BE017"
	BE018 = "BE018"
	BE019 = "BE019"
	BE020 = "BE020"
	BE021 = "BE021"
	BE022 = "BE022"
	BE023 = "BE023"
	BE024 = "BE024"
	BE025 = "BE025"
	BE026 = "BE026"
	BE027 = "BE027"
	BE028 = "BE028"
	BE029 = "BE029"
	BE030 = "BE030"
	BE031 = "BE031"
	BE032 = "BE032"
	BE033 = "BE033"
	BE034 = "BE034"
	BE035 = "BE035"
	BE036 = "BE036"
	BE037 = "BE037"
	BE094 = "BE094"
	BE095 = "BE095"
	BE096 = "BE096"
	BE097 = "BE097"
	BE098 = "BE098"
	BE099 = "BE099"
	BE101 = "BE101"
	BE102 = "BE102"
	BE103 = "BE103"
	BE104 = "BE104"
	BE105 = "BE105"
	BE106 = "BE106"
	BE110 = "BE110"
	BE115 = "BE115"
	BE116 = "BE116"
	BE117 = "BE117"
	BE118 = "BE118"
	BE200 = "BE200"
	EX000 = "EX000"
	EX001 = "EX001"
	EX002 = "EX002"
	EX003 = "EX003"
	EX004 = "EX004"
	EX005 = "EX005"
	EX006 = "EX006"
	EX007 = "EX007"
	EX008 = "EX008"
	EX009 = "EX009"
	EX010 = "EX010"
	EX011 = "EX011"
	SEXXX = "SEXXX"
)

const (
	ErrDescInvalidPrice                   = "Invalid price"
	ErrDescInvalidStopLossOrTakeProfit    = "Invalid StopLoss or TakeProfit"
	ErrDescInvalidVolume                  = "Invalid volume"
	ErrDescLoginDisabled                  = "Login disabled"
	ErrDescUserPasswordCheck              = "userPasswordCheck: Invalid login or password."
	ErrDescMarketForInstrumentClosed      = "Market for instrument is closed"
	ErrDescMismatchedParameters           = "Mismatched parameters"
	ErrDescModificationDenied             = "Modification is denied"
	ErrDescNotEnoughMoneyOnAccount        = "Not enough money on account to perform trade"
	ErrDescOffQuotes                      = "Off quotes"
	ErrDescOppositePositionsProhibited    = "Opposite positions prohibited"
	ErrDescShortPositionsProhibited       = "Short positions prohibited"
	ErrDescPriceHasChanged                = "Price has changed"
	ErrDescRequestTooFrequent             = "Request too frequent"
	ErrDescTooManyTradeRequests           = "Too many trade requests"
	ErrDescTradingOnInstrumentDisabled    = "Trading on instrument disabled"
	ErrDescTradingTimeout                 = "Trading timeout"
	ErrDescOtherError                     = "Other error"
	ErrDescSymbolNotExistsForGivenAccount = "Symbol does not exist for given account"
	ErrDescTradeNotAllowedOnGivenSymbol   = "Account cannot trade on given symbol"
	ErrDescPendingOrderCannotBeClosed     = "Pending order cannot be closed. Pending order must be deleted"
	ErrDescOrderAlreadyClosed             = "Cannot close already closed order"
	ErrDescNoSuchTransaction              = "No such transaction"
	ErrDescUnknownSymbol                  = "Unknown instrument symbol"
	ErrDescUnknownTransactionType         = "Unknown transaction type"
	ErrDescUserNotLogged                  = "User is not logged"
	ErrDescMethodNotExists                = "Method does not exist"
	ErrDescIncorrectPeriodGiven           = "Incorrect period given"
	ErrDescMissingData                    = "Missing data"
	ErrDescIncorrectCommandFormat         = "Incorrect command format"
	ErrDescSymbolNotExits                 = "Symbol does not exist"
	ErrDescInvalidToken                   = "Invalid token"
	ErrDescUserAlreadyLogged              = "User already logged"
	ErrDescSessionTimeout                 = "Session timed out."
	ErrDescInvalidParameters              = "Invalid parameters"
	ErrDescInternalError                  = "Internal error, in case of such error, please contact support"
	ErrDescInternalErrorRequestTimeout    = "Internal error, request timed out"
	ErrDescIncorrectCredentials           = "Login credentials are incorrect or this login is not allowed to use an application with this appId"
	ErrDescInternalErrorSystemOverloaded  = "Internal error, system overloaded"
	ErrDescNoAccess                       = "No access"
	ErrDescLoginTemporaryDisabled         = "userPasswordCheck: Invalid login or password. This login/password is disabled for 10 minutes (the specific login and password pair is blocked after an unsuccessful login attempt)."
	ErrDescReachedConnectionLimit         = "You have reached the connection limit. For details see the Connection validation section"
	ErrDescDataLimitExceeded              = "Data limit potentially exceeded. Please narrow your request range. The potential data size is calculated by: (end_time - start_time) / interval. The limit is 50 000 candles"
	ErrDescLoginOnBlackList               = "Your login is on the black list, perhaps due to previous misuse. For details please contact support."
	ErrDescCommandExecutionNotAllowed     = "You are not allowed to execute this command. For details please contact support."
)

const (
	ImpactLow    = "1"
	ImpactMedium = "2"
	ImpactHigh   = "3"
)

const (
	MarginModeForex        = 101
	MarginModeCFDLeveraged = 102
	MarginModeCFD          = 103
)

const (
	PeriodM1  = 1
	PeriodM5  = 5
	PeriodM15 = 15
	PeriodM30 = 30
	PeriodH1  = 60
	PeriodH4  = 240
	PeriodD1  = 1440
	PeriodW1  = 10080
	PeriodMN1 = 43200
)

const (
	ProfitModeForex = 5
	ProfitModeCFD   = 6
)

const (
	QuoteIdFixed = 1
	QuoteIdFloat = 2
	QuoteIdDepth = 3
	QuoteIdCross = 4
)

const (
	LevelAll       = -1 // all available levels
	LevelBasePrice = 0  // base level bid and ask price for instrument
)
