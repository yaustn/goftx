package goftx

import (
	"encoding/json"

	"github.com/yaustn/goftx/model"
)

const (
	marketsEndpoint = "/markets"
)

func (c *Client) GetMarket(marketName string) (market *model.Market, err error) {
	respBytes, err := c.get(marketsEndpoint + "/" + marketName)
	if err != nil {
		return market, err
	}

	response := new(model.Response)
	response.Result = new(model.Market)

	err = json.Unmarshal(respBytes, response)
	if err != nil {
		return market, err
	}

	market = response.Result.(*model.Market)

	return market, nil
}

func (c *Client) GetMarkets() (markets *[]model.Market, err error) {
	respBytes, err := c.get(marketsEndpoint)
	if err != nil {
		return markets, err
	}

	response := new(model.Response)
	response.Result = new([]model.Market)

	err = json.Unmarshal(respBytes, response)
	if err != nil {
		return markets, err
	}

	markets = response.Result.(*[]model.Market)

	return markets, nil
}

/*
func (c *Client) GetHistoricalMarket() {

}
*/
