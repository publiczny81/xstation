package model

import (
	"encoding/json"
	"github.com/publiczny81/xstation/xapi/utils"
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
			Name: CmdGetAllSymbols,
			RequestFactory: func() any {
				return NewGetAllSymbolsRequest(GetAllSymbolsRequestWithCustomTag("tag"), GetAllSymbolsRequestWithPrettyPrint(true))
			},
			Want: &GetAllSymbolsRequest{
				Command:     CmdGetAllSymbols,
				CustomTag:   "tag",
				PrettyPrint: true,
			},
		},
		{
			Name: CmdGetCalendar,
			RequestFactory: func() any {
				return NewGetCalendarRequest(GetCalendarRequestWithCustomTag("tag"), GetCalendarRequestWithPrettyPrint(true))
			},
			Want: &GetCalendarRequest{
				Command:     CmdGetCalendar,
				CustomTag:   "tag",
				PrettyPrint: true,
			},
		},
		{
			Name: CmdGetChartLastRequest,
			RequestFactory: func() any {
				return NewGetChartLastRequest("symbol", 1, 1234567890,
					GetChartLastRequestWithCustomTAg("tag"),
					GetChartLastRequestWithPrettyPrint(true))
			},
			Want: &GetChartLastRequest{
				Command:     CmdGetChartLastRequest,
				CustomTag:   "tag",
				PrettyPrint: true,
				Arguments: ChartLastInfoRecord{
					Start:  1234567890,
					Period: 1,
					Symbol: "symbol",
				},
			},
		},
		{
			Name: CmdGetChartRangeRequest,
			RequestFactory: func() any {
				return NewGetChartRangeRequest("symbol", 1, 1234567890, 1234567900, 2,
					GetChartRangeRequestWithCustomTag("tag"),
					GetChartRangeRequestWithPrettyPrint(true))
			},
			Want: &GetChartRangeRequest{
				Command:     CmdGetChartRangeRequest,
				CustomTag:   "tag",
				PrettyPrint: true,
				Arguments: ChartRangeInfoRecord{
					End:    1234567900,
					Start:  1234567890,
					Period: 1,
					Symbol: "symbol",
					Ticks:  2,
				},
			},
		},
		{
			Name: CmdGetCommissionDef,
			RequestFactory: func() any {
				return NewGetCommissionDefRequest("symbol", 2,
					GetCommissionDefWithPrettyPrint(true),
					GetCommissionDefWithCustomTag("tag"))
			},
			Want: &GetCommissionDefRequest{
				Command:     CmdGetCommissionDef,
				CustomTag:   "tag",
				PrettyPrint: true,
				Arguments: GetCommissionDefArgs{
					Symbol: "symbol",
					Volume: 2,
				},
			},
		},
		{
			Name: CmdGetCurrentUserData,
			RequestFactory: func() any {
				return NewGetCurrentUserDataRequest(
					GetCurrentUserDataRequestWithCustomTag("tag"),
					GetCurrentUserDataRequestWithPrettyPrint(true))
			},
			Want: &GetCurrentUserDataRequest{
				Command:     CmdGetCurrentUserData,
				PrettyPrint: true,
				CustomTag:   "tag",
			},
		},
		{
			Name: CmdGetIbsHistory,
			RequestFactory: func() any {
				return NewGetIbsHistoryRequest(1, 2,
					GetIbsHistoryRequestWithCustomTag("tag"),
					GetIbsHistoryRequestWithPrettyPrint(true))
			},
			Want: &GetIbsHistoryRequest{
				Command: CmdGetIbsHistory,
				Arguments: IbsHistoryArgs{
					Start: 1,
					End:   2,
				},
				CustomTag:   "tag",
				PrettyPrint: true,
			},
		},
		{
			Name: CmdGetMarginLevel,
			RequestFactory: func() any {
				return NewGetMarginLevelRequest(GetMarginLevelRequestWithCustomTag("tag"), GetMarginLevelRequestWithPrettyPrint(true))
			},
			Want: &GetMarginLevelRequest{
				Command:     CmdGetMarginLevel,
				CustomTag:   "tag",
				PrettyPrint: true,
			},
		},
		{
			Name: CmdGetMarginTrade,
			RequestFactory: func() any {
				return NewGetMarginTradeRequest("EURPLN", 1.0, GetMarginTradeRequestWithCustomTag("tag"),
					GetMarginTradeRequestWithPrettyPrint(true))
			},
			Want: &GetMarginTradeRequest{
				Command: CmdGetMarginTrade,
				Arguments: GetMarginTradeArgs{
					Symbol: "EURPLN",
					Volume: 1.0,
				},
				PrettyPrint: true,
				CustomTag:   "tag",
			},
		},
		{
			Name: CmdLogin,
			RequestFactory: func() any {
				return NewLoginRequest("userId",
					"password",
					LoginRequestWithCustomTag("tag"),
					LoginRequestWithPrettyPrint(true),
					LoginRequestWithAppId("appId"),
					LoginRequestWithAppName("appName"))
			},
			Want: &LoginRequest{
				Command:     CmdLogin,
				CustomTag:   "tag",
				PrettyPrint: true,
				Arguments: LoginArguments{
					UserId:   "userId",
					Password: "password",
					AppId:    "appId",
					AppName:  "appName",
				},
			},
		},
		{
			Name: CmdLogout,
			RequestFactory: func() any {
				return NewLogoutRequest(
					LogoutRequestWithCustomTag("tag"),
					LogoutRequestWithPrettyPrint(true))
			},
			Want: &LogoutRequest{
				Command:     CmdLogout,
				CustomTag:   "tag",
				PrettyPrint: true,
			},
		},
		{
			Name: CmdPing,
			RequestFactory: func() any {
				return NewPingRequest(PingRequestWithCustomTag("tag"), PingRequestWithPrettyPrint(true))
			},
			Want: &PingRequest{
				Request: Request[Nil]{
					Command:     CmdPing,
					CustomTag:   "tag",
					PrettyPrint: true,
				},
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

func TestRequestResponseJsonCoding(t *testing.T) {
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
				Name: CmdGetAllSymbols,
				Data: []testData{
					{
						Want: "testdata/getAllSymbols.request.json",
						Actual: &GetAllSymbolsRequest{
							Command:     CmdGetAllSymbols,
							CustomTag:   "tag",
							PrettyPrint: true,
						},
					},
					{
						Want: "testdata/getAllSymbols.response.json",
						Actual: &GetAllSymbolsResponse{
							Status:    true,
							CustomTag: "tag",
							ReturnData: []SymbolRecord{
								{
									Ask:               4000.0,
									Bid:               4000.0,
									CategoryName:      "Forex",
									ContractSize:      100000,
									Currency:          "USD",
									CurrencyPair:      true,
									CurrencyProfit:    "SEK",
									Description:       "USD/PLN",
									GroupName:         "Minor",
									High:              4000,
									InitialMargin:     0,
									InstantMaxVolume:  0,
									Leverage:          1.5,
									LongOnly:          false,
									LotMax:            10,
									LotMin:            0.1,
									LotStep:           0.1,
									Low:               3500,
									MarginMode:        101,
									Percentage:        100,
									Precision:         2,
									ProfitMode:        5,
									QuoteId:           1,
									ShortSelling:      true,
									SpreadRaw:         0.000003,
									SpreadTable:       0.00042,
									StepRuleId:        1,
									SwapEnable:        true,
									SwapLong:          -2.55929,
									SwapRollover3Days: 1,
									SwapShort:         0.131,
									Symbol:            "USDPLN",
									TickSize:          1,
									TickValue:         1,
									Time:              1272446136891,
									TimeString:        "Thu May 23 12:23:44 EDT 2013",
									TrailingEnabled:   true,
									Type:              21,
								},
							},
						},
					},
				},
			},
			{
				Name: CmdGetCalendar,
				Data: []testData{
					{
						Want:   "testdata/getCalendar.request.json",
						Actual: NewGetCalendarRequest(GetCalendarRequestWithCustomTag("tag")),
					},
					{
						Want: "testdata/getCalendar.response.json",
						Actual: &GetCalendarResponse{
							Status: true,
							ReturnData: []CalendarRecord{
								{
									Country:  "CA",
									Impact:   ImpactHigh,
									Period:   "(FEB)",
									Previous: "58.3",
									Time:     1374846900000,
									Title:    "Ivey Purchasing Managers Index",
								},
								{
									Country:  "PL",
									Impact:   ImpactMedium,
									Period:   "(FEB)",
									Previous: "51.1",
									Time:     1374846900000,
									Title:    "Tax rate",
								},
							},
						},
					},
				},
			},
			{
				Name: CmdGetChartLastRequest,
				Data: []testData{
					{
						Want: "testdata/getChartLast.request.json",
						Actual: NewGetChartLastRequest("PKN.PL", PeriodM5, 1262944112000,
							GetChartLastRequestWithPrettyPrint(true)),
					},
					{
						Want: "testdata/getChartLast.response.json",
						Actual: GetChartLastResponse{
							Status: true,
							ReturnData: RateInfoData{
								Digits: 4,
								RateInfos: []RateInfoRecord{
									{
										Close:     1,
										Ctm:       1389362640000,
										CtmString: "Jan 10, 2014 3:04:00 PM",
										High:      6,
										Low:       0,
										Open:      41848.0,
										Vol:       0,
									},
								},
							},
						},
					},
				},
			},
			{
				Name: CmdGetChartRangeRequest,
				Data: []testData{
					{
						Want: "testdata/getChartRange.request.json",
						Actual: NewGetChartRangeRequest("PKN.PL", PeriodM5, 1262944112000, 1262944412000, 0,
							GetChartRangeRequestWithPrettyPrint(true)),
					},
					{
						Want: "testdata/getChartRange.response.json",
						Actual: GetChartRangeResponse{
							Status: true,
							ReturnData: RateInfoData{
								Digits: 4,
								RateInfos: []RateInfoRecord{
									{
										Close:     1,
										Ctm:       1389362640000,
										CtmString: "Jan 10, 2014 3:04:00 PM",
										High:      6,
										Low:       0,
										Open:      41848.0,
										Vol:       0,
									},
								},
							},
						},
					},
				},
			},
			{
				Name: CmdGetCommissionDef,
				Data: []testData{
					{
						Want:   "testdata/getCommissionDef.request.json",
						Actual: NewGetCommissionDefRequest("T.US", 1.0),
					},
					{
						Want: "testdata/getCommissionDef.response.json",
						Actual: GetCommissionDefResponse{
							Status: true,
							ReturnData: GetCommissionDefData{
								Commission:     0.51,
								RateOfExchange: 0.1609,
							},
						},
					},
				},
			},
			{
				Name: CmdGetCurrentUserData,
				Data: []testData{
					{
						Want: "testdata/getCurrentUserData.request.json",
						Actual: NewGetCurrentUserDataRequest(GetCurrentUserDataRequestWithCustomTag("tag"),
							GetCurrentUserDataRequestWithPrettyPrint(true)),
					},
					{
						Want: "testdata/getCurrentUserData.response.json",
						Actual: GetCurrentUserDataResponse{
							Status: true,
							ReturnData: CurrentUserData{
								CompanyUnit:        8,
								Currency:           "PLN",
								Group:              "demoPLeurSTANDARD200",
								IbAccount:          false,
								Leverage:           1,
								LeverageMultiplier: 0.25,
								SpreadType:         utils.Pointer("FLOAT"),
								TrailingStop:       false,
							},
						},
					},
				},
			},
			{
				Name: CmdGetIbsHistory,
				Data: []testData{
					{
						Want:   "testdata/getIbsHistory.request.json",
						Actual: NewGetIbsHistoryRequest(1394449010991, 1395053810991),
					},
					{
						Want: "testdata/getIbsHistory.response.json",
						Actual: &GetIbsHistoryResponse{
							Status: true,
							ReturnData: []IbRecord{
								{
									ClosePrice: utils.Pointer(1.39302),
									Login:      utils.Pointer("12345"),
									Nominal:    utils.Pointer(6.00),
									OpenPrice:  utils.Pointer(1.39376),
									Side:       utils.Pointer(0),
									Surname:    utils.Pointer("IB_Client_1"),
									Symbol:     utils.Pointer("EURUSD"),
									Timestamp:  utils.Pointer(1395755870000),
									Volume:     utils.Pointer(1.0),
								},
							},
						},
					},
				},
			},
			{
				Name: CmdGetMarginLevel,
				Data: []testData{
					{
						Want:   "testdata/getMarginLevel.request.json",
						Actual: NewGetMarginLevelRequest(),
					},
					{
						Want: "testdata/getMarginLevel.response.json",
						Actual: &GetMarginLevelResponse{
							Status: true,
							ReturnData: GetMarginLevelData{
								Balance:     995800269.43,
								Credit:      1000.00,
								Currency:    "PLN",
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
				Name: CmdGetMarginTrade,
				Data: []testData{
					{
						Want: "testdata/getMarginTrade.request.json",
						Actual: &GetMarginTradeRequest{
							Command: CmdGetMarginTrade,
							Arguments: GetMarginTradeArgs{
								Symbol: "EURPLN",
								Volume: 1.0,
							},
						},
					},
					{
						Want: "testdata/getMarginTrade.response.json",
						Actual: GetMarginTradeResponse{
							Status: true,
							ReturnData: GetMarginTradeData{
								Margin: 4399.350,
							},
						},
					},
				},
			},
			{
				Name: CmdLogin,
				Data: []testData{
					{
						Want: "testdata/login.request.json",
						Actual: &LoginRequest{
							Command:     CmdLogin,
							CustomTag:   "tag",
							PrettyPrint: true,
							Arguments: LoginArguments{
								UserId:   "userId",
								Password: "password",
								AppId:    "appId",
								AppName:  "appName",
							},
						},
					},
					{
						Want: "testdata/login.response.json",
						Actual: LoginResponse{
							Status:          true,
							CustomTag:       "tag",
							StreamSessionId: "streamSessionId"},
					},
					{
						Want: "testdata/login.error.json",
						Actual: LoginResponse{
							Status:     true,
							CustomTag:  "tag",
							ErrorCode:  "EX004",
							ErrorDescr: ErrDescIncorrectCredentials,
						},
					},
				},
			},
			{
				Name: CmdLogout,
				Data: []testData{
					{
						Want: "testdata/logout.request.json",
						Actual: &LogoutRequest{
							Command:     CmdLogout,
							CustomTag:   "tag",
							PrettyPrint: true,
						},
					},
					{
						Want: "testdata/logout.response.json",
						Actual: LogoutResponse{
							Status:    true,
							CustomTag: "tag",
						},
					},
				},
			},
			{
				Name: CmdPing,
				Data: []testData{
					{
						Want: "testdata/ping.request.json",
						Actual: NewPingRequest(
							PingRequestWithCustomTag("tag"),
							PingRequestWithStreamSessionId("1234567890")),
					},
					{
						Want: "testdata/ping.response.json",
						Actual: &PingResponse{
							Status: true,
						},
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
