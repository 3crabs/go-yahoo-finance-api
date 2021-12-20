package main

import (
	"fmt"
	exchange "github.com/3crabs/go-yahoo-finance-api"
	"log"
)

func main() {
	UsRu, err := exchange.GetCurrency("RUB", "USD", "124")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(UsRu.QuoteResponse.Result[0].Bid)
	fmt.Println(UsRu.QuoteResponse.Result[0].ShortName)
	fmt.Println(UsRu.QuoteResponse.Result[0].Symbol)
}
