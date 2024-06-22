package model

import (
	"encoding/json"
)

type Request[ArgumentsType any] struct {
	Command     string        `json:"command"`
	CustomTag   string        `json:"customTag,omitempty"`
	Arguments   ArgumentsType `json:"arguments,omitempty"`
	PrettyPrint bool          `json:"prettyPrint,omitempty"`
}

func (r *Request[ArgumentsType]) String() string {
	buf, _ := json.Marshal(r)
	return string(buf)
}

type Response[ReturnDataType any] struct {
	Status          bool            `json:"status"`
	CustomTag       string          `json:"customTag,omitempty"`
	ErrorCode       string          `json:"errorCode,omitempty"`
	ErrorDescr      string          `json:"errorDescr,omitempty"`
	ReturnData      ReturnDataType  `json:"returnData,omitempty"`
	StreamSessionId string          `json:"streamSessionId,omitempty"`
	Redirect        *RedirectRecord `json:"redirect,omitempty"`
}

func (r *Response[ReturnDataType]) String() string {
	buf, _ := json.Marshal(r)
	return string(buf)
}

type LoginRequest Request[LoginArguments]
type LoginResponse struct {
	Status          bool            `json:"status"`
	CustomTag       string          `json:"customTag,omitempty"`
	ErrorCode       string          `json:"errorCode,omitempty"`
	ErrorDescr      string          `json:"errorDescr,omitempty"`
	StreamSessionId string          `json:"streamSessionId,omitempty"`
	Redirect        *RedirectRecord `json:"redirect,omitempty"`
}
type LoginRequestOption func(request *LoginRequest)

func NewLoginRequest(userId, password string, opts ...LoginRequestOption) (r *LoginRequest) {
	r = &LoginRequest{
		Command: CmdLogin,
		Arguments: LoginArguments{
			UserId:   userId,
			Password: password},
	}
	for _, o := range opts {
		o(r)
	}
	return
}

func LoginRequestWithCustomTag(tag string) LoginRequestOption {
	return func(request *LoginRequest) {
		request.CustomTag = tag
	}
}

func LoginRequestWithPrettyPrint(flag bool) LoginRequestOption {
	return func(request *LoginRequest) {
		request.PrettyPrint = flag
	}
}

func LoginRequestWithAppId(appId string) LoginRequestOption {
	return func(request *LoginRequest) {
		request.Arguments.AppId = appId
	}
}

func LoginRequestWithAppName(appName string) LoginRequestOption {
	return func(request *LoginRequest) {
		request.Arguments.AppName = appName
	}
}

type LogoutRequest Request[Nil]
type LogoutResponse Response[Nil]
type LogoutRequestOption func(request *LogoutRequest)

func LogoutRequestWithCustomTag(tag string) LogoutRequestOption {
	return func(request *LogoutRequest) {
		request.CustomTag = tag
	}
}

func LogoutRequestWithPrettyPrint(flag bool) LogoutRequestOption {
	return func(request *LogoutRequest) {
		request.PrettyPrint = flag
	}
}

func NewLogoutRequest(opts ...LogoutRequestOption) (r *LogoutRequest) {
	r = &LogoutRequest{
		Command: CmdLogout,
	}
	for _, o := range opts {
		o(r)
	}
	return
}

type GetAllSymbolsRequest Request[Nil]
type GetAllSymbolsResponse Response[[]SymbolRecord]
type GetAllSymbolsRequestOption func(request *GetAllSymbolsRequest)

func NewGetAllSymbolsRequest(opts ...GetAllSymbolsRequestOption) (r *GetAllSymbolsRequest) {
	r = &GetAllSymbolsRequest{
		Command: CmdGetAllSymbols,
	}
	for _, o := range opts {
		o(r)
	}
	return
}

func GetAllSymbolsRequestWithCustomTag(tag string) GetAllSymbolsRequestOption {
	return func(request *GetAllSymbolsRequest) {
		request.CustomTag = tag
	}
}

func GetAllSymbolsRequestWithPrettyPrint(flag bool) GetAllSymbolsRequestOption {
	return func(request *GetAllSymbolsRequest) {
		request.PrettyPrint = flag
	}
}

type GetCalendarRequest Request[Nil]
type GetCalendarResponse Response[[]CalendarRecord]
type GetCalendarRequestOption func(response *GetCalendarRequest)

func GetCalendarRequestWithCustomTag(tag string) GetCalendarRequestOption {
	return func(request *GetCalendarRequest) {
		request.CustomTag = tag
	}
}

func GetCalendarRequestWithPrettyPrint(flag bool) GetCalendarRequestOption {
	return func(request *GetCalendarRequest) {
		request.PrettyPrint = flag
	}
}

func NewGetCalendarRequest(opts ...GetCalendarRequestOption) (r *GetCalendarRequest) {
	r = &GetCalendarRequest{
		Command: CmdGetCalendar,
	}
	for _, o := range opts {
		o(r)
	}
	return
}

