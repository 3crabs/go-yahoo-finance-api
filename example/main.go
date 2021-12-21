package main

import (
	"fmt"
	exchange "github.com/3crabs/go-yahoo-finance-api"
	"log"
)

func main() {
	pair := exchange.Pair{From: exchange.Ruble, To: exchange.DollarUSA}
	UsRu, err := exchange.GetCurrency(pair, "111")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(UsRu.QuoteResponse.Result[0].Bid)
	fmt.Println(UsRu.QuoteResponse.Result[0].ShortName)
	fmt.Println(UsRu.QuoteResponse.Result[0].Symbol)

	pairs := [...]exchange.Pair{
		{From: exchange.Ruble, To: exchange.Euro},
		{From: exchange.Euro, To: exchange.Ruble},
		{From: exchange.Ruble, To: exchange.DollarUSA},
		{From: exchange.DollarUSA, To: exchange.Ruble},
		{From: exchange.Euro, To: exchange.DollarUSA},
		{From: exchange.DollarUSA, To: exchange.Euro},
	}
	quote, err := exchange.GetCurrencies(pairs[:], "111")
	if err != nil {
		log.Println(err)
	}
	for _, pair := range quote.QuoteResponse.Result {
		fmt.Printf("Pair:%s -> Bid: %f\n", pair.ShortName, pair.Bid)
	}
}
