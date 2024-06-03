package model

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

var streamingDataTypeFactory = map[string]func() any{
	DataStreamBalance: func() any {
		return new(StreamingBalanceRecord)
	},
	DataStreamCandle: func() any {
		return new(StreamingCandleRecord)
	},
	DataStreamKeepAlive: func() any {
		return new(StreamingKeepAliveRecord)
	},
	DataStreamNews: func() any {
		return new(StreamingNewsRecord)
	},
	DataStreamProfit: func() any {
		return new(StreamingProfitRecord)
	},
	DataStreamTickPrices: func() any {
		return new(StreamingTickRecord)
	},
	DataStreamTrade: func() any {
		return new(StreamingTradeRecord)
	},
	DataStreamTradeStatus: func() any {
		return new(StreamingTradeStatusRecord)
	},
}

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

type GetBalanceRequest StreamRequest
type StopBalanceRequest StreamRequest

func NewGetBalanceRequest(streamSessionId string) *GetBalanceRequest {
	return &GetBalanceRequest{
		Command:         CmdGetBalance,
		StreamSessionId: streamSessionId,
	}
}

func NewStopBalanceRequest() *StopBalanceRequest {
	return &StopBalanceRequest{
		Command: CmdStopBalance,
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

type GetCandlesRequest struct {
	StreamRequest
	Symbol string `json:"symbol"`
}
type StopCandlesRequest struct {
	StreamRequest
	Symbol string `json:"symbol"`
}

func NewGetCandlesRequest(symbol, streamSessionId string) *GetCandlesRequest {
	return &GetCandlesRequest{
		StreamRequest: StreamRequest{
			Command:         CmdGetCandles,
			StreamSessionId: streamSessionId,
		},
		Symbol: symbol,
	}
}

func NewStopCandlesRequest(symbol string) *StopCandlesRequest {
	return &StopCandlesRequest{
		StreamRequest: StreamRequest{
			Command: CmdStopCandles,
		},
		Symbol: symbol,
	}
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

type GetKeepAliveRequest StreamRequest
type StopKeepAliveRequest StreamRequest

func NewGetKeepAliveRequest(streamSessionId string) *GetKeepAliveRequest {
	return &GetKeepAliveRequest{
		Command:         CmdGetKeepAlive,
		StreamSessionId: streamSessionId,
	}
}

func NewStopKeepAliveRequest() *StopKeepAliveRequest {
	return &StopKeepAliveRequest{
		Command: CmdStopKeepAlive,
	}
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
	StreamSessionId string `json:"streamSessionId,omitempty"`
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

func PingRequestWithStreamSessionId(id string) PingRequestOption {
	return func(request *PingRequest) {
		request.StreamSessionId = id
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

type Nil any
type ReturnData struct {
	SymbolRecords []*SymbolRecord
}

type StreamRequest struct {
	Command         string `json:"command"`
	StreamSessionId string `json:"streamSessionId,omitempty"`
}

type DataStream struct {
	Command string `json:"command"`
	Data    any    `json:"data"`
}

func (ds *DataStream) UnmarshalJSON(data []byte) (err error) {
	var helper = struct {
		Command string      `json:"command"`
		Data    interface{} `json:"data"`
	}{}

	if err = json.Unmarshal(data, &helper); err != nil {
		return
	}

	if helper.Data == nil {
		err = errors.New("empty data frame in stream")
		return
	}

	var factory, ok = streamingDataTypeFactory[helper.Command]
	if !ok {
		err = errors.Errorf("unsupported stream data type %s", helper.Command)
		return
	}
	var value = factory()

	if err = mapstructure.Decode(helper.Data, value); err != nil {
		return
	}
	ds.Command = helper.Command
	ds.Data = value
	return
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

type StreamingBalanceRecord struct {
	Balance     float64 `json:"balance"`
	Credit      float64 `json:"credit"`
	Equity      float64 `json:"equity"`
	Margin      float64 `json:"margin"`
	MarginFree  float64 `json:"marginFree"`
	MarginLevel float64 `json:"marginLevel"`
}

func (r *StreamingBalanceRecord) String() string {
	var data, _ = json.Marshal(r)
	return string(data)
}

type StreamingCandleRecord struct {
	Close     float64 `json:"close"`
	Ctm       int     `json:"ctm"`
	CtmString string  `json:"ctmString"`
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Open      float64 `json:"open"`
	QuoteId   int     `json:"quoteId"`
	Symbol    string  `json:"symbol"`
	Vol       float64 `json:"vol"`
}

func (r *StreamingCandleRecord) String() string {
	var data, _ = json.Marshal(r)
	return string(data)
}

type StreamingKeepAliveRecord struct {
	Timestamp int `json:"timestamp"`
}

func (r *StreamingKeepAliveRecord) String() string {
	var data, _ = json.Marshal(r)
	return string(data)
}

type StreamingNewsRecord struct {
	Body  string `json:"body"`
	Key   string `json:"key"`
	Time  int    `json:"time"`
	Title string `json:"title"`
}

type StreamingProfitRecord struct {
	Order         int     `json:"order"`
	TransactionId int     `json:"order2"`
	Position      int     `json:"position"`
	Profit        float64 `json:"profit"`
}

type StreamingTickRecord struct {
	Ask         float64 `json:"ask"`
	AskVolume   int     `json:"askVolume"`
	Bid         float64 `json:"bid"`
	BidVolume   int     `json:"bidVolume"`
	High        float64 `json:"high"`
	Low         float64 `json:"low"`
	QuoteId     int     `json:"quoteId"`
	SpreadRaw   float64 `json:"spreadRaw"`
	SpreadTable float64 `json:"spreadTable"`
	Symbol      string  `json:"symbol"`
	Timestamp   int     `json:"timestamp"`
}

type StreamingTradeRecord struct {
	ClosePrice    float64 `json:"close_price"`
	CloseTime     *int    `json:"close_time,omitempty"`
	Closed        bool    `json:"closed"`
	Cmd           int     `json:"cmd"`
	Comment       string  `json:"comment"`
	Commission    float64 `json:"commission"`
	CustomComment string  `json:"customComment"`
	Digits        int     `json:"digits"`
	Expiration    *int    `json:"expiration,omitempty"`
	MarginRate    float64 `json:"margin_rate"`
	Offset        int     `json:"offset"`
	OpenPrice     float64 `json:"open_price"`
	OpenTime      int     `json:"open_time"`
	Order         int     `json:"order"`
	TransactionId int     `json:"order2"`
	Position      int     `json:"position"`
	Profit        float64 `json:"profit"`
	StopLoss      float64 `json:"sl"`
	State         string  `json:"state"`
	Storage       float64 `json:"storage"`
	Symbol        string  `json:"symbol"`
	TakeProfit    float64 `json:"tp"`
	Type          int     `json:"type"`
	Volume        float64 `json:"volume"`
}

type StreamingTradeStatusRecord struct {
	CustomComment string  `json:"customComment,omitempty"`
	Message       *string `json:"message,omitempty"`
	Order         int     `json:"order"`
	Price         float64 `json:"price"`
	RequestStatus int     `json:"requestStatus"`
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