type CalendarRecord struct {
	Country  string `json:"country"`
	Current  string `json:"current"`
	Forecast string `json:"forecast"`
	Impact   string `json:"impact"`
	Period   string `json:"period"`
	Previous string `json:"previous"`
	Time     int    `json:"time"`
	Title    string `json:"title"`
}

type ChartLastInfoRecord struct {
	Period int    `json:"period"`
	Start  int    `json:"start"`
	Symbol string `json:"symbol"`
}

type RateInfoRecord struct {
	Close     float64 `json:"close"`
	Ctm       int     `json:"ctm"`
	CtmString string  `json:"ctmString"`
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Open      float64 `json:"open"`
	Vol       float64 `json:"vol"`
}

type RateInfoData struct {
	Digits    int              `json:"digits"`
	RateInfos []RateInfoRecord `json:"rateInfos"`
}

type GetChartLastRequest Request[ChartLastInfoRecord]
type GetChartLastResponse Response[RateInfoData]
type GetChartLastRequestOption func(request *GetChartLastRequest)

func GetChartLastRequestWithCustomTAg(tag string) GetChartLastRequestOption {
	return func(request *GetChartLastRequest) {
		request.CustomTag = tag
	}
}

func GetChartLastRequestWithPrettyPrint(flag bool) GetChartLastRequestOption {
	return func(request *GetChartLastRequest) {
		request.PrettyPrint = true
	}
}

func NewGetChartLastRequest(symbol string, period int, start int, opts ...GetChartLastRequestOption) (r *GetChartLastRequest) {
	r = &GetChartLastRequest{
		Command: CmdGetChartLastRequest,
		Arguments: ChartLastInfoRecord{
			Period: period,
			Start:  start,
			Symbol: symbol,
		},
	}
	for _, o := range opts {
		o(r)
	}
	return
}

type GetChartRangeRequest Request[ChartRangeInfoRecord]
type GetChartRangeResponse Response[RateInfoData]
type GetChartRangeRequestOption func(*GetChartRangeRequest)

func GetChartRangeRequestWithPrettyPrint(flag bool) GetChartRangeRequestOption {
	return func(request *GetChartRangeRequest) {
		request.PrettyPrint = flag
	}
}

func GetChartRangeRequestWithCustomTag(tag string) GetChartRangeRequestOption {
	return func(request *GetChartRangeRequest) {
		request.CustomTag = tag
	}
}

func NewGetChartRangeRequest(symbol string, period, start, end, ticks int, opts ...GetChartRangeRequestOption) (r *GetChartRangeRequest) {
	r = &GetChartRangeRequest{
		Command: CmdGetChartRangeRequest,
		Arguments: ChartRangeInfoRecord{
			End:    end,
			Period: period,
			Start:  start,
			Symbol: symbol,
			Ticks:  ticks,
		},
	}
	for _, o := range opts {
		o(r)
	}
	return
}

type ChartRangeInfoRecord struct {
	End    int    `json:"end"`
	Period int    `json:"period"`
	Start  int    `json:"start"`
	Symbol string `json:"symbol"`
	Ticks  int    `json:"ticks"`
}

type GetCommissionDefArgs struct {
	Symbol string  `json:"symbol"`
	Volume float64 `json:"volume"`
}

type GetCommissionDefData struct {
	Commission     float64 `json:"commission"`
	RateOfExchange float64 `json:"rateOfExchange"`
}

type GetCommissionDefRequest Request[GetCommissionDefArgs]
type GetCommissionDefResponse Response[GetCommissionDefData]
type GetCommissionDefRequestOption func(request *GetCommissionDefRequest)

func GetCommissionDefWithPrettyPrint(flag bool) GetCommissionDefRequestOption {
	return func(request *GetCommissionDefRequest) {
		request.PrettyPrint = flag
	}
}

func GetCommissionDefWithCustomTag(tag string) GetCommissionDefRequestOption {
	return func(request *GetCommissionDefRequest) {
		request.CustomTag = tag
	}
}

func NewGetCommissionDefRequest(symbol string, volume float64, opts ...GetCommissionDefRequestOption) (r *GetCommissionDefRequest) {
	r = &GetCommissionDefRequest{
		Command: CmdGetCommissionDef,
		Arguments: GetCommissionDefArgs{
			Symbol: symbol,
			Volume: volume,
		},
	}
	for _, o := range opts {
		o(r)
	}
	return
}

type CurrentUserData struct {
	CompanyUnit        int     `json:"companyUnit"`
	Currency           string  `json:"currency"`
	Group              string  `json:"group"`
	IbAccount          bool    `json:"ibAccount"`
	Leverage           int     `json:"leverage"`
	LeverageMultiplier float64 `json:"leverageMultiplier"`
	SpreadType         *string `json:"spreadType"`
	TrailingStop       bool    `json:"trailingStop"`
}

