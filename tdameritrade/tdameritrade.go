package tdameritrade

import (
	tdameritradeTypes "chatgpt4/tdameritrade/types"
	"chatgpt4/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

// converts milliseconds since epoch to timestamp string
func MillisecondsToTimestampString(milliseconds int64) string {
	t := time.Unix(0, milliseconds*int64(time.Millisecond))
	return t.Format("2006-01-02 15:04:05")
}

func CreateTDAmeritradeSocket(streamerSocketUrl string) (*websocket.Conn, error) {
	addr := streamerSocketUrl
	url := url.URL{Scheme: "wss", Host: addr, Path: "/ws"}
	log.Printf("connecting to %s", url.String())

	websocketConn, _, err := websocket.DefaultDialer.Dial(url.String(), nil)

	if err != nil {
		log.Fatal("dial:", err)
		return nil, err
	}

	return websocketConn, nil
}

func ConnectToTDAmeritradeWithExpressConnection(userPrincipalsString string) (*websocket.Conn, *int, *tdameritradeTypes.UserPrincipals, error) {
	userPrincipals := tdameritradeTypes.UserPrincipals{}

	// marshall into a struct
	err := json.Unmarshal([]byte(userPrincipalsString), &userPrincipals)
	if err != nil {
		log.Println(err)
		return nil, nil, nil, err
	}

	// create TD Ameritrade socket connection
	conn, err := CreateTDAmeritradeSocket(userPrincipals.StreamerInfo.StreamerSocketUrl)
	if err != nil {
		log.Println(err)
		return nil, nil, nil, err
	}

	// login to the socket connection
	requestId := 1
	err = Login(requestId, conn, userPrincipals)
	if err != nil {
		log.Println(err)
		return nil, nil, nil, err
	}

	// read message from login
	_, message, err := conn.ReadMessage()
	if err != nil {
		log.Println("Error reading login message:", err)
	}

	// print message to console
	log.Printf("recv: %s", message)

	// set express connection on the socket connection
	requestId += 1
	err = SetExpressConnection(requestId, conn, userPrincipals)
	if err != nil {
		log.Println(err)
		return nil, nil, nil, err
	}

	// read message from setting express connection
	_, message, err = conn.ReadMessage()
	if err != nil {
		log.Println("Error reading express connection message:", err)
	}

	// print message to console
	log.Printf("recv: %s", message)

	return conn, &requestId, &userPrincipals, nil
}

func Login(requestID int, conn *websocket.Conn, userPrincipals tdameritradeTypes.UserPrincipals) error {

	tokenTimeStampAsDateObj, err := time.Parse("2006-01-02T15:04:05-0700", userPrincipals.StreamerInfo.TokenTimestamp)
	if err != nil {
		log.Println(err)
		return err
	}
	tokenTimeStampAsMs := tokenTimeStampAsDateObj.UnixNano() / int64(time.Millisecond)

	credentials := tdameritradeTypes.Credentials{
		UserId:      userPrincipals.Accounts[0].AccountId,
		Token:       userPrincipals.StreamerInfo.Token,
		Company:     userPrincipals.Accounts[0].Company,
		Segment:     userPrincipals.Accounts[0].Segment,
		CdDomain:    userPrincipals.Accounts[0].AccountCdDomainId,
		UserGroup:   userPrincipals.StreamerInfo.UserGroup,
		AccessLevel: userPrincipals.StreamerInfo.AccessLevel,
		Authorized:  "Y",
		Timestamp:   tokenTimeStampAsMs,
		AppID:       userPrincipals.StreamerInfo.AppID,
		ACL:         userPrincipals.StreamerInfo.ACL,
	}

	webSocketRequests := tdameritradeTypes.WebSocketRequests{
		Requests: []tdameritradeTypes.WebSocketRequest{
			{
				Service:   "ADMIN",
				Command:   "LOGIN",
				RequestID: requestID,
				Account:   userPrincipals.Accounts[0].AccountId,
				Source:    userPrincipals.StreamerInfo.AppID,
				Parameters: map[string]string{
					"credential": jsonToQueryString(credentials),
					"token":      userPrincipals.StreamerInfo.Token,
					"version":    "1.0",
				},
			},
		},
	}

	// serialize webSocketRequests to json
	webSocketRequestsJson, err := json.Marshal(webSocketRequests)
	if err != nil {
		return err
	}

	// go turns ampersands into \u0026, so we need to unescape them
	webSocketRequestsJson, err = _UnescapeUnicodeCharactersInJSON(webSocketRequestsJson)

	// print json to console
	log.Printf("LOGIN send: %s", webSocketRequestsJson)

	err = conn.WriteMessage(websocket.TextMessage, webSocketRequestsJson)
	if err != nil {
		return err
	}

	return nil
}

func _UnescapeUnicodeCharactersInJSON(_jsonRaw json.RawMessage) (json.RawMessage, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(_jsonRaw)), `\\u`, `\u`, -1))
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}

