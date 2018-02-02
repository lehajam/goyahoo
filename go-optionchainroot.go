package goyahoo

type OptionChainRoot struct {
	OptionChain Header `json:"optionChain"`
}

type Underlying struct {
	Ask                               float64 `json:"ask"`
	AskSize                           float64 `json:"askSize"`
	AverageDailyVolume10Day           float64 `json:"averageDailyVolume10Day"`
	AverageDailyVolume3Month          float64 `json:"averageDailyVolume3Month"`
	Bid                               float64 `json:"bid"`
	BidSize                           float64 `json:"bidSize"`
	BookValue                         float64 `json:"bookValue"`
	Currency                          string  `json:"currency"`
	DividendDate                      int64   `json:"dividendDate"`
	EarningsTimestamp                 int64   `json:"earningsTimestamp"`
	EarningsTimestampEnd              int64   `json:"earningsTimestampEnd"`
	EarningsTimestampStart            int64   `json:"earningsTimestampStart"`
	EpsForward                        float64 `json:"epsForward"`
	EpsTrailingTwelveMonths           float64 `json:"epsTrailingTwelveMonths"`
	Exchange                          string  `json:"exchange"`
	ExchangeTimezoneName              string  `json:"exchangeTimezoneName"`
	ExchangeTimezoneShortName         string  `json:"exchangeTimezoneShortName"`
	FiftyDayAverage                   float64 `json:"fiftyDayAverage"`
	FiftyDayAverageChange             float64 `json:"fiftyDayAverageChange"`
	FiftyDayAverageChangePercent      float64 `json:"fiftyDayAverageChangePercent"`
	FiftyTwoWeekHigh                  float64 `json:"fiftyTwoWeekHigh"`
	FiftyTwoWeekHighChange            float64 `json:"fiftyTwoWeekHighChange"`
	FiftyTwoWeekHighChangePercent     float64 `json:"fiftyTwoWeekHighChangePercent"`
	FiftyTwoWeekLow                   float64 `json:"fiftyTwoWeekLow"`
	FiftyTwoWeekLowChange             float64 `json:"fiftyTwoWeekLowChange"`
	FiftyTwoWeekLowChangePercent      float64 `json:"fiftyTwoWeekLowChangePercent"`
	ForwardPE                         float64 `json:"forwardPE"`
	FullExchangeName                  string  `json:"fullExchangeName"`
	GmtOffSetMilliseconds             int64   `json:"gmtOffSetMilliseconds"`
	LongName                          string  `json:"longName"`
	Market                            string  `json:"market"`
	MarketCap                         float64 `json:"marketCap"`
	MarketState                       string  `json:"marketState"`
	MessageBoardID                    string  `json:"messageBoardId"`
	PostMarketChange                  float64 `json:"postMarketChange"`
	PostMarketChangePercent           float64 `json:"postMarketChangePercent"`
	PostMarketPrice                   float64 `json:"postMarketPrice"`
	PostMarketTime                    int64   `json:"postMarketTime"`
	PriceHint                         float64 `json:"priceHint"`
	PriceToBook                       float64 `json:"priceToBook"`
	QuoteSourceName                   string  `json:"quoteSourceName"`
	QuoteType                         string  `json:"quoteType"`
	RegularMarketChange               float64 `json:"regularMarketChange"`
	RegularMarketChangePercent        float64 `json:"regularMarketChangePercent"`
	RegularMarketDayHigh              float64 `json:"regularMarketDayHigh"`
	RegularMarketDayLow               float64 `json:"regularMarketDayLow"`
	RegularMarketOpen                 float64 `json:"regularMarketOpen"`
	RegularMarketPreviousClose        float64 `json:"regularMarketPreviousClose"`
	RegularMarketPrice                float64 `json:"regularMarketPrice"`
	RegularMarketTime                 int64   `json:"regularMarketTime"`
	RegularMarketVolume               float64 `json:"regularMarketVolume"`
	SharesOutstanding                 float64 `json:"sharesOutstanding"`
	ShortName                         string  `json:"shortName"`
	SourceInterval                    float64 `json:"sourceInterval"`
	Symbol                            string  `json:"symbol"`
	TrailingAnnualDividendRate        float64 `json:"trailingAnnualDividendRate"`
	TrailingAnnualDividendYield       float64 `json:"trailingAnnualDividendYield"`
	TrailingPE                        float64 `json:"trailingPE"`
	TwoHundredDayAverage              float64 `json:"twoHundredDayAverage"`
	TwoHundredDayAverageChange        float64 `json:"twoHundredDayAverageChange"`
	TwoHundredDayAverageChangePercent float64 `json:"twoHundredDayAverageChangePercent"`
}

type Quote struct {
	Ask               float64 `json:"ask"`
	Bid               float64 `json:"bid"`
	Change            float64 `json:"change"`
	ContractSize      string  `json:"contractSize"`
	ContractSymbol    string  `json:"contractSymbol"`
	Currency          string  `json:"currency"`
	Expiration        float64 `json:"expiration"`
	ImpliedVolatility float64 `json:"impliedVolatility"`
	InTheMoney        bool    `json:"inTheMoney"`
	LastPrice         float64 `json:"lastPrice"`
	LastTradeDate     int64   `json:"lastTradeDate"`
	OpenInterest      float64 `json:"openInterest"`
	PercentChange     float64 `json:"percentChange"`
	Strike            float64 `json:"strike"`
	Volume            float64 `json:"volume"`
}

type Quotes struct {
	Calls          []Quote `json:"calls"`
	ExpirationDate int64   `json:"expirationDate"`
	HasMiniOptions bool    `json:"hasMiniOptions"`
	Puts           []Quote `json:"puts"`
}

type Header struct {
	Error  interface{} `json:"error"`
	Result []Chain     `json:"result"`
}

type Chain struct {
	ExpirationDates  []int64    `json:"expirationDates"`
	HasMiniOptions   bool       `json:"hasMiniOptions"`
	Options          []Quotes   `json:"options"`
	Quote            Underlying `json:"quote"`
	Strikes          []float64  `json:"strikes"`
	UnderlyingSymbol string     `json:"underlyingSymbol"`
}