type GetCurrentUserDataRequest Request[Nil]
type GetCurrentUserDataResponse Response[CurrentUserData]
type GetCurrentUserDataRequestOption func(request *GetCurrentUserDataRequest)

func GetCurrentUserDataRequestWithCustomTag(tag string) GetCurrentUserDataRequestOption {
	return func(request *GetCurrentUserDataRequest) {
		request.CustomTag = tag
	}
}

func GetCurrentUserDataRequestWithPrettyPrint(flag bool) GetCurrentUserDataRequestOption {
	return func(request *GetCurrentUserDataRequest) {
		request.PrettyPrint = flag
	}
}

func NewGetCurrentUserDataRequest(opts ...GetCurrentUserDataRequestOption) (r *GetCurrentUserDataRequest) {
	r = &GetCurrentUserDataRequest{
		Command: CmdGetCurrentUserData,
	}
	for _, o := range opts {
		o(r)
	}
	return
}

type IbsHistoryArgs struct {
	End   int `json:"end"`
	Start int `json:"start"`
}

type IbRecord struct {
	ClosePrice *float64 `json:"closePrice,omitempty"`
	Login      *string  `json:"login,omitempty"`
	Nominal    *float64 `json:"nominal,omitempty"`
	OpenPrice  *float64 `json:"openPrice,omitempty"`
	Side       *int     `json:"side,omitempty"`
	Surname    *string  `json:"surname,omitempty"`
	Symbol     *string  `json:"symbol,omitempty"`
	Timestamp  *int     `json:"timestamp,omitempty"`
	Volume     *float64 `json:"volume,omitempty"`
}

type GetIbsHistoryRequest Request[IbsHistoryArgs]
type GetIbsHistoryRequestOption func(request *GetIbsHistoryRequest)
type GetIbsHistoryResponse Response[[]IbRecord]

func GetIbsHistoryRequestWithCustomTag(tag string) GetIbsHistoryRequestOption {
	return func(request *GetIbsHistoryRequest) {
		request.CustomTag = tag
	}
}

func GetIbsHistoryRequestWithPrettyPrint(flag bool) GetIbsHistoryRequestOption {
	return func(request *GetIbsHistoryRequest) {
		request.PrettyPrint = flag
	}
}

func NewGetIbsHistoryRequest(start, end int, opts ...GetIbsHistoryRequestOption) (r *GetIbsHistoryRequest) {
	r = &GetIbsHistoryRequest{
		Command: CmdGetIbsHistory,
		Arguments: IbsHistoryArgs{
			End:   end,
			Start: start,
		},
	}
	for _, o := range opts {
		o(r)
	}
	return
}

type GetMarginLevelData struct {
	Balance     float64 `json:"balance"`
	Credit      float64 `json:"credit"`
	Currency    string  `json:"currency"`
	Equity      float64 `json:"equity"`
	Margin      float64 `json:"margin"`
	MarginFree  float64 `json:"margin_free"`
	MarginLevel float64 `json:"margin_level"`
}

type GetMarginLevelRequest Request[Nil]
type GetMarginLevelRequestOption func(request *GetMarginLevelRequest)
type GetMarginLevelResponse Response[GetMarginLevelData]

func GetMarginLevelRequestWithCustomTag(tag string) GetMarginLevelRequestOption {
	return func(request *GetMarginLevelRequest) {
		request.CustomTag = tag
	}
}

func GetMarginLevelRequestWithPrettyPrint(flag bool) GetMarginLevelRequestOption {
	return func(request *GetMarginLevelRequest) {
		request.PrettyPrint = flag
	}
}

func NewGetMarginLevelRequest(opts ...GetMarginLevelRequestOption) (r *GetMarginLevelRequest) {
	r = &GetMarginLevelRequest{
		Command: CmdGetMarginLevel,
	}
	for _, o := range opts {
		o(r)
	}
	return
}

type GetMarginTradeArgs struct {
	Symbol string  `json:"symbol"`
	Volume float64 `json:"volume"`
}

type GetMarginTradeData struct {
	Margin float64 `json:"margin"`
}

type GetMarginTradeRequest Request[GetMarginTradeArgs]
type GetMarginTradeRequestOption func(*GetMarginTradeRequest)
type GetMarginTradeResponse Response[GetMarginTradeData]

func GetMarginTradeRequestWithCustomTag(tag string) GetMarginTradeRequestOption {
	return func(request *GetMarginTradeRequest) {
		request.CustomTag = tag
	}
}

func GetMarginTradeRequestWithPrettyPrint(flag bool) GetMarginTradeRequestOption {
	return func(request *GetMarginTradeRequest) {
		request.PrettyPrint = flag
	}
}

func NewGetMarginTradeRequest(symbol string, volume float64, opts ...GetMarginTradeRequestOption) (r *GetMarginTradeRequest) {
	r = &GetMarginTradeRequest{
		Command: CmdGetMarginTrade,
		Arguments: GetMarginTradeArgs{
			Symbol: symbol,
			Volume: volume,
		},
	}

	for _, o := range opts {
		o(r)
	}
	return
}

