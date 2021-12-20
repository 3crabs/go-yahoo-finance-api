package main

import exchange "github.com/3crabs/go-yahoo-finance-api"

func main() {
	exchange.GetCurrency("USD", "RUS", "111")
}
