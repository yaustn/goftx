package main

import (
	"io/ioutil"
	"log"

	"github.com/yaustn/goftx"
	"gopkg.in/yaml.v2"
)

type apiSecret struct {
	Key    string `yaml:"ftx-api-key"`
	Secret string `yaml:"ftx-api-secret"`
}

func main() {
	var secrets apiSecret
	secretsFile, _ := ioutil.ReadFile("config")
	_ = yaml.Unmarshal(secretsFile, &secrets)

	client := goftx.NewClient(secrets.Secret, secrets.Key)

	market, _ := client.GetMarket("BTC/USD")
	log.Printf("GetMarket BTC/USD: %+v", market)

	markets, _ := client.GetMarkets()
	log.Printf("GetMarkets: %+v", markets)

	balance, _ := client.GetBalances()
	log.Printf("GetBalance: %+v", balance)
}