type PingRequest struct {
	Request[Nil]
}

type PingResponse Response[Nil]
type PingRequestOption func(request *PingRequest)

func PingRequestWithCustomTag(tag string) PingRequestOption {
	return func(request *PingRequest) {
		request.CustomTag = tag
	}
}

func PingRequestWithPrettyPrint(flag bool) PingRequestOption {
	return func(request *PingRequest) {
		request.PrettyPrint = flag
	}
}

func NewPingRequest(opts ...PingRequestOption) (r *PingRequest) {
	r = &PingRequest{
		Request: Request[Nil]{
			Command: CmdPing,
		},
	}
	for _, o := range opts {
		o(r)
	}
	return
}

type ServerTimeRecord struct {
	Time       int    `json:"time"`
	TimeString string `json:"timeString"`
}

type GetServerTimeRequest Request[Nil]
type GetServerTimeRequestOption func(request *GetServerTimeRequest)
type GetServerTimeResponse Response[ServerTimeRecord]

func NewGetServerTimeRequest(opts ...GetServerTimeRequestOption) (r *GetServerTimeRequest) {
	r = &GetServerTimeRequest{
		Command: CmdGetServerTime,
	}
	for _, o := range opts {
		o(r)
	}
	return
}

func GetServerTimeRequestWithCustomTag(tag string) GetServerTimeRequestOption {
	return func(request *GetServerTimeRequest) {
		request.CustomTag = tag
	}
}

func GetServerTimeRequestWithPrettyPrint(flag bool) GetServerTimeRequestOption {
	return func(request *GetServerTimeRequest) {
		request.PrettyPrint = flag
	}
}

type GetStepRulesRequest Request[Nil]
type GetStepRulesRequestOption func(rules *GetStepRulesRequest)
type GetStepRulesResponse Response[StepRuleRecords]

func NewGetStepRulesRequest(opts ...GetStepRulesRequestOption) (r *GetStepRulesRequest) {
	r = &GetStepRulesRequest{
		Command: CmdGetStepRules,
	}

	for _, o := range opts {
		o(r)
	}
	return r
}

func GetStepRulesRequestWithCustomTag(tag string) GetStepRulesRequestOption {
	return func(rules *GetStepRulesRequest) {
		rules.CustomTag = tag
	}
}

func GetStepRulesRequestWithPrettyPrint(flag bool) GetStepRulesRequestOption {
	return func(rules *GetStepRulesRequest) {
		rules.PrettyPrint = flag
	}
}

type Nil any
type ReturnData struct {
	SymbolRecords []*SymbolRecord
}

type LoginArguments struct {
	UserId   string `json:"userId"`
	Password string `json:"password"`
	AppId    string `json:"appId,omitempty"`
	AppName  string `json:"appName,omitempty"`
}

type RedirectRecord struct {
	MainPort      int    `json:"mainPort"`
	StreamingPort int    `json:"streamingPort"`
	Address       string `json:"address"`
}

type SymbolRecord struct {
	Ask                float64 `json:"ask"`
	Bid                float64 `json:"bid"`
	CategoryName       string  `json:"categoryName"`
	ContractSize       int     `json:"contractSize"`
	Currency           string  `json:"currency"`
	CurrencyPair       bool    `json:"currencyPair"`
	CurrencyProfit     string  `json:"currencyProfit"`
	Description        string  `json:"description"`
	Expiration         *int    `json:"expiration"`
	GroupName          string  `json:"groupName"`
	High               float64 `json:"high"`
	InitialMargin      int     `json:"initialMargin"`
	InstantMaxVolume   int     `json:"instantMaxVolume"`
	Leverage           float64 `json:"leverage"`
	LongOnly           bool    `json:"longOnly"`
	LotMax             float64 `json:"lotMax"`
	LotMin             float64 `json:"lotMin"`
	LotStep            float64 `json:"lotStep"`
	Low                float64 `json:"low"`
	MarginHedged       int     `json:"marginHedged"`
	MarginHedgedStrong bool    `json:"marginHedgedStrong"`
	MarginMaintenance  *int    `json:"marginMaintenance"`
	MarginMode         int     `json:"marginMode"`
	Percentage         float64 `json:"percentage"`
	PipsPrecision      int     `json:"pipsPrecision"`
	Precision          int     `json:"precision"`
	ProfitMode         int     `json:"profitMode"`
	QuoteId            int     `json:"quoteId"`
	ShortSelling       bool    `json:"shortSelling"`
	SpreadRaw          float64 `json:"spreadRaw"`
	SpreadTable        float64 `json:"spreadTable"`
	Starting           *int    `json:"starting"`
	StepRuleId         int     `json:"stepRuleId"`
	StopsLevel         int     `json:"stopsLevel"`
	SwapRollover3Days  int     `json:"swap_rollover3days"`
	SwapEnable         bool    `json:"swapEnable"`
	SwapLong           float64 `json:"swapLong"`
	SwapShort          float64 `json:"swapShort"`
	SwapType           int     `json:"swapType"`
	Symbol             string  `json:"symbol"`
	TickSize           float64 `json:"tickSize"`
	TickValue          float64 `json:"tickValue"`
	Time               int     `json:"time"`
	TimeString         string  `json:"timeString"`
	TrailingEnabled    bool    `json:"trailingEnabled"`
	Type               int     `json:"type"`
}

