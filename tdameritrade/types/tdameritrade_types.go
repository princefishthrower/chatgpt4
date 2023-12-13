package tdamertrade

import "encoding/json"

type Credentials struct {
	UserId      string `json:"userid"`
	Token       string `json:"token"`
	Company     string `json:"company"`
	Segment     string `json:"segment"`
	CdDomain    string `json:"cddomain"`
	UserGroup   string `json:"usergroup"`
	AccessLevel string `json:"accesslevel"`
	Authorized  string `json:"authorized"`
	Timestamp   int64  `json:"timestamp"`
	AppID       string `json:"appID"`
	ACL         string `json:"acl"`
}

type StreamerInfo struct {
	StreamerBinaryUrl string `json:"streamerBinaryUrl"`
	StreamerSocketUrl string `json:"streamerSocketUrl"`
	Token             string `json:"token"`
	TokenTimestamp    string `json:"tokenTimestamp"`
	UserGroup         string `json:"userGroup"`
	AccessLevel       string `json:"accessLevel"`
	ACL               string `json:"acl"`
	AppID             string `json:"appID"`
}

type Quotes struct {
	IsNyseDelayed   bool `json:"isNyseDelayed"`
	IsNasdaqDelayed bool `json:"isNasdaqDelayed"`
	IsOpraDelayed   bool `json:"isOpraDelayed"`
	IsAmexDelayed   bool `json:"isAmexDelayed"`
	IsCmeDelayed    bool `json:"isCmeDelayed"`
	IsIceDelayed    bool `json:"isIceDelayed"`
	IsForexDelayed  bool `json:"isForexDelayed"`
}

type Key struct {
	Key string `json:"key"`
}

type StreamerSubscriptionKeys struct {
	Keys []Key `json:"keys"`
}

type ExchangeAgreements struct {
	NasdaqExchangeAgreement string `json:"NASDAQ_EXCHANGE_AGREEMENT"`
	NyseExchangeAgreement   string `json:"NYSE_EXCHANGE_AGREEMENT"`
	OpraExchangeAgreement   string `json:"OPRA_EXCHANGE_AGREEMENT"`
}

type Authorizations struct {
	Apex               bool   `json:"apex"`
	LevelTwoQuotes     bool   `json:"levelTwoQuotes"`
	StockTrading       bool   `json:"stockTrading"`
	MarginTrading      bool   `json:"marginTrading"`
	StreamingNews      bool   `json:"streamingNews"`
	OptionTradingLevel string `json:"optionTradingLevel"`
	ScottradeAccount   bool   `json:"scottradeAccount"`
	AutoPositionEffect bool   `json:"autoPositionEffect"`
}

type Account struct {
	AccountId         string         `json:"accountId"`
	DisplayName       string         `json:"displayName"`
	AccountCdDomainId string         `json:"accountCdDomainId"`
	Company           string         `json:"company"`
	Segment           string         `json:"segment"`
	Acl               string         `json:"acl"`
	Authorizations    Authorizations `json:"authorizations"`
}

type UserPrincipals struct {
	UserId                   string                   `json:"userId"`
	UserCdDomainId           string                   `json:"userCdDomainId"`
	PrimaryAccountId         string                   `json:"primaryAccountId"`
	LastLoginTime            string                   `json:"lastLoginTime"`
	TokenExpirationTime      string                   `json:"tokenExpirationTime"`
	LoginTime                string                   `json:"loginTime"`
	AccessLevel              string                   `json:"accessLevel"`
	StalePassword            bool                     `json:"stalePassword"`
	StreamerInfo             StreamerInfo             `json:"streamerInfo"`
	ProfessionalStatus       string                   `json:"professionalStatus"`
	Quotes                   Quotes                   `json:"quotes"`
	StreamerSubscriptionKeys StreamerSubscriptionKeys `json:"streamerSubscriptionKeys"`
	ExchangeAgreements       ExchangeAgreements       `json:"exchangeAgreements"`
	Accounts                 []Account                `json:"accounts"`
}

type RequestParameters struct {
	Credential string `json:"credential"`
	Token      string `json:"token"`
	Version    string `json:"version"`
}

