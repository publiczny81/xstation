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

func NewGetNewsRequest(streamSessionId string) *GetNewsRequest {
	return &GetNewsRequest{
		Command:         CmdGetNews,
		StreamSessionId: streamSessionId,
	}
}

type PingRequest Request

func NewPingRequest(streamSessionId string) *PingRequest {
	return &PingRequest{
		Command:         CmdPing,
		StreamSessionId: streamSessionId,
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