type NewsRequestArg struct {
	End   int `json:"end"`
	Start int `json:"start"`
}

type NewsTopicRecord struct {
	Body       string `json:"body"`
	BodyLength int    `json:"bodylen"`
	Key        string `json:"key"`
	Time       int    `json:"time"`
	TimeString string `json:"timeString"`
	Title      string `json:"title"`
}

type NewsTopicRecords []NewsTopicRecord
type GetNewsRequest Request[NewsRequestArg]
type GetNewsRequestOption func(request *GetNewsRequest)
type GetNewsResponse Response[NewsTopicRecords]

func NewGetNewsRequest(start, end int, opts ...GetNewsRequestOption) (r *GetNewsRequest) {
	r = &GetNewsRequest{
		Command:   CmdGetNews,
		Arguments: NewsRequestArg{End: end, Start: start},
	}
	for _, o := range opts {
		o(r)
	}
	return r
}

type GetProfitCalculationArgs struct {
	ClosePrice float64 `json:"closePrice"`
	Cmd        int     `json:"cmd"`
	OpenPrice  float64 `json:"openPrice"`
	Symbol     string  `json:"symbol"`
	Volume     float64 `json:"volume"`
}

type ProfitRecord struct {
	Profit float64 `json:"profit"`
}

type GetProfitCalculationRequest Request[GetProfitCalculationArgs]
type GetProfitCalculationRequestOption func(*GetProfitCalculationRequest)
type GetProfitCalculationResponse Response[ProfitRecord]

func NewGetProfitCalculationRequest(symbol string, cmd int, openPrice, closePrice, volume float64, opts...GetProfitCalculationRequestOption) (r *GetProfitCalculationRequest) {
	r = &GetProfitCalculationRequest{
		Command: CmdGetProfitCalculation,
		Arguments: GetProfitCalculationArgs{
			ClosePrice: closePrice,
			Cmd:        cmd,
			OpenPrice:  openPrice,
			Symbol:     symbol,
			Volume:     volume,
		},
	}

	for _, o := range opts {
		o(r)
	}

	return
}

func GetProfitCalculationRequestWithCustomTag(tag string) GetProfitCalculationRequestOption {
	return func(request *GetProfitCalculationRequest) {
		request.CustomTag = tag
	}
}

func GetProfitCalculationRequestWithPrettyPrint(flag bool) GetProfitCalculationRequestOption {
	return func(request *GetProfitCalculationRequest) {
		request.PrettyPrint = flag
	}
}

type StepRecord struct {
	// lower border of the volume range
	FromValue float64 `json:"fromValue"`
	//lot step value in the given volume range
	Step float64 `json:"step"`
}

type StepRuleRecord struct {
	// step rule ID
	Id int `json:"id"`
	// step rule name
	Name string `json:"name"`
	// array of steps
	Steps []StepRecord `json:"steps"`
}

type StepRuleRecords []StepRuleRecord

type GetSymbolRequestArgs struct {
	// symbol
	Symbol string `json:"symbol"`
}

type GetSymbolRequest Request[GetSymbolRequestArgs]
type GetSymbolRequestOption func(request *GetSymbolRequest)
type GetSymbolResponse Response[SymbolRecord]

func NewGetSymbolRequest(symbol string, opts ...GetSymbolRequestOption) (r *GetSymbolRequest) {
	r = &GetSymbolRequest{
		Command: CmdGetSymbol,
		Arguments: GetSymbolRequestArgs{
			Symbol: symbol,
		},
	}
	for _, o := range opts {
		o(r)
	}
	return
}

func GetSymbolRequestWithCustomTag(tag string) GetSymbolRequestOption {
	return func(request *GetSymbolRequest) {
		request.CustomTag = tag
	}
}

func GetSymbolRequestWithPrettyPrint(flag bool) GetSymbolRequestOption {
	return func(request *GetSymbolRequest) {
		request.PrettyPrint = true
	}
}

type GetTickPricesRequestArgs struct {
	// price level
	Level int `json:"level"`
	// Array of symbol names
	Symbols []string `json:"symbols"`
	// The time from which the most recent tick should be looked for. Historical prices cannot be obtained using this parameter.
	// It can only be used to verify whether a price has changed since the given time
	Timestamp int `json:"timestamp"`
}