type WebSocketRequest struct {
	Service    string            `json:"service"`
	Command    string            `json:"command"`
	RequestID  int               `json:"requestid"`
	Account    string            `json:"account"`
	Source     string            `json:"source"`
	Parameters map[string]string `json:"parameters"`
}

type WebSocketRequests struct {
	Requests []WebSocketRequest `json:"requests"`
}

type QuoteContent struct {
	Key         string  `json:"key"`
	Zero        string  `json:"0"`
	One         float64 `json:"1"`
	Two         float64 `json:"2"`
	Three       float64 `json:"3"`
	Four        float64 `json:"4"`
	Five        float64 `json:"5"`
	Six         string  `json:"6"`
	Seven       string  `json:"7"`
	Eight       float64 `json:"8"`
	Nine        float64 `json:"9"`
	Ten         float64 `json:"10"`
	Eleven      float64 `json:"11"`
	Twelve      float64 `json:"12"`
	Thirteen    float64 `json:"13"`
	Fourteen    string  `json:"14"`
	Fifteen     float64 `json:"15"`
	Sixteen     string  `json:"16"`
	Seventeen   bool    `json:"17"`
	Eighteen    bool    `json:"18"`
	Nineteen    float64 `json:"19"`
	Twenty      float64 `json:"20"`
	TwentyOne   float64 `json:"21"`
	TwentyTwo   float64 `json:"22"`
	TwentyThree float64 `json:"23"`
	TwentyFour  float64 `json:"24"`
	TwentyFive  string  `json:"25"`
	TwentySix   string  `json:"26"`
	TwentySeven float64 `json:"27"`
	TwentyEight float64 `json:"28"`
	TwentyNine  float64 `json:"29"`
	Thirty      float64 `json:"30"`
	ThirtyOne   float64 `json:"31"`
	ThirtyTwo   float64 `json:"32"`
	ThirtyThree float64 `json:"33"`
	ThirtyFour  float64 `json:"34"`
	ThirtyFive  float64 `json:"35"`
	ThirtySix   float64 `json:"36"`
	ThirtySeven float64 `json:"37"`
	ThirtyEight float64 `json:"38"`
	ThirtyNine  string  `json:"39"`
	Forty       string  `json:"40"`
	FortyOne    bool    `json:"41"`
	FortyTwo    bool    `json:"42"`
	FortyThree  float64 `json:"43"`
	FortyFour   float64 `json:"44"`
	FortyFive   float64 `json:"45"`
	FortySix    float64 `json:"46"`
	FortySeven  float64 `json:"47"`
	FortyEight  string  `json:"48"`
	FortyNine   float64 `json:"49"`
	Fifty       float64 `json:"50"`
	FiftyOne    float64 `json:"51"`
	FiftyTwo    int64   `json:"52"`
}

type ExchangeOrder struct {
	Zero string `json:"0"` // exchange name
	One  int64  `json:"1"` // size
}

type OptionOrder struct {
	Zero  string          `json:"0"` // price
	One   float64         `json:"1"` // size
	Two   string          `json:"2"` // exchange count
	Three []ExchangeOrder `json:"3"` // array exchanges
}

type OptionBookContent struct {
	Key   string        `json:"key"` // option ID
	One   int64         `json:"1"`   // quoteTimeInLong
	Two   []OptionOrder `json:"2"`   // buys
	Three []OptionOrder `json:"3"`   // sells
}

type ChartFuturesContent struct {
	Key    string  `json:"key"` // future ID
	Time   float64 `json:"1"`   // time
	Open   float64 `json:"2"`   // open
	High   float64 `json:"3"`   // high
	Low    float64 `json:"4"`   // low
	Close  float64 `json:"5"`   // close
	Volume float64 `json:"6"`   // volume
}

type StreamData struct {
	Service   string      `json:"service"`
	Timestamp int64       `json:"timestamp"`
	Command   string      `json:"command"`
	Content   interface{} `json:"content"` // left as string so we can dynamically unmarshal to various types
}

func (b *StreamData) UnmarshalJSON(data []byte) error {
	switch b.Service {
	case "OPTIONS_BOOK":
		b.Content = new([]OptionBookContent)
	case "CHART_FUTURES":
		b.Content = new([]ChartFuturesContent)
	case "QUOTE":
		fallthrough
	case "LEVELONE_FUTURES":
		fallthrough
	default:
		b.Content = new([]QuoteContent)
	}

	type tmp StreamData // avoids infinite recursion
	return json.Unmarshal(data, (*tmp)(b))
}

