package xapi

import (
	"github.com/publiczny81/xstation/xapi/model"
	"github.com/publiczny81/xstation/xapi/stream"
)

// Login logins to the service
//
//	conn - connection to the servier
//	user - user id
//	password - password
//	opts - options for the request
//
// The options allowed for the request are:
//   - appId - application id (deprecated)
//   - appName - application name
//
// In order to perform any action client application have to perform login process. No functionality is available before proper login process.
// After initial login, a new session is created and all commands are executed for a logged user until he/she logs out or drops the connection
//
// After successful login the system responds with the status message that can contain the String representing streamSessionId field.
// The streamSessionId field of the string, if present, is a token that can be used to establish a streaming subscription on a separate
// network connection. streamSessionId is used in streaming subscription commands.
// streamSessionId is unique for the given main session and will change between login sessions.
func Login(conn connection, user, password string, opts ...model.LoginRequestOption) (resp *model.LoginResponse, err error) {
	resp, err = sendReceive[model.LoginResponse](conn, model.NewLoginRequest(user, password, opts...))
	return
}

// Logout logouts from the session
//
// It returns no returnData field in output. Only status message is sent
func Logout(conn connection) (resp *model.LogoutResponse, err error) {
	resp, err = sendReceive[model.LogoutResponse](conn, model.NewLogoutRequest())
	return
}

// GetAllSymbols returns array of all symbols available for the user.
//
//	conn - connection to the servier
//	opts - options for the request
//
// The allowed options are:
//   - customTag - custom tag
//   - prettyPrint - pretty print flag
func GetAllSymbols(conn connection, opts ...model.GetAllSymbolsRequestOption) (resp *model.GetAllSymbolsResponse, err error) {
	resp, err = sendReceive[model.GetAllSymbolsResponse](conn, model.NewGetAllSymbolsRequest(opts...))
	return
}

// GetCalendar returns calendar with market events
//
//	conn - connection to the servier
func GetCalendar(conn connection, opts ...model.GetCalendarRequestOption) (resp *model.GetCalendarResponse, err error) {
	return sendReceive[model.GetCalendarResponse](conn, model.NewGetCalendarRequest(opts...))
}

// GetChartLastRequest returns chart info, from start date to the current time. If the chosen period of model.ChartLastInfoRecord
// is greater than 1 minute, the last candle returned by the API can change until the end of the period (the candle is being
// automatically updated every minute)
//
//	conn - connection to the servier
//	symbol - symbol
//	period - period code
//	start - start of chart block (rounded down to the nearest interval and excluding)
func GetChartLastRequest(conn connection, symbol string, period, start int, opts ...model.GetChartLastRequestOption) (*model.GetChartLastResponse, error) {
	return sendReceive[model.GetChartLastResponse](conn,
		model.NewGetChartLastRequest(symbol, period, start, opts...))
}

// GetChartRangeRequest returns chart info with data between given start and end dates
//
//	c - connection to a server
//	symbol - symbol
//	period - period code
//	start - start of chart block (rounded down to the nearest interval and excluding)
//	end - end of chart block (rounded down to the nearest interval and excluding)
//	ticks - number of ticks needed, this field is optional, please read the description above
//
//	Limitations: there are limitations in charts data availability. Detailed ranges for charts data, what can be accessed
//	with specific period, are as follows:
//	-
func GetChartRangeRequest(conn connection, symbol string, period, start, end, ticks int, opts ...model.GetChartRangeRequestOption) (*model.GetChartRangeResponse, error) {
	return sendReceive[model.GetChartRangeResponse](conn, model.NewGetChartRangeRequest(symbol, period, start, end, ticks, opts...))
}

// GetCommissionDef returns calculation of commission and rate of exchange. The value is calculated as expected value,
// and therefore might not be perfectly accurate
//
//	conn - connection to the servier
//	symbol - symbol
//	volume - volume
//	opts - options for the request
//
// The allowed options are:
//   - customTag - custom tag
//   - prettyPrint - pretty print flag
func GetCommissionDef(conn connection, symbol string, volume float64, opts ...model.GetCommissionDefRequestOption) (*model.GetCommissionDefResponse, error) {
	return sendReceive[model.GetCommissionDefResponse](conn,
		model.NewGetCommissionDefRequest(symbol, volume, opts...))
}