type TickRecord struct {
	//Ask price in base currency
	Ask float64 `json:"ask"`
	//Number of available lots to buy at given price or nil if not applicable
	AskVolume *int `json:"askVolume"`
	// Bid price in base currency
	Bid float64 `json:"bid"`
	// Number of available lots to sell at given price or nil if not applicable
	BidVolume *int `json:"bidVolume"`
	// The highest price of the day in base currency
	High float64 `json:"high"`
	// Price level
	Level int `json:"level"`
	// The lowest price of the day in base currency
	Low float64 `json:"low"`
	// The difference between raw ask and bid prices
	SpreadRaw float64 `json:"spreadRaw"`
	// Spread representation
	SpreadTable float64 `json:"spreadTable"`
	// Symbol
	Symbol string `json:"symbol"`
	// Timestamp
	Timestamp int `json:"timestamp"`
}

type GetTickPricesData struct {
	Quotations []TickRecord `json:"quotations"`
}

type GetTickPricesRequest Request[GetTickPricesRequestArgs]
type GetTickPricesRequestOption func(r *GetTickPricesRequest)
type GetTickPricesResponse Response[GetTickPricesData]

func NewGetTickPricesRequest(level int, timestamp int, opts ...GetTickPricesRequestOption) (r *GetTickPricesRequest) {
	r = &GetTickPricesRequest{
		Command: CmdGetTickPrices,
		Arguments: GetTickPricesRequestArgs{
			Level:     level,
			Timestamp: timestamp,
		},
	}
	for _, o := range opts {
		o(r)
	}
	return
}

func GetTickPricesRequestWithSymbol(symbol string) GetTickPricesRequestOption {
	return func(r *GetTickPricesRequest) {
		r.Arguments.Symbols = append(r.Arguments.Symbols, symbol)
	}
}

func GetTickPricesRequestWithCustomTag(tag string) GetTickPricesRequestOption {
	return func(r *GetTickPricesRequest) {
		r.CustomTag = tag
	}
}

func GetTickPricesRequestWithPrettyPrint(flag bool) GetTickPricesRequestOption {
	return func(r *GetTickPricesRequest) {
		r.PrettyPrint = flag
	}
}

type GetTradeRecordsRequestArgs struct {
	// Array of orders (position numbers)
	Orders []int `json:"orders"`
}

type TradeRecord struct {
	// Close price in base currency
	ClosePrice float64 `json:"close_price"`
	// Null if order is not closed
	CloseTime *int `json:"close_time"`
	// Null if order is not closed
	CloseTimeString *string `json:"close_timeString"`
	// Closed
	Closed bool `json:"closed"`
	// Operation code
	Cmd int `json:"cmd"`
	// Comment
	Comment string `json:"comment"`
	// Commission in account currency, null if not applicable
	Commission float64 `json:"commission"`
	// The value the customer may provide in order to retrieve it later.
	CustomComment string `json:"customComment"`
	// Number of decimal places
	Digits int `json:"digits"`
	// Null if order is not closed
	Expiration *int `json:"expiration"`
	// Null if order is not closed
	ExpirationString *string `json:"expirationString"`
	// Margin rate
	MarginRate float64 `json:"margin_rate"`
	// Trailing offset
	Offset int `json:"offset"`
	// Open price in base currency
	OpenPrice float64 `json:"open_price"`
	// Open time
	OpenTime int `json:"open_time"`
	// Open time string
	OpenTimeString string `json:"open_timeString"`
	// Order number for opened transaction
	Order int `json:"order"`
	// 	Order number for closed transaction
	Order2 int `json:"order2"`
	// Order number common both for opened and closed transaction
	Position int `json:"position"`
	// Profit in account currency
	Profit float64 `json:"profit"`
	// Zero if stop loss is not set (in base currency)
	StopLoss float64 `json:"sl"`
	// order swaps in account currency
	Storage float64 `json:"storage"`
	// symbol name or null for deposit/withdrawal operations
	Symbol string `json:"symbol"`
	// Timestamp
	Timestamp int `json:"timestamp"`
	// Zero if take profit is not set (in base currency)
	TakeProfit float64 `json:"tp"`
	// Volume in lots
	Volume float64 `json:"volume"`
}

type TradeRecords []TradeRecord
type GetTradeRecordsRequest Request[GetTradeRecordsRequestArgs]
type GetTradeRecordsRequestOption func(*GetTradeRecordsRequest)
type GetTradeRecordsResponse Response[TradeRecords]

func NewGetTradeRecordsRequest(opts ...GetTradeRecordsRequestOption) (r *GetTradeRecordsRequest) {
	r = &GetTradeRecordsRequest{
		Command: CmdGetTradeRecords,
	}
	for _, o := range opts {
		o(r)
	}
	return
}

func GetTradeRecordsRequestWithOrder(order int) GetTradeRecordsRequestOption {
	return func(request *GetTradeRecordsRequest) {
		request.Arguments.Orders = append(request.Arguments.Orders, order)
	}
}

func GetTradeRecordsRequestWithCustomTag(tag string) GetTradeRecordsRequestOption {
	return func(request *GetTradeRecordsRequest) {
		request.CustomTag = tag
	}
}

