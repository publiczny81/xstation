package stream

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

type Request struct {
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

type GetBalanceRequest Request
type StopBalanceRequest Request

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

type GetCandlesRequest struct {
	Request
	Symbol string `json:"symbol"`
}

type StopCandlesRequest struct {
	Request
	Symbol string `json:"symbol"`
}

func NewGetCandlesRequest(symbol, streamSessionId string) *GetCandlesRequest {
	return &GetCandlesRequest{
		Request: Request{
			Command:         CmdGetCandles,
			StreamSessionId: streamSessionId,
		},
		Symbol: symbol,
	}
}

func NewStopCandlesRequest(symbol string) *StopCandlesRequest {
	return &StopCandlesRequest{
		Request: Request{
			Command: CmdStopCandles,
		},
		Symbol: symbol,
	}
}

type GetKeepAliveRequest Request
type StopKeepAliveRequest Request

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

type GetNewsRequest Request
type StopNewsRequest Request

func NewGetNewsRequest(streamSessionId string) *GetNewsRequest {
	return &GetNewsRequest{
		Command:         CmdGetNews,
		StreamSessionId: streamSessionId,
	}
}

func NewStopNewsRequest() *StopNewsRequest {
	return &StopNewsRequest{
		Command: CmdStopNews,
	}
}

type PingRequest Request

func NewPingRequest(streamSessionId string) *PingRequest {
	return &PingRequest{
		Command:         CmdPing,
		StreamSessionId: streamSessionId,
	}
}

type GetProfits Request
type StopProfits Request

func NewGetProfits(streamSessionId string) *GetProfits {
	return &GetProfits{
		Command:         CmdGetProfits,
		StreamSessionId: streamSessionId,
	}
}

func NewStopProfits() *StopProfits {
	return &StopProfits{
		Command: CmdStopProfits,
	}
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
	Order    int     `json:"order"`
	Order2   int     `json:"order2"`
	Position int     `json:"position"`
	Profit   float64 `json:"profit"`
}

type StreamingTickRecord struct {
	Ask         float64 `json:"ask"`
	AskVolume   int     `json:"askVolume"`
	Bid         float64 `json:"bid"`
	BidVolume   int     `json:"bidVolume"`
	High        float64 `json:"high"`
	Level       int     `json:"level"`
	Low         float64 `json:"low"`
	QuoteId     int     `json:"quoteId"`
	SpreadRaw   float64 `json:"spreadRaw"`
	SpreadTable float64 `json:"spreadTable"`
	Symbol      string  `json:"symbol"`
	Timestamp   int     `json:"timestamp"`
}

type StreamingTradeRecord struct {
	ClosePrice    float64 `json:"close_price" mapstructure:"close_price"`
	CloseTime     *int    `json:"close_time" mapstructure:"close_time"`
	Closed        bool    `json:"closed" mapstructure:"closed"`
	Cmd           int     `json:"cmd" mapstructure:"cmd"`
	Comment       string  `json:"comment" mapstructure:"comment"`
	Commission    float64 `json:"commission" mapstructure:"commission"`
	CustomComment string  `json:"customComment" mapstructure:"customComment"`
	Digits        int     `json:"digits" mapstructure:"digits"`
	Expiration    *int    `json:"expiration" mapstructure:"expiration,omitempty"`
	MarginRate    float64 `json:"margin_rate" mapstructure:"margin_rate"`
	Offset        int     `json:"offset" mapstructure:"offset"`
	OpenPrice     float64 `json:"open_price" mapstructure:"open_price"`
	OpenTime      int     `json:"open_time" mapstructure:"open_time"`
	Order         int     `json:"order"mapstructure:"order"`
	TransactionId int     `json:"order2" mapstructure:"order2"`
	Position      int     `json:"position" mapstructure:"position"`
	Profit        float64 `json:"profit" mapstructure:"profit"`
	StopLoss      float64 `json:"sl" mapstructure:"sl"`
	State         string  `json:"state" mapstructure:"state"`
	Storage       float64 `json:"storage" mapstructure:"storage"`
	Symbol        string  `json:"symbol" mapstructure:"symbol"`
	TakeProfit    float64 `json:"tp" mapstructure:"tp"`
	Type          int     `json:"type" mapstructure:"type"`
	Volume        float64 `json:"volume" mapstructure:"volume"`
}

type StreamingTradeStatusRecord struct {
	CustomComment string  `json:"customComment,omitempty"`
	Message       *string `json:"message"`
	Order         int     `json:"order"`
	Price         float64 `json:"price"`
	RequestStatus int     `json:"requestStatus"`
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

type GetTickPricesRequest struct {
	Request
	Symbol         string `json:"symbol"`
	MinArrivalTime int    `json:"minArrivalTime"`
	MaxLevel       int    `json:"maxLevel"`
}

func NewGetTickPricesRequest(streamSessionId, symbol string, minArrivalTime, maxLevel int) *GetTickPricesRequest {
	return &GetTickPricesRequest{
		Request: Request{
			Command:         CmdGetTickPrices,
			StreamSessionId: streamSessionId,
		},
		Symbol:         symbol,
		MinArrivalTime: minArrivalTime,
		MaxLevel:       maxLevel,
	}
}

type StopTickPricesRequest struct {
	Request
	Symbol string `json:"symbol"`
}

func NewStopTickPricesRequest(symbol string) *StopTickPricesRequest {
	return &StopTickPricesRequest{
		Request: Request{Command: CmdStopTickPrices},
		Symbol:  symbol,
	}
}

type GetTradesRequest Request

func NewGetTradesRequest(streamSessionId string) *GetTradesRequest {
	return &GetTradesRequest{
		Command:         CmdGetTrades,
		StreamSessionId: streamSessionId,
	}
}

type StopTradesRequest Request

func NewStopTradesRequest() *StopTradesRequest {
	return &StopTradesRequest{
		Command: CmdStopTrades,
	}
}

type GetTradeStatusRequest Request

func NewGetTradeStatusRequest(streamSessionId string) *GetTradeStatusRequest {
	return &GetTradeStatusRequest{
		Command:         CmdGetTradeStatus,
		StreamSessionId: streamSessionId,
	}
}

type StopTradeStatusRequest Request

func NewStopTradeStatusRequest() *StopTradeStatusRequest {
	return &StopTradeStatusRequest{
		Command: CmdStopTradeStatus,
	}
}
