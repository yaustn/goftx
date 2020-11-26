package goftx

import "github.com/yaustn/goftx/model"

const (
	marketsEndpoint = "/markets"
)

func (c *Client) GetMarket(marketName string) (*model.Market, error) {
	var market model.Market
	err := c.get(marketsEndpoint+"/"+marketName, &market)
	if err != nil {
		return nil, err
	}

	return &market, nil
}

func (c *Client) GetMarkets() ([]model.Market, error) {
	var markets []model.Market
	err := c.get(marketsEndpoint, &markets)
	if err != nil {
		return nil, err
	}

	return markets, nil
}

// func (c *Client) GetOrderbook() () {
//
// }

// func (c *Client) GetTrades() () {
//
//}

// todo
// func (c *Client) GetHistoricalMarket() (*[]model.Market, error) {
//	 markets := new([]model.Market)
//
// }