func GetTradeRecordsRequestWithPrettyPrint(flag bool) GetTradeRecordsRequestOption {
	return func(request *GetTradeRecordsRequest) {
		request.PrettyPrint = flag
	}
}

type GetTradesArgs struct {
	// If true then only opened trades will be returned
	OpenedOnly bool `json:"openedOnly"`
}

type GetTradesRequest Request[GetTradesArgs]
type GetTradesRequestOption func(request *GetTradesRequest)
type GetTradesResponse Response[TradeRecords]

func NewGetTradesRequest(opts ...GetTradesRequestOption) (r *GetTradesRequest) {
	r = &GetTradesRequest{
		Command: CmdGetTrades,
	}

	for _, o := range opts {
		o(r)
	}
	return
}

func GetTradesRequestWithOpenedOnly(flag bool) GetTradesRequestOption {
	return func(request *GetTradesRequest) {
		request.Arguments.OpenedOnly = flag
	}
}

func GetTradesRequestWithCustomTag(tag string) GetTradesRequestOption {
	return func(request *GetTradesRequest) {
		request.CustomTag = tag
	}
}

func GetTradesRequestWithPrettyPrint(flag bool) GetTradesRequestOption {
	return func(request *GetTradesRequest) {
		request.PrettyPrint = flag
	}
}

type GetTradesHistoryArgs struct {
	End   int `json:"end"`
	Start int `json:"start"`
}

type GetTradesHistoryRequest Request[GetTradesHistoryArgs]
type GetTradesHistoryRequestOption func(request *GetTradesHistoryRequest)
type GetTradesHistoryResponse Response[TradeRecords]

func NewGetTradesHistoryRequest(start, end int, opts ...GetTradesHistoryRequestOption) (r *GetTradesHistoryRequest) {
	r = &GetTradesHistoryRequest{
		Command: CmdGetTradesHistory,
		Arguments: GetTradesHistoryArgs{
			Start: start,
			End:   end,
		},
	}
	for _, o := range opts {
		o(r)
	}
	return r
}

func GetTradesHistoryRequestWithCustomTag(tag string) GetTradesHistoryRequestOption {
	return func(request *GetTradesHistoryRequest) {
		request.CustomTag = tag
	}
}

func GetTradesHistoryRequestWithPrettyPrint(flag bool) GetTradesHistoryRequestOption {
	return func(request *GetTradesHistoryRequest) {
		request.PrettyPrint = flag
	}
}

type GetTradingHoursArgs struct {
	// Array of symbol names
	Symbols []string `json:"symbols"`
}

type QuoteRecord struct {
	// Day of week
	Day int `json:"day"`
	// Start time in ms from 00:00 CET / CEST time zone (see Daylight Saving Time, DST)
	FromT int `json:"fromT"`
	// End time in ms from 00:00 CET / CEST time zone (see Daylight Saving Time, DST)
	ToT int `json:"toT"`
}

type TradingRecord QuoteRecord

type TradingHoursRecord struct {
	// Array of QuoteRecord
	Quotes []QuoteRecord `json:"quotes"`
	// Symbol
	Symbol string `json:"symbol"`
	// Array of TradingRecord
	Trading []TradingRecord `json:"trading"`
}
type TradingHoursRecords []TradingHoursRecord

type GetTradingHoursRequest Request[GetTradingHoursArgs]
type GetTradingHoursRequestOption func(*GetTradingHoursRequest)
type GetTradingHoursResponse Response[TradingHoursRecords]

func NewGetTradingHoursRequest(opts ...GetTradingHoursRequestOption) (r *GetTradingHoursRequest) {
	r = &GetTradingHoursRequest{
		Command: CmdGetTradingHours,
	}
	for _, o := range opts {
		o(r)
	}
	return
}

func GetTradingHoursRequestWithSymbol(symbol string) GetTradingHoursRequestOption {
	return func(request *GetTradingHoursRequest) {
		request.Arguments.Symbols = append(request.Arguments.Symbols, symbol)
	}
}

func GetTradingHoursRequestWithCustomTag(tag string) GetTradingHoursRequestOption {
	return func(request *GetTradingHoursRequest) {
		request.CustomTag = tag
	}
}

func GetTradingHoursRequestWithPrettyPrint(flag bool) GetTradingHoursRequestOption {
	return func(request *GetTradingHoursRequest) {
		request.PrettyPrint = flag
	}
}

type VersionData struct {
	Version string `json:"version"`
}

type GetVersionRequest Request[Nil]
type GetVersionRequestOption func(request *GetVersionRequest)
type GetVersionResponse Response[VersionData]

func NewGetVersionRequest(opts ...GetVersionRequestOption) (r *GetVersionRequest) {
	r = &GetVersionRequest{
		Command: CmdGetVersion,
	}
	for _, o := range opts {
		o(r)
	}
	return r
}