func SetExpressConnection(requestID int, conn *websocket.Conn, userPrincipals tdameritradeTypes.UserPrincipals) error {
	webSocketRequests := tdameritradeTypes.WebSocketRequests{
		Requests: []tdameritradeTypes.WebSocketRequest{
			{
				Service:   "ADMIN",
				Command:   "QOS",
				RequestID: requestID,
				Account:   userPrincipals.Accounts[0].AccountId,
				Source:    userPrincipals.StreamerInfo.AppID,
				Parameters: map[string]string{
					"qoslevel": "0",
				},
			},
		},
	}

	// serialize webSocketRequests to json
	webSocketRequestsJson, err := json.Marshal(webSocketRequests)
	if err != nil {
		return err
	}

	err = conn.WriteMessage(websocket.TextMessage, webSocketRequestsJson)
	if err != nil {
		return err
	}

	return nil

}

func StreamFuturesQuotes(symbol string, requestID int, conn *websocket.Conn, userPrincipals tdameritradeTypes.UserPrincipals) error {
	webSocketRequests := tdameritradeTypes.WebSocketRequests{
		Requests: []tdameritradeTypes.WebSocketRequest{
			{
				Service:   "LEVELONE_FUTURES",
				Command:   "SUBS",
				RequestID: 1,
				Account:   userPrincipals.Accounts[0].AccountId,
				Source:    userPrincipals.StreamerInfo.AppID,
				Parameters: map[string]string{
					"keys":   symbol,
					"fields": "0,1,2,3",
				},
			},
		},
	}

	// serialize webSocketRequests to json
	webSocketRequestsJson, err := json.Marshal(webSocketRequests)
	if err != nil {
		return err
	}

	err = conn.WriteMessage(websocket.TextMessage, webSocketRequestsJson)
	if err != nil {
		return err
	}

	return nil
}

func StreamStockQuotes(symbol string, requestID int, conn *websocket.Conn, userPrincipals tdameritradeTypes.UserPrincipals) error {
	webSocketRequests := tdameritradeTypes.WebSocketRequests{
		Requests: []tdameritradeTypes.WebSocketRequest{
			{
				Service:   "QUOTE",
				Command:   "SUBS",
				RequestID: 1,
				Account:   userPrincipals.Accounts[0].AccountId,
				Source:    userPrincipals.StreamerInfo.AppID,
				Parameters: map[string]string{
					"keys":   symbol,
					"fields": "0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52",
				},
			},
		},
	}

	// serialize webSocketRequests to json
	webSocketRequestsJson, err := json.Marshal(webSocketRequests)
	if err != nil {
		return err
	}

	err = conn.WriteMessage(websocket.TextMessage, webSocketRequestsJson)
	if err != nil {
		return err
	}

	return nil
}

func StreamOptionsBook(symbol string, requestID int, conn *websocket.Conn, userPrincipals tdameritradeTypes.UserPrincipals) error {
	webSocketRequests := tdameritradeTypes.WebSocketRequests{
		Requests: []tdameritradeTypes.WebSocketRequest{
			{
				Service:   "OPTIONS_BOOK",
				Command:   "SUBS",
				RequestID: 1,
				Account:   userPrincipals.Accounts[0].AccountId,
				Source:    userPrincipals.StreamerInfo.AppID,
				Parameters: map[string]string{
					"keys":   symbol,
					"fields": "0,1,2,3,4",
				},
			},
		},
	}

	// serialize webSocketRequests to json
	webSocketRequestsJson, err := json.Marshal(webSocketRequests)
	if err != nil {
		return err
	}

	err = conn.WriteMessage(websocket.TextMessage, webSocketRequestsJson)
	if err != nil {
		return err
	}

	return nil
}

