package goftx

import (
	"encoding/json"
	"fmt"

	"github.com/yaustn/goftx/model"
)

const (
	ordersEndpoint        = "/orders"
	ordersHistoryEndpoint = "/orders/history"
	marketParam           = "?market="
)

func (c *Client) GetOrders() ([]model.Order, error) {
	var orders []model.Order
	err := c.get(ordersEndpoint, &orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (c *Client) GetOrdersByMarket(marketName string) ([]model.Order, error) {
	var orders []model.Order
	err := c.get(ordersEndpoint+marketParam+marketName, &orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (c *Client) GetOrderHistoryByMarket(marketName string) ([]model.Order, error) {
	var orders []model.Order
	err := c.get(ordersHistoryEndpoint+marketParam+marketName, &orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (c *Client) PlaceOrder(market, side, _type string, price, size float64) (*model.Order, error) {
	request, _ := json.Marshal(model.PlaceOrderRequest{
		Market: market,
		Side:   side,
		Type:   _type,
		Price:  price,
		Size:   size,
	})

	var order model.Order
	err := c.post(ordersEndpoint, request, &order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (c *Client) CancelOrder(orderID int64) (bool, error) {
	request, _ := json.Marshal(model.CancelOrderRequest{OrderID: orderID})
	var result string
	err := c.delete(ordersEndpoint, request, &result)
	if err != nil {
		return false, err
	}

	fmt.Println(result)

	return false, nil
}
