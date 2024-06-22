package model

import (
	"encoding/json"
	"github.com/publiczny81/xstation/xapi/common"
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
			Name: CmdGetNews,
			RequestFactory: func() any {
				return NewGetNewsRequest(100, 120)
			},
			Want: &GetNewsRequest{
				Command:   CmdGetNews,
				Arguments: NewsRequestArg{Start: 100, End: 120},
			},
		},
		{
			Name: CmdGetProfitCalculation,
			RequestFactory: func() any {
				return NewGetProfitCalculationRequest("USDPLN", 1, 1.1, 1.2, 3)
			},
			Want: &GetProfitCalculationRequest{
				Command: CmdGetProfitCalculation,
				Arguments: GetProfitCalculationArgs{
					Symbol:     "USDPLN",
					Cmd:        1,
					OpenPrice:  1.1,
					ClosePrice: 1.2,
					Volume:     3,
				},
			},
		},
		{
			Name: CmdGetServerTime,
			RequestFactory: func() any {
				return NewGetServerTimeRequest()
			},
			Want: &GetServerTimeRequest{
				Command: CmdGetServerTime,
			},
		},
		{
			Name: CmdGetStepRules,
			RequestFactory: func() any {
				return NewGetStepRulesRequest(GetStepRulesRequestWithCustomTag("tag"), GetStepRulesRequestWithPrettyPrint(true))
			},
			Want: &GetStepRulesRequest{
				Command:     CmdGetStepRules,
				CustomTag:   "tag",
				PrettyPrint: true,
			},
		},
		{
			Name: CmdGetSymbol,
			RequestFactory: func() any {
				return NewGetSymbolRequest("USDPLN", GetSymbolRequestWithCustomTag("tag"), GetSymbolRequestWithPrettyPrint(true))
			},
			Want: &GetSymbolRequest{
				Command: CmdGetSymbol,
				Arguments: GetSymbolRequestArgs{
					Symbol: "USDPLN",
				},
				PrettyPrint: true,
				CustomTag:   "tag",
			},
		},
		{
			Name: CmdGetTickPrices,
			RequestFactory: func() any {
				return NewGetTickPricesRequest(LevelAll,
					1234567890,
					GetTickPricesRequestWithSymbol("EURPLN"),
					GetTickPricesRequestWithCustomTag("tag"),
					GetTickPricesRequestWithPrettyPrint(true))
			},
			Want: &GetTickPricesRequest{
				Command:     CmdGetTickPrices,
				CustomTag:   "tag",
				PrettyPrint: true,
				Arguments: GetTickPricesRequestArgs{
					Level:     LevelAll,
					Timestamp: 1234567890,
					Symbols:   []string{"EURPLN"},
				},
			},
		},
		{
			Name: CmdGetTradeRecords,
			RequestFactory: func() any {
				return NewGetTradeRecordsRequest(GetTradeRecordsRequestWithOrder(1),
					GetTradeRecordsRequestWithCustomTag("tag"),
					GetTradeRecordsRequestWithPrettyPrint(true))
			},
			Want: &GetTradeRecordsRequest{
				Command:     CmdGetTradeRecords,
				CustomTag:   "tag",
				PrettyPrint: true,
				Arguments: GetTradeRecordsRequestArgs{
					Orders: []int{1},
				},
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
		{
			Name: CmdGetTrades,
			RequestFactory: func() any {
				return NewGetTradesRequest(GetTradesRequestWithOpenedOnly(true),
					GetTradesRequestWithCustomTag("tag"),
					GetTradesRequestWithPrettyPrint(true))
			},
			Want: &GetTradesRequest{
				Command:     CmdGetTrades,
				CustomTag:   "tag",
				PrettyPrint: true,
				Arguments: GetTradesArgs{
					OpenedOnly: true,
				},
			},
		},
		{
			Name: CmdGetTradesHistory,
			RequestFactory: func() any {
				return NewGetTradesHistoryRequest(1, 2,
					GetTradesHistoryRequestWithCustomTag("tag"),
					GetTradesHistoryRequestWithPrettyPrint(true))
			},
			Want: &GetTradesHistoryRequest{
				Command:     CmdGetTradesHistory,
				CustomTag:   "tag",
				PrettyPrint: true,
				Arguments: GetTradesHistoryArgs{
					Start: 1,
					End:   2,
				},
			},
		},
		{
			Name: CmdGetTradingHours,
			RequestFactory: func() any {
				return NewGetTradingHoursRequest(GetTradingHoursRequestWithSymbol("USDPLN"),
					GetTradingHoursRequestWithCustomTag("tag"),
					GetTradingHoursRequestWithPrettyPrint(true))
			},
			Want: &GetTradingHoursRequest{
				Command:     CmdGetTradingHours,
				CustomTag:   "tag",
				PrettyPrint: true,
				Arguments: GetTradingHoursArgs{
					Symbols: []string{"USDPLN"},
				},
			},
		},
		{
			Name: CmdGetVersion,
			RequestFactory: func() any {
				return NewGetVersionRequest(GetVersionRequestWithCustomTag("tag"),
					GetVersionRequestWithPrettyPrint(true))
			},
			Want: &GetVersionRequest{
				Command:     CmdGetVersion,
				CustomTag:   "tag",
				PrettyPrint: true,
			},
		},
		{
			Name: CmdTradeTransaction,
			RequestFactory: func() (r any) {
				r, _ = NewTradeTransactionRequest("USDPLN",
					common.OperationCodeSell,
					common.OperationTypeOpen,
					4.01,
					0.1,
					TradeTransactionRequestWithCustomComment("comment"),
					TradeTransactionRequestWithPrettyPrint(true),
					TradeTransactionRequestWithCustomTag("tag"),
					TradeTransactionRequestWithStopLoss(4.02),
					TradeTransactionRequestWithTakeProfit(3.98),
					TradeTransactionRequestWithOffset(3),
					TradeTransactionRequestWithExpiration(1234567890))
				return
			},
			Want: &TradeTransactionRequest{
				Command:     CmdTradeTransaction,
				CustomTag:   "tag",
				PrettyPrint: true,
				Arguments: TradeTransactionArgs{
					TradeTransInfo{
						Cmd:           common.OperationCodeSell,
						CustomComment: "comment",
						Type:          common.OperationTypeOpen,
						Price:         4.01,
						Volume:        0.1,
						TakeProfit:    3.98,
						StopLoss:      4.02,
						Symbol:        "USDPLN",
						Offset:        3,
						Expiration:    1234567890,
					},
				},
			},
		},
		{
			Name: CmdTradeTransactionStatus,
			RequestFactory: func() any {
				return NewTradeTransactionStatusRequest(43,
					TradeTransactionStatusRequestWithCustomTag("tag"),
					TradeTransactionStatusRequestWithPrettyPrint(true))
			},
			Want: &TradeTransactionStatusRequest{
				Command:     CmdTradeTransactionStatus,
				CustomTag:   "tag",
				PrettyPrint: true,
				Arguments: TradeTransactionStatusArgs{
					Order: 43,
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
				Name: CmdGetNews,
				Data: []testData{
					{
						Want:   "testdata/getNews.request.json",
						Actual: NewGetNewsRequest(1275993488000, 0),
					},
					{
						Want: "testdata/getNews.response.json",
						Actual: &GetNewsResponse{
							Status: true,
							ReturnData: NewsTopicRecords{
								{
									Body:       "<html>...</html>",
									BodyLength: 110,
									Key:        "1f6da766abd29927aa854823f0105c23",
									Time:       1262944112000,
									TimeString: "May 17, 2013 4:30:00 PM",
									Title:      "Breaking trend",
								},
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
						Actual: NewPingRequest(PingRequestWithCustomTag("tag")),
					},
					{
						Want: "testdata/ping.response.json",
						Actual: &PingResponse{
							Status: true,
						},
					},
				},
			},
			{
				Name: CmdGetProfitCalculation,
				Data: []testData{
					{
						Want:   "testdata/getProfitCalculation.request.json",
						Actual: NewGetProfitCalculationRequest("EURPLN", 0, 1.2233, 1.3000, 1),
					},
					{
						Want: "testdata/getProfitCalculation.response.json",
						Actual: &GetProfitCalculationResponse{
							Status: true,
							ReturnData: ProfitRecord{
								Profit: 714.303,
							},
						},
					},
				},
			},
			{
				Name: CmdGetServerTime,
				Data: []testData{
					{
						Want:   "testdata/getServerTime.request.json",
						Actual: NewGetServerTimeRequest(GetServerTimeRequestWithCustomTag("tag"), GetServerTimeRequestWithPrettyPrint(true)),
					},
					{
						Want: "testdata/getServerTime.response.json",
						Actual: &GetServerTimeResponse{
							Status: true,
							ReturnData: ServerTimeRecord{
								Time:       1392211379731,
								TimeString: "Feb 12, 2014 2:22:59 PM",
							},
						},
					},
				},
			},
			{
				Name: CmdGetStepRules,
				Data: []testData{
					{
						Want:   "testdata/getStepRules.request.json",
						Actual: NewGetStepRulesRequest(GetStepRulesRequestWithCustomTag("tag"), GetStepRulesRequestWithPrettyPrint(true)),
					},
					{
						Want: "testdata/getStepRules.response.json",
						Actual: &GetStepRulesResponse{
							Status: true,
							ReturnData: []StepRuleRecord{
								{
									Id:   1,
									Name: "Forex",
									Steps: []StepRecord{
										{
											FromValue: 0.1,
											Step:      0.0025,
										},
									},
								},
							},
						},
					},
				},
			},
			{
				Name: CmdGetSymbol,
				Data: []testData{
					{
						Want:   "testdata/getSymbol.request.json",
						Actual: NewGetSymbolRequest("EURPLN"),
					},
					{
						Want: "testdata/getSymbol.response.json",
						Actual: &GetSymbolResponse{
							Status: true,
							ReturnData: SymbolRecord{
								Ask:                4000.0,
								Bid:                4000.0,
								CategoryName:       "Forex",
								ContractSize:       100000,
								Currency:           "USD",
								CurrencyPair:       true,
								CurrencyProfit:     "SEK",
								Description:        "USD/PLN",
								Expiration:         nil,
								GroupName:          "Minor",
								High:               4000.0,
								InitialMargin:      0,
								InstantMaxVolume:   0,
								Leverage:           1.5,
								LongOnly:           false,
								LotMax:             10.0,
								LotMin:             0.1,
								LotStep:            0.1,
								Low:                3500.0,
								MarginHedged:       0,
								MarginHedgedStrong: false,
								MarginMaintenance:  nil,
								MarginMode:         MarginModeForex,
								Percentage:         100,
								Precision:          2,
								ProfitMode:         ProfitModeForex,
								QuoteId:            common.QuoteIdFixed,
								ShortSelling:       true,
								SpreadRaw:          0.000003,
								SpreadTable:        0.00042,
								Starting:           nil,
								StepRuleId:         1,
								StopsLevel:         0,
								SwapRollover3Days:  0,
								SwapEnable:         true,
								SwapLong:           -2.55929,
								SwapShort:          0.131,
								SwapType:           0,
								Symbol:             "USDPLN",
								TickSize:           1.0,
								TickValue:          1.0,
								Time:               1272446136891,
								TimeString:         "Thu May 23 12:23:44 EDT 2013",
								TrailingEnabled:    true,
								Type:               21,
							},
						},
					},
				},
			},
			{
				Name: CmdGetTickPrices,
				Data: []testData{
					{
						Want: "testdata/getTickPrices.request.json",
						Actual: &GetTickPricesRequest{
							Command: CmdGetTickPrices,
							Arguments: GetTickPricesRequestArgs{
								Level:     0,
								Symbols:   []string{"EURPLN", "AGO.PL"},
								Timestamp: 1262944112000,
							},
						},
					},
					{
						Want: "testdata/getTickPrices.response.json",
						Actual: &GetTickPricesResponse{
							Status: true,
							ReturnData: GetTickPricesData{
								Quotations: []TickRecord{
									{
										Ask:         4000.0,
										AskVolume:   utils.Pointer(15000),
										Bid:         4000.0,
										BidVolume:   utils.Pointer(16000),
										High:        4000.0,
										Level:       0,
										Low:         3500,
										SpreadRaw:   0.000003,
										SpreadTable: 0.00042,
										Symbol:      "KOMB.CZ",
										Timestamp:   1272529161605,
									},
								},
							},
						},
					},
				},
			},
			{
				Name: CmdGetTradeRecords,
				Data: []testData{
					{
						Want: "testdata/getTradeRecords.request.json",
						Actual: NewGetTradeRecordsRequest(
							GetTradeRecordsRequestWithOrder(7489839),
							GetTradeRecordsRequestWithOrder(7489841),
						),
					},
					{
						Want: "testdata/getTradeRecords.response.json",
						Actual: &GetTradeRecordsResponse{
							Status: true,
							ReturnData: TradeRecords{
								{
									ClosePrice:       1.3256,
									CloseTime:        nil,
									CloseTimeString:  nil,
									Closed:           false,
									Cmd:              common.OperationCodeBuy,
									Comment:          "Web Trader",
									Commission:       0.0,
									CustomComment:    "Some text",
									Digits:           4,
									Expiration:       nil,
									ExpirationString: nil,
									MarginRate:       0,
									Offset:           0,
									OpenPrice:        1.4,
									OpenTime:         1272380927000,
									OpenTimeString:   "Fri Jan 11 10:03:36 CET 2013",
									Order:            7497776,
									Order2:           1234567,
									Position:         1234567,
									Profit:           -2196.44,
									StopLoss:         0.0,
									Storage:          -4.46,
									Symbol:           "EURUSD",
									Timestamp:        1272540251000,
									TakeProfit:       0.0,
									Volume:           0.1,
								},
							},
						},
					},
				},
			},
			{
				Name: CmdGetTrades,
				Data: []testData{
					{
						Want:   "testdata/getTrades.request.json",
						Actual: NewGetTradesRequest(GetTradesRequestWithOpenedOnly(true)),
					},
					{
						Want: "testdata/getTrades.response.json",
						Actual: &GetTradesResponse{
							Status: true,
							ReturnData: []TradeRecord{
								{
									ClosePrice:       1.3256,
									CloseTime:        nil,
									CloseTimeString:  nil,
									Closed:           false,
									Cmd:              common.OperationCodeBuy,
									Comment:          "Web Trader",
									Commission:       0,
									CustomComment:    "Some text",
									Digits:           4,
									Expiration:       nil,
									ExpirationString: nil,
									MarginRate:       0,
									Offset:           0,
									OpenPrice:        1.4,
									OpenTime:         1272380927000,
									OpenTimeString:   "Fri Jan 11 10:03:36 CET 2013",
									Order:            7497776,
									Order2:           1234567,
									Position:         1234567,
									Profit:           -2196.44,
									StopLoss:         0,
									Storage:          -4.46,
									Symbol:           "EURUSD",
									Timestamp:        1272540251000,
									TakeProfit:       0,
									Volume:           0.1,
								},
							},
						},
					},
				},
			},
			{
				Name: CmdGetTradesHistory,
				Data: []testData{
					{
						Want:   "testdata/getTradesHistory.request.json",
						Actual: NewGetTradesHistoryRequest(1275993488000, 0),
					},
					{
						Want: "testdata/getTradesHistory.response.json",
						Actual: &GetTradesHistoryResponse{
							Status: true,
							ReturnData: []TradeRecord{
								{
									ClosePrice:       1.3256,
									CloseTime:        nil,
									CloseTimeString:  nil,
									Closed:           false,
									Cmd:              common.OperationCodeBuy,
									Comment:          "Web Trader",
									Commission:       0,
									CustomComment:    "Some text",
									Digits:           4,
									Expiration:       nil,
									ExpirationString: nil,
									MarginRate:       0,
									Offset:           0,
									OpenPrice:        1.4,
									OpenTime:         1272380927000,
									OpenTimeString:   "Fri Jan 11 10:03:36 CET 2013",
									Order:            7497776,
									Order2:           1234567,
									Position:         1234567,
									Profit:           -2196.44,
									StopLoss:         0,
									Storage:          -4.46,
									Symbol:           "EURUSD",
									Timestamp:        1272540251000,
									TakeProfit:       0,
									Volume:           0.1,
								},
							},
						},
					},
				},
			},
			{
				Name: CmdGetTradingHours,
				Data: []testData{
					{
						Want: "testdata/getTradingHours.request.json",
						Actual: &GetTradingHoursRequest{
							Command: CmdGetTradingHours,
							Arguments: GetTradingHoursArgs{
								Symbols: []string{"EURPLN", "AGO.PL"},
							},
						},
					},
					{
						Want: "testdata/getTradingHours.response.json",
						Actual: &GetTradingHoursResponse{
							Status: true,
							ReturnData: TradingHoursRecords{
								{
									Quotes: []QuoteRecord{
										{
											Day:   2,
											FromT: 63000000,
											ToT:   63300000,
										},
									},
									Symbol: "USDPLN",
									Trading: []TradingRecord{
										{
											Day:   2,
											FromT: 63000000,
											ToT:   63300000,
										},
									},
								},
							},
						},
					},
				},
			},
			{
				Name: CmdGetVersion,
				Data: []testData{
					{
						Want:   "testdata/getVersion.request.json",
						Actual: NewGetVersionRequest(),
					},
					{
						Want: "testdata/getVersion.response.json",
						Actual: &GetVersionResponse{
							Status: true,
							ReturnData: VersionData{
								Version: "2.4.15",
							},
						},
					},
				},
			},
			{
				Name: CmdTradeTransaction,
				Data: []testData{
					{
						Want: "testdata/tradeTransaction.request.json",
						Actual: &TradeTransactionRequest{
							Command: "tradeTransaction",
							Arguments: TradeTransactionArgs{
								TradeTransInfo{
									Cmd:           common.OperationCodeBuyLimit,
									CustomComment: "Some text",
									Expiration:    1462006335000,
									Offset:        0,
									Order:         82188055,
									Price:         1.12,
									Symbol:        "EURUSD",
									Volume:        5.0,
								},
							},
						},
					},
					{
						Want: "testdata/tradeTransaction.response.json",
						Actual: &TradeTransactionResponse{
							Status: true,
							ReturnData: TradeTransactionData{
								Order: 43,
							},
						},
					},
				},
			},
			{
				Name: CmdTradeTransactionStatus,
				Data: []testData{
					{
						Want:   "testdata/tradeTransactionStatus.request.json",
						Actual: NewTradeTransactionStatusRequest(43),
					},
					{
						Want: "testdata/tradeTransactionStatus.response.json",
						Actual: &TradeTransactionStatusResponse{
							Status: true,
							ReturnData: TradeTransactionStatusData{
								Ask:           1.392,
								Bid:           1.392,
								CustomComment: "Some text",
								Message:       nil,
								Order:         43,
								RequestStatus: RequestStatusAccepted,
							},
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
