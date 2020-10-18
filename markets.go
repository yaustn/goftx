package goftx

import (
	"encoding/json"
	"log"

	"github.com/yaustn/goftx/model"
)

const (
	getMarketEndpoint = "markets/"
)

// GetMarket will take a market pair and return a Market's state.
// ex: https://ftx.com/api/markets/BTC/USD
func (c *Client) GetMarket(marketName string) (market *model.Market, err error) {
	getMarketURL := url + getMarketEndpoint + marketName

	respBytes, err := c.get(getMarketURL)
	if err != nil {
		log.Printf("[ERROR] Failed GET /markets request: %+v", err)
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
// ex: https://ftx.com/api/markets
func (c *Client) GetMarkets() (markets []model.Market, err error) {
	// todo

	return markets, err
}
