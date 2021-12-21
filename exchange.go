package exchange

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Quote struct {
	QuoteResponse struct {
		Result []struct {
			Language                          string  `json:"language"`
			Region                            string  `json:"region"`
			QuoteType                         string  `json:"quoteType"`
			QuoteSourceName                   string  `json:"quoteSourceName"`
			Triggerable                       bool    `json:"triggerable"`
			Currency                          string  `json:"currency"`
			FiftyTwoWeekLow                   float64 `json:"fiftyTwoWeekLow"`
			FiftyTwoWeekHigh                  float64 `json:"fiftyTwoWeekHigh"`
			FiftyDayAverage                   float64 `json:"fiftyDayAverage"`
			TwoHundredDayAverage              float64 `json:"twoHundredDayAverage"`
			RegularMarketChange               float64 `json:"regularMarketChange"`
			RegularMarketChangePercent        float64 `json:"regularMarketChangePercent"`
			RegularMarketPrice                float64 `json:"regularMarketPrice"`
			RegularMarketDayHigh              float64 `json:"regularMarketDayHigh"`
			RegularMarketDayLow               float64 `json:"regularMarketDayLow"`
			RegularMarketPreviousClose        float64 `json:"regularMarketPreviousClose"`
			Bid                               float64 `json:"bid"`
			Ask                               float64 `json:"ask"`
			RegularMarketOpen                 float64 `json:"regularMarketOpen"`
			ShortName                         string  `json:"shortName"`
			FiftyTwoWeekRange                 string  `json:"fiftyTwoWeekRange"`
			FiftyTwoWeekHighChange            float64 `json:"fiftyTwoWeekHighChange"`
			FiftyTwoWeekHighChangePercent     float64 `json:"fiftyTwoWeekHighChangePercent"`
			FiftyDayAverageChange             float64 `json:"fiftyDayAverageChange"`
			FiftyDayAverageChangePercent      float64 `json:"fiftyDayAverageChangePercent"`
			TwoHundredDayAverageChange        float64 `json:"twoHundredDayAverageChange"`
			TwoHundredDayAverageChangePercent float64 `json:"twoHundredDayAverageChangePercent"`
			SourceInterval                    int     `json:"sourceInterval"`
			ExchangeDataDelayedBy             int     `json:"exchangeDataDelayedBy"`
			Tradeable                         bool    `json:"tradeable"`
			FirstTradeDateMilliseconds        int64   `json:"firstTradeDateMilliseconds"`
			PriceHint                         int     `json:"priceHint"`
			MessageBoardId                    string  `json:"messageBoardId"`
			ExchangeTimezoneName              string  `json:"exchangeTimezoneName"`
			ExchangeTimezoneShortName         string  `json:"exchangeTimezoneShortName"`
			GmtOffSetMilliseconds             int     `json:"gmtOffSetMilliseconds"`
			Market                            string  `json:"market"`
			EsgPopulated                      bool    `json:"esgPopulated"`
			Exchange                          string  `json:"exchange"`
			MarketState                       string  `json:"marketState"`
			RegularMarketTime                 int     `json:"regularMarketTime"`
			RegularMarketDayRange             string  `json:"regularMarketDayRange"`
			RegularMarketVolume               int     `json:"regularMarketVolume"`
			BidSize                           int     `json:"bidSize"`
			AskSize                           int     `json:"askSize"`
			FullExchangeName                  string  `json:"fullExchangeName"`
			AverageDailyVolume3Month          int     `json:"averageDailyVolume3Month"`
			AverageDailyVolume10Day           int     `json:"averageDailyVolume10Day"`
			FiftyTwoWeekLowChange             float64 `json:"fiftyTwoWeekLowChange"`
			FiftyTwoWeekLowChangePercent      float64 `json:"fiftyTwoWeekLowChangePercent"`
			Symbol                            string  `json:"symbol"`
		} `json:"result"`
		Error interface{} `json:"error"`
	} `json:"quoteResponse"`
}

const (
	Euro      = "EUR"
	DollarUSA = "USD"
	Ruble     = "RUB"
)

type Pair struct {
	From string
	To   string
}

func (p *Pair) ToString() string {
	return p.From + p.To + "%3DX"
}

func GetCurrencies(pairList []Pair, apiKey string) (*Quote, error) {
	var symbols string
	for _, pair := range pairList {
		v := "," + pair.ToString()
		symbols += v
	}
	if len(symbols) > 0 {
		symbols = strings.TrimPrefix(symbols, ",")
	}
	quote, err := sendYahooRequest(symbols, apiKey)
	if err != nil {
		return nil, err
	}
	return quote, nil
}

func GetCurrency(pair Pair, apiKey string) (*Quote, error) {
	symbols := pair.ToString()
	quote, err := sendYahooRequest(symbols, apiKey)
	if err != nil {
		return nil, err
	}
	return quote, nil
}

func sendYahooRequest(symbols string, apiKey string) (*Quote, error) {
	v := url.Values{
		"region":  []string{"US"},
		"lang":    []string{"en"},
		"symbols": []string{symbols},
	}

	targetUrl := url.URL{
		Scheme:   "https",
		Host:     "yfapi.net",
		Path:     "v6/finance/quote",
		RawQuery: v.Encode(),
	}

	req, _ := http.NewRequest("GET", targetUrl.String(), nil)
	req.Header.Set("accept", "application/json")
	req.Header.Set("X-API-KEY", apiKey)
	res, _ := http.DefaultClient.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("file doesn't closed")
		}
	}(res.Body)
	body, _ := ioutil.ReadAll(res.Body)

	result := &Quote{}
	err := json.Unmarshal(body, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
