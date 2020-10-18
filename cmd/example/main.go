package main

import (
	"log"

	"github.com/yaustn/goftx"
)

func main() {
	client := goftx.NewClient()

	market, _ := client.GetMarket("BTC/USD")
	log.Printf("GetMarkets: %+v", market)
}
