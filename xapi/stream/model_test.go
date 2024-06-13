package stream

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/wI2L/jsondiff"
	"os"
	"testing"
)

func TestConstructor(t *testing.T) {
	var tests = []struct {
		Name           string
		RequestFactory func() any
		Want           any
	}{
		{
			Name: CmdGetBalance,
			RequestFactory: func() any {
				return NewGetBalanceRequest("streamSessionId")
			},
			Want: &GetBalanceRequest{
				Command:         CmdGetBalance,
				StreamSessionId: "streamSessionId",
			},
		},
		{
			Name: CmdGetCandles,
			RequestFactory: func() any {
				return NewGetCandlesRequest("symbol", "streamSessionId")
			},
			Want: &GetCandlesRequest{
				Request: Request{
					Command:         CmdGetCandles,
					StreamSessionId: "streamSessionId",
				},
				Symbol: "symbol",
			},
		},
		{
			Name: CmdGetKeepAlive,
			RequestFactory: func() any {
				return NewGetKeepAliveRequest("8469308861804289383")
			},
			Want: &GetKeepAliveRequest{
				Command:         CmdGetKeepAlive,
				StreamSessionId: "8469308861804289383",
			},
		},
		{
			Name: CmdStopKeepAlive,
			RequestFactory: func() any {
				return NewStopKeepAliveRequest()
			},
			Want: &StopKeepAliveRequest{
				Command: CmdStopKeepAlive,
			},
		},
		{
			Name: CmdGetNews,
			RequestFactory: func() any {
				return NewGetNewsRequest("1234567890")
			},
			Want: &GetNewsRequest{
				Command:         CmdGetNews,
				StreamSessionId: "1234567890",
			},
		},
		{
			Name: CmdPing,
			RequestFactory: func() any {
				return NewPingRequest("8469308861804289383")
			},
			Want: &PingRequest{
				Command:         CmdPing,
				StreamSessionId: "8469308861804289383",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var actual = test.RequestFactory()
			assert.Equal(t, test.Want, actual)
		})
	}
}

func TestRequestJsonCoding(t *testing.T) {
	type testData struct {
		Want   string
		Actual any
	}
	var (
		tests = []struct {
			Name string
			Data []testData
		}{
			{
				Name: CmdGetBalance,
				Data: []testData{
					{
						Want:   "testdata/getBalance.request.json",
						Actual: NewGetBalanceRequest("8469308861804289383"),
					},
					{
						Want:   "testdata/stopBalance.request.json",
						Actual: NewStopBalanceRequest(),
					},
					{
						Want: "testdata/balance.stream.json",
						Actual: &DataStream{
							Command: DataStreamBalance,
							Data: &StreamingBalanceRecord{
								Balance:     995800269.43,
								Credit:      1000.0,
								Equity:      995985397.56,
								Margin:      572634.43,
								MarginFree:  995227635.00,
								MarginLevel: 173930.41,
							},
						},
					},
				},
			},
			{
				Name: CmdGetCandles,
				Data: []testData{
					{
						Want:   "testdata/getCandles.request.json",
						Actual: NewGetCandlesRequest("EURUSD", "8469308861804289383"),
					},
					{
						Want: "testdata/candle.stream.json",
						Actual: &DataStream{
							Command: DataStreamCandle,
							Data: &StreamingCandleRecord{
								Close:     4.1849,
								Ctm:       1378369375000,
								CtmString: "Sep 05, 2013 10:22:55 AM",
								High:      4.1854,
								Low:       4.1848,
								Open:      4.1848,
								QuoteId:   2,
								Symbol:    "EURUSD",
								Vol:       0,
							},
						},
					},
					{
						Want:   "testdata/stopCandles.request.json",
						Actual: NewStopCandlesRequest("EURUSD"),
					},
				},
			},
			{
				Name: CmdGetKeepAlive,
				Data: []testData{
					{
						Want: "testdata/getKeepAlive.request.json",
						Actual: &GetKeepAliveRequest{
							Command:         CmdGetKeepAlive,
							StreamSessionId: "8469308861804289383",
						},
					},
					{
						Want: "testdata/stopKeepAlive.request.json",
						Actual: &StopKeepAliveRequest{
							Command: CmdStopKeepAlive,
						},
					},
					{
						Want: "testdata/keepAlive.stream.json",
						Actual: &DataStream{
							Command: DataStreamKeepAlive,
							Data: &StreamingKeepAliveRecord{
								Timestamp: 1362944112000,
							},
						},
					},
				},
			},
			{
				Name: CmdGetNews,
				Data: []testData{
					{
						Want:   "testdata/getNews.request.json",
						Actual: NewGetNewsRequest("1234567890"),
					},
					{
						Want: "testdata/news.stream.json",
						Actual: &DataStream{
							Command: DataStreamNews,
							Data: &StreamingNewsRecord{
								Body:  "<html>...</html>",
								Key:   "1f6da766abd29927aa854823f0105c23",
								Time:  1262944112000,
								Title: "Breaking trend",
							},
						},
					},
				},
			},
			{
				Name: CmdPing,
				Data: []testData{
					{
						Want:   "testdata/ping.request.json",
						Actual: NewPingRequest("1234567890"),
					},
				},
			},
		}
	)

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var (
				actual, want []byte
				err          error
				patch        jsondiff.Patch
			)
			for idx, data := range test.Data {
				want, err = os.ReadFile(data.Want)
				assert.NoError(t, err)
				actual, err = json.Marshal(data.Actual)
				patch, err = jsondiff.CompareJSON(want, actual, jsondiff.Equivalent())
				assert.Truef(t, len(patch) == 0, "data[%d] differs %s", idx, patch.String())
			}
		})
	}
}

func TestDataStreamUnmarshalJSON(t *testing.T) {
	var (
		data, _ = os.ReadFile("testdata/data.stream.json")
		actual  = make([]DataStream, 0)
		err     = json.Unmarshal(data, &actual)
		want    = []DataStream{
			{
				Command: DataStreamBalance,
				Data: &StreamingBalanceRecord{
					Balance:     995800269.43,
					Credit:      1000.00,
					Equity:      995985397.56,
					Margin:      572634.43,
					MarginFree:  995227635.00,
					MarginLevel: 173930.41,
				},
			},
			{
				Command: DataStreamCandle,
				Data: &StreamingCandleRecord{
					Close:     4.1849,
					Ctm:       1378369375000,
					CtmString: "Sep 05, 2013 10:22:55 AM",
					High:      4.1854,
					Low:       4.1848,
					Open:      4.1848,
					QuoteId:   2,
					Symbol:    "EURUSD",
					Vol:       0,
				},
			},
			{
				Command: DataStreamKeepAlive,
				Data: &StreamingKeepAliveRecord{
					Timestamp: 1362944112000,
				},
			},
			{
				Command: DataStreamNews,
				Data: &StreamingNewsRecord{
					Body:  "<html>...</html>",
					Key:   "1f6da766abd29927aa854823f0105c23",
					Time:  1262944112000,
					Title: "Breaking trend",
				},
			},
		}
	)
	assert.NoError(t, err)
	assert.Equal(t, want, actual)
}