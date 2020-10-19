package main

import (
	"log"

	"github.com/yaustn/goftx"
)

func main() {
	client := goftx.NewClient()

	market, _ := client.GetMarket("BTC/USD")
	log.Printf("GetMarket BTC/USD: %+v", market)

	markets, _ := client.GetMarkets()
	log.Printf("GetMarkets: %+v", markets)
}