// GetCurrentUserData returns information about account currency, and account leverage
//
//	conn - connection to the servier
//	opts - options for the request
//
// The allowed options are:
//   - customTag - custom tag
//   - prettyPrint - pretty print flag
func GetCurrentUserData(conn connection, opts ...model.GetCurrentUserDataRequestOption) (*model.GetCurrentUserDataResponse, error) {
	return sendReceive[model.GetCurrentUserDataResponse](conn, model.NewGetCurrentUserDataRequest(opts...))
}

// GetIbsHistory returns IBs data from the given time range
//
//	conn - connection to the servier
//	start - start of IBs history block
//	end - end of IBs history block
//	opts - options for the request
//
// The allowed options are:
//   - customTag - custom tag
//   - prettyPrint - pretty print flag
func GetIbsHistory(conn connection, start, end int, opts ...model.GetIbsHistoryRequestOption) (*model.GetCurrentUserDataResponse, error) {
	return sendReceive[model.GetCurrentUserDataResponse](conn, model.NewGetIbsHistoryRequest(start, end, opts...))
}

// GetMarginLevel returns various account indicators
//
//	conn - connection to the servier
//	opts - options for the request
//
// The allowed options are:
//   - customTag - custom tag
//   - prettyPrint - pretty print flag
func GetMarginLevel(conn connection, opts ...model.GetMarginLevelRequestOption) (*model.GetMarginLevelResponse, error) {
	return sendReceive[model.GetMarginLevelResponse](conn, model.NewGetMarginLevelRequest(opts...))
}

// GetMarginTrade returns expected margin for given instrument and volume. The value is calculated as expected margin value,
// and therefore might not be perfectly accurate
//
//	conn - connection to the servier
//	symbol - symbol
//	volume - volume
//	opts - options for the request
//
// The allowed options are:
//   - customTag - custom tag
//   - prettyPrint - pretty print flag
func GetMarginTrade(conn connection, symbol string, volume float64, opts ...model.GetMarginTradeRequestOption) (*model.GetMarginTradeResponse, error) {
	return sendReceive[model.GetMarginTradeResponse](conn, model.NewGetMarginTradeRequest(symbol, volume, opts...))
}

// GetNews returns news from trading server which were sent within specified period of time
//
//	c - connection
//	start - start time
//	end - end time, 0 means current time for simplicity
//	opts - options for the request
//
// The allowed options are:
//   - customTag - custom tag
//   - prettyPrint - pretty print flag
func GetNews(conn connection, start, end int) (*model.GetNewsResponse, error) {
	return sendReceive[model.GetNewsResponse](conn, model.NewGetNewsRequest(start, end))
}

// GetProfitCalculation calculates estimated profit for given deal data Should be used for calculator-like apps only.
// Profit for opened transactions should be taken from server, due to higher precision of server calculation
//
//	c - connection
//	symbol - symbol
//	cmd - operation code
//	openPrice - theoretical open price
//	closePrice - theoretical close price
//	volume - volume
//	opts - options for the request
//
// The allowed options are:
//   - customTag - custom tag
//   - prettyPrint - pretty print flag
func GetProfitCalculation(conn connection, symbol string, cmd int, openPrice, closePrice, volume float64, opts ...model.GetProfitCalculationRequestOption) (*model.GetProfitCalculationResponse, error) {
	return sendReceive[model.GetProfitCalculationResponse](conn, model.NewGetProfitCalculationRequest(symbol, cmd, openPrice, closePrice, volume, opts...))
}

// GetServerTime returns current time on trading server
//
//	c - connection
//	opts - options for the request
//
// The allowed options are:
//   - customTag - custom tag
//   - prettyPrint - pretty print flag
func GetServerTime(conn connection, opts ...model.GetServerTimeRequestOption) (*model.GetServerTimeResponse, error) {
	return sendReceive[model.GetServerTimeResponse](conn, model.NewGetServerTimeRequest(opts...))
}