func GetVersionRequestWithCustomTag(tag string) GetVersionRequestOption {
	return func(request *GetVersionRequest) {
		request.CustomTag = tag
	}
}

func GetVersionRequestWithPrettyPrint(flag bool) GetVersionRequestOption {
	return func(request *GetVersionRequest) {
		request.PrettyPrint = flag
	}
}

type TradeTransInfo struct {
	Cmd           int     `json:"cmd"`
	CustomComment string  `json:"customComment,omitempty"`
	Expiration    int     `json:"expiration"`
	Offset        int     `json:"offset"`
	Order         int     `json:"order"`
	Price         float64 `json:"price"`
	StopLoss      float64 `json:"sl"`
	Symbol        string  `json:"symbol"`
	TakeProfit    float64 `json:"tp"`
	Type          int     `json:"type"`
	Volume        float64 `json:"volume"`
}

type TradeTransactionArgs struct {
	TradeTransInfo `json:"tradeTransInfo"`
}

type TradeTransactionData struct {
	Order int `json:"order"`
}

type TradeTransactionRequest Request[TradeTransactionArgs]
type TradeTransactionRequestOption func(request *TradeTransactionRequest) error
type TradeTransactionResponse Response[TradeTransactionData]

func NewTradeTransactionRequest(symbol string, operationCode, operationType int, price, volume float64, opts ...TradeTransactionRequestOption) (r *TradeTransactionRequest, err error) {
	r = &TradeTransactionRequest{
		Command: CmdTradeTransaction,
		Arguments: TradeTransactionArgs{
			TradeTransInfo{
				Symbol: symbol,
				Cmd:    operationCode,
				Type:   operationType,
				Price:  price,
				Volume: volume,
			},
		},
	}
	for _, o := range opts {
		if err = o(r); err != nil {
			return
		}
	}
	return
}

func TradeTransactionRequestWithCustomTag(tag string) TradeTransactionRequestOption {
	return func(request *TradeTransactionRequest) error {
		request.CustomTag = tag
		return nil
	}
}

func TradeTransactionRequestWithCustomComment(comment string) TradeTransactionRequestOption {
	return func(request *TradeTransactionRequest) error {
		request.Arguments.CustomComment = comment
		return nil
	}
}

func TradeTransactionRequestWithExpiration(expiration int) TradeTransactionRequestOption {
	return func(request *TradeTransactionRequest) error {
		request.Arguments.Expiration = expiration
		return nil
	}
}

func TradeTransactionRequestWithOffset(offset int) TradeTransactionRequestOption {
	return func(request *TradeTransactionRequest) error {
		request.Arguments.Offset = offset
		return nil
	}
}

func TradeTransactionRequestWithOrder(order int) TradeTransactionRequestOption {
	return func(request *TradeTransactionRequest) error {
		request.Arguments.Order = order
		return nil
	}
}

func TradeTransactionRequestWithPrettyPrint(flag bool) TradeTransactionRequestOption {
	return func(request *TradeTransactionRequest) error {
		request.PrettyPrint = flag
		return nil
	}
}

func TradeTransactionRequestWithStopLoss(stopLoss float64) TradeTransactionRequestOption {
	return func(request *TradeTransactionRequest) error {
		request.Arguments.StopLoss = stopLoss
		return nil
	}
}

func TradeTransactionRequestWithTakeProfit(takeProfit float64) TradeTransactionRequestOption {
	return func(request *TradeTransactionRequest) error {
		request.Arguments.TakeProfit = takeProfit
		return nil
	}
}

type TradeTransactionStatusArgs struct {
	Order int `json:"order"`
}

type TradeTransactionStatusData struct {
	Ask           float64 `json:"ask"`
	Bid           float64 `json:"bid"`
	CustomComment string  `json:"customComment,omitempty"`
	Message       *string `json:"message"`
	Order         int     `json:"order"`
	RequestStatus int     `json:"requestStatus"`
}

type TradeTransactionStatusRequest Request[TradeTransactionStatusArgs]
type TradeTransactionStatusRequestOption func(*TradeTransactionStatusRequest)
type TradeTransactionStatusResponse Response[TradeTransactionStatusData]

func NewTradeTransactionStatusRequest(order int, opts ...TradeTransactionStatusRequestOption) (r *TradeTransactionStatusRequest) {
	r = &TradeTransactionStatusRequest{
		Command: CmdTradeTransactionStatus,
		Arguments: TradeTransactionStatusArgs{
			Order: order,
		},
	}
	for _, o := range opts {
		o(r)
	}
	return
}

func TradeTransactionStatusRequestWithCustomTag(tag string) TradeTransactionStatusRequestOption {
	return func(request *TradeTransactionStatusRequest) {
		request.CustomTag = tag
	}
}

func TradeTransactionStatusRequestWithPrettyPrint(flag bool) TradeTransactionStatusRequestOption {
	return func(request *TradeTransactionStatusRequest) {
		request.PrettyPrint = flag
	}
}