func StreamFuturesChart(symbol string, requestID int, conn *websocket.Conn, userPrincipals tdameritradeTypes.UserPrincipals) error {
	webSocketRequests := tdameritradeTypes.WebSocketRequests{
		Requests: []tdameritradeTypes.WebSocketRequest{
			{
				Service:   "CHART_FUTURES",
				Command:   "SUBS",
				RequestID: 1,
				Account:   userPrincipals.Accounts[0].AccountId,
				Source:    userPrincipals.StreamerInfo.AppID,
				Parameters: map[string]string{
					"keys":   symbol,
					"fields": "0,1,2,3,4,5,6",
				},
			},
		},
	}

	// serialize webSocketRequests to json
	webSocketRequestsJson, err := json.Marshal(webSocketRequests)
	if err != nil {
		return err
	}

	err = conn.WriteMessage(websocket.TextMessage, webSocketRequestsJson)
	if err != nil {
		return err
	}

	return nil
}

func GetChartHistoryFutures(symbol string, frequency string, period string, requestID int, conn *websocket.Conn, userPrincipals tdameritradeTypes.UserPrincipals) error {
	webSocketRequests := tdameritradeTypes.WebSocketRequests{
		Requests: []tdameritradeTypes.WebSocketRequest{
			{
				Service:   "CHART_HISTORY_FUTURES",
				Command:   "GET",
				RequestID: 1,
				Account:   userPrincipals.Accounts[0].AccountId,
				Source:    userPrincipals.StreamerInfo.AppID,
				Parameters: map[string]string{
					"symbol":    symbol,
					"frequency": frequency,
					"period":    period,
				},
			},
		},
	}

	// serialize webSocketRequests to json
	webSocketRequestsJson, err := json.Marshal(webSocketRequests)
	if err != nil {
		return err
	}

	err = conn.WriteMessage(websocket.TextMessage, webSocketRequestsJson)
	if err != nil {
		return err
	}

	return nil
}

func jsonToQueryString(json tdameritradeTypes.Credentials) string {
	var queryParams []string

	v := reflect.ValueOf(json)

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		tag := field.Tag.Get("json")
		value := fmt.Sprintf("%v", v.Field(i))
		queryParams = append(queryParams, url.QueryEscape(tag)+"="+url.QueryEscape(fmt.Sprintf("%v", value)))
	}

	return strings.Join(queryParams, "&")
}

func GetTDAmeritradeOptionChain(symbol string) (optionChain *tdameritradeTypes.OptionChain, error error) {
	// the url to call
	requestUrl := "https://api.tdameritrade.com/v1/marketdata/chains"

	// the query parameters to pass
	queryParameters := url.Values{}
	queryParameters.Add("apikey", os.Getenv("TD_AMERITRADE_API_KEY"))
	queryParameters.Add("includeQuotes", "TRUE")
	queryParameters.Add("symbol", symbol)

	// the type to unmarshal the response into the option chain type
	response := tdameritradeTypes.OptionChain{}

	// call the function
	response, err := utils.MakeHTTPRequest(requestUrl, "GET", nil, queryParameters, nil, response)
	if err != nil {
		// don't panic, just return the error
		return nil, err
	}

	return &response, nil
}

func GetAllOptionsInOptionsChain(optionChain tdameritradeTypes.OptionChain) ([]tdameritradeTypes.Option, []int64) {
	// first initialize an empty slice of options
	options := []tdameritradeTypes.Option{}
	expirations := []int64{}

	// first do calls
	for _, expiries := range optionChain.CallExpDateMap {
		for _, expiryStrikes := range expiries {
			// append the option to the slice
			options = append(options, expiryStrikes...)

			// append the expiration to the slice
			expirations = append(expirations, expiryStrikes[0].ExpirationDate)
		}
	}

	// then do puts
	for _, expiries := range optionChain.PutExpDateMap {
		for _, expiryStrikes := range expiries {
			// append the option to the slice
			options = append(options, expiryStrikes...)

			// append the expiration to the slice
			expirations = append(expirations, expiryStrikes[0].ExpirationDate)
		}
	}

	// ensure that expirations are unique
	uniqueExpirations := utils.Unique(expirations)

	// sort the expirations
	sort.Slice(uniqueExpirations, func(i, j int) bool { return uniqueExpirations[i] < uniqueExpirations[j] })

	return options, uniqueExpirations
}

func GetCommaSeparatedOptionIDs(options []tdameritradeTypes.Option, filter string) string {
	// first initialize an empty slice of option ids
	optionIDs := []string{}

	// loop through the options and append the option id to the slice
	for _, option := range options {
		optionIDs = append(optionIDs, option.Symbol)
	}

	// if there is a filter, then filter the option ids
	if filter != "" {
		optionIDs = utils.Filter(optionIDs, func(optionID string) bool {
			return strings.Contains(optionID, filter)
		})
	}

	// join the option ids with a comma
	return strings.Join(optionIDs, ",")
}