type StreamMessage struct {
	Data []StreamData `json:"data"`
}

// JSON schema for TD Ameritrade Option Chain
type OptionChain struct {
	// CallExpDateMap corresponds to the JSON schema field "callExpDateMap".
	CallExpDateMap OptionChainJsonCallExpDateMap `json:"callExpDateMap,omitempty"`

	// DaysToExpiration corresponds to the JSON schema field "daysToExpiration".
	DaysToExpiration *float64 `json:"daysToExpiration,omitempty"`

	// InterestRate corresponds to the JSON schema field "interestRate".
	InterestRate *float64 `json:"interestRate,omitempty"`

	// Interval corresponds to the JSON schema field "interval".
	Interval *float64 `json:"interval,omitempty"`

	// IsDelayed corresponds to the JSON schema field "isDelayed".
	IsDelayed *bool `json:"isDelayed,omitempty"`

	// IsIndex corresponds to the JSON schema field "isIndex".
	IsIndex *bool `json:"isIndex,omitempty"`

	// PutExpDateMap corresponds to the JSON schema field "putExpDateMap".
	PutExpDateMap OptionChainJsonPutExpDateMap `json:"putExpDateMap,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status *string `json:"status,omitempty"`

	// Strategy corresponds to the JSON schema field "strategy".
	Strategy *OptionChainJsonStrategy `json:"strategy,omitempty"`

	// Symbol corresponds to the JSON schema field "symbol".
	Symbol *string `json:"symbol,omitempty"`

	// Underlying corresponds to the JSON schema field "underlying".
	Underlying *OptionChainJsonUnderlying `json:"underlying,omitempty"`

	// UnderlyingPrice corresponds to the JSON schema field "underlyingPrice".
	UnderlyingPrice *float64 `json:"underlyingPrice,omitempty"`

	// Volatility corresponds to the JSON schema field "volatility".
	Volatility *float64 `json:"volatility,omitempty"`
}

type Option struct {
	PutCall                string          `json:"putCall"`
	Symbol                 string          `json:"symbol"`
	Description            string          `json:"description"`
	ExchangeName           string          `json:"exchangeName"`
	Bid                    float64         `json:"bid"`
	Ask                    float64         `json:"ask"`
	Last                   float64         `json:"last"`
	Mark                   float64         `json:"mark"`
	BidSize                int             `json:"bidSize"`
	AskSize                int             `json:"askSize"`
	BidAskSize             string          `json:"bidAskSize"`
	LastSize               int             `json:"lastSize"`
	HighPrice              float64         `json:"highPrice"`
	LowPrice               float64         `json:"lowPrice"`
	OpenPrice              float64         `json:"openPrice"`
	ClosePrice             float64         `json:"closePrice"`
	TotalVolume            int             `json:"totalVolume"`
	TradeDate              interface{}     `json:"tradeDate"`
	TradeTimeInLong        int64           `json:"tradeTimeInLong"`
	QuoteTimeInLong        int64           `json:"quoteTimeInLong"`
	NetChange              float64         `json:"netChange"`
	Volatility             json.RawMessage `json:"volatility"`
	Delta                  json.RawMessage `json:"delta"`
	Gamma                  json.RawMessage `json:"gamma"`
	Theta                  json.RawMessage `json:"theta"`
	Vega                   json.RawMessage `json:"vega"`
	Rho                    json.RawMessage `json:"rho"`
	OpenInterest           int             `json:"openInterest"`
	TimeValue              float64         `json:"timeValue"`
	TheoreticalOptionValue json.RawMessage `json:"theoreticalOptionValue"`
	TheoreticalVolatility  json.RawMessage `json:"theoreticalVolatility"`
	OptionDeliverablesList interface{}     `json:"optionDeliverablesList"`
	StrikePrice            float64         `json:"strikePrice"`
	ExpirationDate         int64           `json:"expirationDate"`
	DaysToExpiration       int             `json:"daysToExpiration"`
	ExpirationType         string          `json:"expirationType"`
	LastTradingDay         int64           `json:"lastTradingDay"`
	Multiplier             float64         `json:"multiplier"`
	SettlementType         string          `json:"settlementType"`
	DeliverableNote        string          `json:"deliverableNote"`
	IsIndexOption          interface{}     `json:"isIndexOption"`
	PercentChange          float64         `json:"percentChange"`
	MarkChange             float64         `json:"markChange"`
	MarkPercentChange      float64         `json:"markPercentChange"`
	IntrinsicValue         float64         `json:"intrinsicValue"`
	NonStandard            bool            `json:"nonStandard"`
	InTheMoney             bool            `json:"inTheMoney"`
	Mini                   bool            `json:"mini"`
}

type StrikePriceMap map[string][]Option

type OptionChainJsonCallExpDateMap map[string]StrikePriceMap

type OptionChainJsonPutExpDateMap map[string]StrikePriceMap

type OptionChainJsonStrategy string

type OptionChainJsonUnderlying struct {
	// Ask corresponds to the JSON schema field "ask".
	Ask *float64 `json:"ask,omitempty"`

	// AskSize corresponds to the JSON schema field "askSize".
	AskSize *int `json:"askSize,omitempty"`

	// Bid corresponds to the JSON schema field "bid".
	Bid *float64 `json:"bid,omitempty"`

	// BidSize corresponds to the JSON schema field "bidSize".
	BidSize *int `json:"bidSize,omitempty"`

	// Change corresponds to the JSON schema field "change".
	Change *float64 `json:"change,omitempty"`

	// Close corresponds to the JSON schema field "close".
	Close *float64 `json:"close,omitempty"`

	// Delayed corresponds to the JSON schema field "delayed".
	Delayed *bool `json:"delayed,omitempty"`

	// Description corresponds to the JSON schema field "description".
	Description *string `json:"description,omitempty"`

	// ExchangeName corresponds to the JSON schema field "exchangeName".
	ExchangeName *OptionChainJsonUnderlyingExchangeName `json:"exchangeName,omitempty"`

	// FiftyTwoWeekHigh corresponds to the JSON schema field "fiftyTwoWeekHigh".
	FiftyTwoWeekHigh *float64 `json:"fiftyTwoWeekHigh,omitempty"`

	// FiftyTwoWeekLow corresponds to the JSON schema field "fiftyTwoWeekLow".
	FiftyTwoWeekLow *float64 `json:"fiftyTwoWeekLow,omitempty"`

	// HighPrice corresponds to the JSON schema field "highPrice".
	HighPrice *float64 `json:"highPrice,omitempty"`

	// Last corresponds to the JSON schema field "last".
	Last *float64 `json:"last,omitempty"`

	// LowPrice corresponds to the JSON schema field "lowPrice".
	LowPrice *float64 `json:"lowPrice,omitempty"`

	// Mark corresponds to the JSON schema field "mark".
	Mark *float64 `json:"mark,omitempty"`

	// MarkChange corresponds to the JSON schema field "markChange".
	MarkChange *float64 `json:"markChange,omitempty"`

	// MarkPercentChange corresponds to the JSON schema field "markPercentChange".
	MarkPercentChange *float64 `json:"markPercentChange,omitempty"`

	// OpenPrice corresponds to the JSON schema field "openPrice".
	OpenPrice *float64 `json:"openPrice,omitempty"`

	// PercentChange corresponds to the JSON schema field "percentChange".
	PercentChange *float64 `json:"percentChange,omitempty"`

	// QuoteTime corresponds to the JSON schema field "quoteTime".
	QuoteTime *int `json:"quoteTime,omitempty"`

	// Symbol corresponds to the JSON schema field "symbol".
	Symbol *string `json:"symbol,omitempty"`

	// TotalVolume corresponds to the JSON schema field "totalVolume".
	TotalVolume *int `json:"totalVolume,omitempty"`

	// TradeTime corresponds to the JSON schema field "tradeTime".
	TradeTime *int `json:"tradeTime,omitempty"`
}

type OptionChainJsonUnderlyingExchangeName string

type OHLCV struct {
	O float64 `json:"o"`
	H float64 `json:"h"`
	L float64 `json:"l"`
	C float64 `json:"c"`
	V float64 `json:"v"`
}
