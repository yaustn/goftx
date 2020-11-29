package goftx

import (
	"encoding/json"
	"strconv"
)

const (
	ordersEndpoint        = "/orders"
	ordersHistoryEndpoint = "/orders/history"
	marketParam           = "?market="
)

type Order struct {
	ID            int64   `json:"id"`
	CreatedAt     string  `json:"createdAt"`
	Market        string  `json:"market"`
	Side          string  `json:"side"`
	Type          string  `json:"type"`
	Price         float64 `json:"price"`
	Size          float64 `json:"size"`
	FilledSize    float64 `json:"filledSize"`
	RemainingSize float64 `json:"remainingSize"`
	AvgFillPrice  float64 `json:"avgFillPrice"`
	Status        string  `json:"status"` // new, open, or closed
	ReduceOnly    bool    `json:"reduceOnly"`
	IOC           bool    `json:"ioc"`
	PostOnly      bool    `json:"postOnly"`
	ClientID      string  `json:"clientId"`
}

// PlaceOrderRequest to place a new order
//
// ReduceOnly will only close out current positions
// IOC (immediate-or-cancel) orders will only take
// PostOnly orders will only make
type PlaceOrderRequest struct {
	Market     string  `json:"market"`
	Side       string  `json:"side"`  // "buy" or "sell"
	Type       string  `json:"type"`  // "limit" or "market"
	Price      float64 `json:"price"` // nil for market orders
	Size       float64 `json:"size"`
	ReduceOnly bool    `json:"reduceOnly,omitempty"` // optional - default is false
	IOC        bool    `json:"ioc,omitempty"`        // optional - default is false
	PostOnly   bool    `json:"postOnly,omitempty"`   // optional - default is false
	ClientID   string  `json:"clientId,omitempty"`   // optional
}

type CancelOrdersRequest struct {
	Market          string `json:"market,omitempty"`
	ConditionalOnly string `json:"conditionalOrdersOnly,omitempty"`
	LimitOnly       string `json:"limitOrdersOnly,omitempty"`
}

func (c *Client) GetOrders() ([]Order, error) {
	var orders []Order
	err := c.get(ordersEndpoint, &orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (c *Client) GetOrdersByMarket(marketName string) ([]Order, error) {
	var orders []Order
	err := c.get(ordersEndpoint+marketParam+marketName, &orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (c *Client) GetOrderHistoryByMarket(marketName string) ([]Order, error) {
	var orders []Order
	err := c.get(ordersHistoryEndpoint+marketParam+marketName, &orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (c *Client) PlaceOrder(request PlaceOrderRequest) (*Order, error) {
	jsonRequest, _ := json.Marshal(request)

	var order Order
	err := c.post(ordersEndpoint, jsonRequest, &order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (c *Client) CancelOrders(request CancelOrdersRequest) (bool, error) {
	jsonRequest, _ := json.Marshal(request)
	var result string
	err := c.delete(ordersEndpoint, jsonRequest, &result)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (c *Client) CancelOrder(orderID int64) (bool, error) {
	var result string
	err := c.delete(ordersEndpoint+"/"+strconv.FormatInt(orderID, 10), []byte{}, &result)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (c *Client) CancelAllOrders() (bool, error) {
	var result string
	err := c.delete(ordersEndpoint, []byte{}, &result)
	if err != nil {
		return false, err
	}

	return true, nil
}