// GetStepRules returns a list of step rules for DMAs
//
//	c - connection
//	opts - options for the request
//
// The allowed options are:
//   - customTag - custom tag
//   - prettyPrint - pretty print flag
func GetStepRules(conn connection, opts ...model.GetStepRulesRequestOption) (*model.GetStepRulesResponse, error) {
	return sendReceive[model.GetStepRulesResponse](conn, model.NewGetStepRulesRequest(opts...))
}

// GetSymbol returns information about symbol available for the user.
//
//	c - connection to the servier
//	symbol - symbol
//	opts - options for the request
//
// The allowed options are:
//   - customTag - custom tag
//   - prettyPrint - pretty print flag
func GetSymbol(conn connection, symbol string, opts ...model.GetSymbolRequestOption) (*model.GetSymbolResponse, error) {
	return sendReceive[model.GetSymbolResponse](conn, model.NewGetSymbolRequest(symbol, opts...))
}

// GetTickPrices returns array of current quotations for given symbols, only quotations that changed from given timestamp
// are returned. New timestamp obtained from output will be used as an argument of the next call of this command
//
//	conn - connection to the servier
//	level - price level
//	timestamp - The time from which the most recent tick should be looked for. Historical prices cannot be obtained using
//	this parameter. It can only be used to verify whether a price has changed since the given time
//	opts - options for the request
//
// The allowed options are:
//   - customTag - custom tag
//   - prettyPrint - pretty print flag
//   - symbol - symbol (one or more)
//
// Note: This function can be usually replaced by its streaming equivalent getTickPrices  which is the preferred way of retrieving ticks data
func GetTickPrices(conn connection, level int, timestamp int, opts ...model.GetTickPricesRequestOption) (*model.GetTickPricesResponse, error) {
	return sendReceive[model.GetTickPricesResponse](conn, model.NewGetTickPricesRequest(level, timestamp, opts...))
}

// GetTradeRecords Returns array of trades listed in orders argument.
//
//	conn - connection to the servier
//	opts - options for the request
//
// The allowed options are:
//   - customTag - custom tag
//   - prettyPrint - pretty print flag
//   - order - order (one or more)
func GetTradeRecords(conn connection, opts ...model.GetTradeRecordsRequestOption) (*model.GetTradeRecordsResponse, error) {
	return sendReceive[model.GetTradeRecordsResponse](conn, model.NewGetTradeRecordsRequest(opts...))
}

// GetTrades returns array of user's trades
//
//	conn - connection to the servier
//	opts - options for the request
//
// The allowed options are:
//   - customTag - custom tag
//   - prettyPrint - pretty print flag
//   - openedOnly - opened trades only flag
//
// Note that this function can be usually replaced by its streaming equivalent getTrades
func GetTrades(conn connection, opts ...model.GetTradesRequestOption) (*model.GetTradesResponse, error) {
	return sendReceive[model.GetTradesResponse](conn, model.NewGetTradesRequest(opts...))
}

// GetTradesHistory returns array of user's trades which were closed within specified period of time
//
//	conn - connection to the servier
//	start - Time, 0 means last month interval
//	end - Time, 0 means current time for simplicity
//	opts - options for the request
//
// The allowed options are:
//   - customTag - custom tag
//   - prettyPrint - pretty print flag
func GetTradesHistory(conn connection, start, end int, opts ...model.GetTradesHistoryRequestOption) (*model.GetTradesHistoryResponse, error) {
	return sendReceive[model.GetTradesHistoryResponse](conn, model.NewGetTradesHistoryRequest(start, end, opts...))
}

// GetTradingHours returns quotes and trading times
//
//	conn - connection to the servier
//	opts - options for the request
//
// The allowed options are:
//   - model.GetTradingHoursRequestWithCustomTag - custom tag
//   - model.GetTradingHoursRequestWithPrettyPrint - pretty print flag
//   - model.GetTradingHoursRequestWithSymbol - symbol (one or more)
func GetTradingHours(conn connection, opts ...model.GetTradingHoursRequestOption) (*model.GetTradingHoursResponse, error) {
	return sendReceive[model.GetTradingHoursResponse](conn, model.NewGetTradingHoursRequest(opts...))
}

