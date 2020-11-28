package goftx

const (
	marketsEndpoint = "/markets"
)

type Market struct {
	Type           string  `json:"type"`
	Name           string  `json:"name"`
	BaseCurrency   string  `json:"baseCurrency"`
	QuoteCurrency  string  `json:"quoteCurrency"`
	Underlying     string  `json:"underlying"`
	VolumeUsd24H   float64 `json:"volumeUsd24h"`
	QuoteVolume24H float64 `json:"quoteVolume24h"`
	Price          float64 `json:"price"`
	Ask            float64 `json:"ask"`
	Bid            float64 `json:"bid"`
	Last           float64 `json:"last"`
	PriceIncrement float64 `json:"priceIncrement"`
	SizeIncrement  float64 `json:"sizeIncrement"`
	MinProvideSize float64 `json:"minProvideSize"`
	Change1H       float64 `json:"change1h"`
	Change24H      float64 `json:"change24h"`
	ChangeBody     float64 `json:"changeBod"`
	Enabled        bool    `json:"enabled"`
	Restricted     bool    `json:"restricted"`
	PostOnly       bool    `json:"postOnly"`
}

func (c *Client) GetMarket(marketName string) (*Market, error) {
	var market Market
	err := c.get(marketsEndpoint+"/"+marketName, &market)
	if err != nil {
		return nil, err
	}

	return &market, nil
}

func (c *Client) GetMarkets() ([]Market, error) {
	var markets []Market
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
