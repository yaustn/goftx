package goftx

import (
	"encoding/json"
	"log"

	"github.com/yaustn/goftx/model"
)

const (
	getMarketsEndpoint = "/markets"
)

// GetMarket will take a market pair and return a Market's state.
// https://ftx.com/api/markets/BTC/USD
func (c *Client) GetMarket(marketName string) (market *model.Market, err error) {
	getMarketURL := apiURL + getMarketsEndpoint + "/" + marketName

	respBytes, err := c.get(getMarketURL)
	if err != nil {
		log.Printf("[ERROR] Failed GET /markets/%s request: %+v", marketName, err)
		return market, err
	}

	response := new(model.Response)
	response.Result = new(model.Market)

	err = json.Unmarshal(respBytes, response)
	if err != nil {
		log.Printf("[ERROR] Failed to unmarshal GET /markets response: %+v", err)
		return market, err
	}

	market = response.Result.(*model.Market)

	return market, nil
}

// GetMarkets returns a list of all Market states on FTX.
// https://ftx.com/api/markets
func (c *Client) GetMarkets() (markets *[]model.Market, err error) {
	getMarketsURL := apiURL + getMarketsEndpoint

	respBytes, err := c.get(getMarketsURL)
	if err != nil {
		log.Printf("[ERROR] Failed GET /markets request: %+v", err)
		return markets, err
	}

	response := new(model.Response)
	response.Result = new([]model.Market)

	err = json.Unmarshal(respBytes, response)
	if err != nil {
		log.Printf("[ERROR] Failed to unmarshal GET /markets response: %+v", err)
		return markets, err
	}

	markets = response.Result.(*[]model.Market)

	return markets, err
}

/*
func (c *Client) GetHistoricalMarket() {

}
*/
