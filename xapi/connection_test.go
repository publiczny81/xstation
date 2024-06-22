package xapi

import (
	"encoding/json"
	"fmt"
	"github.com/publiczny81/xstation/xapi/model"
	"github.com/publiczny81/xstation/xapi/websocket"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

var (
	accountId     = os.Getenv(XStationUserId)
	password      = os.Getenv(XStationPassword)
	address       = os.Getenv(XStationUrl)
	streamAddress = os.Getenv(XStationStreamUrl)
	origin        = os.Getenv(XStationOrigin)
)

func TestConnect(t *testing.T) {
	conn, err := websocket.Connect(address, origin)
	assert.NoError(t, err)

	loginResponse, err := Login(conn, accountId, password)
	assert.NoError(t, err)
	buf, _ := json.Marshal(loginResponse)
	fmt.Println(string(buf))
	getAllResponse, err := GetAllSymbols(conn)
	assert.NoError(t, err)
	groups := make(map[string][]model.SymbolRecord)
	for _, s := range getAllResponse.ReturnData {
		groups[s.GroupName] = append(groups[s.GroupName], s)
	}
	//for _, s := range groups["Poland"] {
	//	value, _ := json.Marshal(s)
	//	fmt.Println(string(value))
	//}
	getCalendarResponse, err := GetCalendar(conn)
	for _, r := range getCalendarResponse.ReturnData {
		buf, _ := json.Marshal(r)
		fmt.Println(string(buf))
	}
	serverTimeResponse, err := GetServerTime(conn)
	if err != nil {
		return
	}
	fmt.Println(serverTimeResponse.ReturnData.Time, serverTimeResponse.ReturnData.TimeString)
	fmt.Println(time.UnixMilli(int64(serverTimeResponse.ReturnData.Time)))
	resp, err := Logout(conn)
	if resp.Status {

	}
}
