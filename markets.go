package goftx

import (
	"log"
	"strconv"
)

const (
	marketsEndpoint   = "/markets"
	orderbookEndpoint = "/orderbook"
	tradesEndpoint    = "/trades"
	candlesEndpoint   = "/candles"
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
	ChangeBod      float64 `json:"changeBod"`
	Enabled        bool    `json:"enabled"`
	Restricted     bool    `json:"restricted"`
	PostOnly       bool    `json:"postOnly"`
}

type Orderbook struct {
	Asks [][]float64 `json:"asks"`
	Bids [][]float64 `json:"bids"`
}

type Trade struct {
	ID    int64   `json:"id"`
	Side  string  `json:"side"`
	Price float64 `json:"price"`
	Size  float64 `json:"size"`
	Time  string  `json:"time"`
}

type Candle struct {
	StartTime string  `json:"startTime"`
	Open      float64 `json:"open"`
	Close     float64 `json:"close"`
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Volume    float64 `json:"volume"`
}

type GetOrderbookRequest struct {
	Market string
	Depth  int // default 20, max 100
}

type GetTradesRequest struct {
	Market    string
	Limit     int   // default 20, max 100
	StartTime int64 // epoch seconds
	EndTime   int64 // epoch seconds
}

type GetCandlesRequest struct {
	Market     string
	Resolution int   // seconds; 15, 60, 300, 900, 3600, 14400, 86400
	Limit      int   // default 20, max 100
	StartTime  int64 // epoch seconds
	EndTime    int64 // epoch second
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

func (c *Client) GetOrderbook(request GetOrderbookRequest) (*Orderbook, error) {
	var orderbook Orderbook
	getOrderbookEndpoint := marketsEndpoint + "/" + request.Market + orderbookEndpoint
	if request.Depth != 0 {
		getOrderbookEndpoint += "?depth=" + strconv.Itoa(request.Depth)
	}

	log.Printf("%s", getOrderbookEndpoint)

	err := c.get(getOrderbookEndpoint, &orderbook)
	if err != nil {
		return nil, err
	}

	return &orderbook, nil
}

func (c *Client) GetTrades(request GetTradesRequest) ([]Trade, error) {
	var trades []Trade
	getTradesEndpoint := marketsEndpoint + "/" + request.Market + tradesEndpoint
	if request.Limit != 0 || request.StartTime != 0 || request.EndTime != 0 {
		getTradesEndpoint += "?"
	}

	if request.Limit != 0 {
		getTradesEndpoint += "&limit=" + strconv.Itoa(request.Limit)
	}

	if request.StartTime != 0 {
		getTradesEndpoint += "&start_time=" + strconv.FormatInt(request.StartTime, 10)
	}

	if request.EndTime != 0 {
		getTradesEndpoint += "&end_time=" + strconv.FormatInt(request.EndTime, 10)
	}

	err := c.get(getTradesEndpoint, &trades)
	if err != nil {
		return nil, err
	}

	return trades, nil
}

func (c *Client) GetCandles(request GetCandlesRequest) ([]Candle, error) {
	var candles []Candle
	getCandlesEndpoint := marketsEndpoint + "/" + request.Market + candlesEndpoint
	if request.Resolution != 0 || request.Limit != 0 || request.StartTime != 0 || request.EndTime != 0 {
		getCandlesEndpoint += "?"
	}

	if request.Resolution != 0 {
		getCandlesEndpoint += "&resolution=" + strconv.Itoa(request.Resolution)
	}

	if request.Limit != 0 {
		getCandlesEndpoint += "&limit=" + strconv.Itoa(request.Limit)
	}

	if request.StartTime != 0 {
		getCandlesEndpoint += "&start_time=" + strconv.FormatInt(request.StartTime, 10)
	}

	if request.EndTime != 0 {
		getCandlesEndpoint += "&end_time=" + strconv.FormatInt(request.EndTime, 10)
	}

	err := c.get(getCandlesEndpoint, &candles)
	if err != nil {
		return nil, err
	}

	return candles, nil
}
