package exchange

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://yfapi.net/v6/finance/quote?region=US&lang=en"

type Quote struct {
	Region    string
	QuoteType string
	ShortName string
	Bid       float32
	Ask       float32
}

func GetCurrency(fromCurrency string, toCurrency string, apiKey string) Quote {
	currencyParam := "&symbols=" + fromCurrency + toCurrency + "%3DX"
	finalUrl := url + currencyParam
	req, _ := http.NewRequest("GET", finalUrl, nil)
	req.Header.Set("accept", "application/json")
	req.Header.Set("X-API-KEY", apiKey)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))

	return Quote{
		Region:    "US",
		QuoteType: "CURRENCY",
		ShortName: "USD/RUB",
		Bid:       74.19673,
		Ask:       74.20387,
	}
}