// GetVersion returns the current API version
//
//	conn - connection to the servier
//	opts - options for the request
//
// The allowed options are:
//   - model.GetVersionRequestWithCustomTag - custom tag
//   - model.GetVersionRequestWithPrettyPrint - pretty print flag
func GetVersion(conn connection, opts ...model.GetVersionRequestOption) (*model.GetVersionResponse, error) {
	return sendReceive[model.GetVersionResponse](conn, model.NewGetVersionRequest(opts...))
}

// Ping pings the servier.
//
//	conn - connection to the servier
//	opts - options for the request
//
// The allowed options are:
//   - model.PingRequestWithCustomTag - custom tag
//   - model.PingRequestWithPrettyPrint - pretty print flag
//
// Regularly calling this function is enough to refresh the internal state of all the components in the system.
// It is recommended that any application that does not execute other commands, should call this command at least once every 10 minutes.
// Please note that the streaming counterpart of this function is combination of ping  and getKeepAlive .
// Request:
func Ping(conn connection, opts ...model.PingRequestOption) (*model.PingResponse, error) {
	return sendReceive[model.PingResponse](conn, model.NewPingRequest(opts...))
}

// TradeTransaction starts trade transaction. tradeTransaction sends main transaction information to the server.
//
//	conn - connection to the servier
//	symbol - symbol
//	operationCode - operation code
//	operationType - operation type
//	price - trade price
//	volume - trade volume
//	opts - options for the request
//
// The allowed operationCode:
//   - common.OperationCodeBuy 	- buy
//   - common.OperationCodeSell - sell
//   - common.OperationCodeBuyLimit - buy limit
//   - common.OperationCodeSellLimit - sell limit
//   - common.OperationCodeBuyStop 	- buy stop
//   - common.OperationCodeSellStop - sell stop
//
// The allowed operation types:
//   - common.OperationTypeOpen - order open
//   - common.OperationTypeClose - order close
//
// The allowed options are:
//   - model.TradeTransactionRequestWithCustomTag - custom tag
//   - model.TradeTransactionRequestWithCustomComment - The value the customer may provide in order to retrieve it later
//   - model.TradeTransactionRequestWithExpiration - Pending order expiration time
//   - model.TradeTransactionRequestWithPrettyPrint - pretty print flag
//   - model.TradeTransactionRequestWithOffset - trailing offset
//   - model.TradeTransactionRequestWithOrder - position number for closing/modifications
//   - model.TradeTransactionRequestWithStopLoss - stop loss
//   - model.TradeTransactionRequestWithTakeProfit - take profit
func TradeTransaction(conn connection, symbol string, operationCode, operationType int, price, volume float64, opts ...model.TradeTransactionRequestOption) (response *model.TradeTransactionResponse, err error) {
	var (
		request *model.TradeTransactionRequest
	)
	if request, err = model.NewTradeTransactionRequest(symbol, operationCode, operationType, price, volume, opts...); err != nil {
		return
	}
	return sendReceive[model.TradeTransactionResponse](conn, request)
}

// TradeTransactionStatus returns current transaction status. At any time of transaction processing client might check
// the status of transaction on server side. In order to do that client must provide unique order taken from tradeTransaction  invocation
//
//	conn - connection to the servier
//	order - order
//	opts - options for the request
//
// The allowed options are:
//   - model.TradeTransactionStatusRequestWithCustomTag - custom tag
//   - model.TradeTransactionStatusRequestWithPrettyPrint - pretty print flag
func TradeTransactionStatus(conn connection, order int, opts ...model.TradeTransactionStatusRequestOption) (*model.TradeTransactionStatusResponse, error) {
	return sendReceive[model.TradeTransactionStatusResponse](conn, model.NewTradeTransactionStatusRequest(order, opts...))
}

func GetBalanceStream(conn connection, streamSessionId string) error {
	return conn.Send(stream.NewGetBalanceRequest(streamSessionId))
}

func StopBalanceStream(conn connection) error {
	return conn.Send(stream.NewStopBalanceRequest())
}

func sendReceive[ResponseType any](conn connection, request any) (response *ResponseType, err error) {

	if err = conn.Send(request); err != nil {
		return
	}

	response = new(ResponseType)

	err = conn.Receive(response)

	return
}
