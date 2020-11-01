package goftx

import (
	"github.com/yaustn/goftx/model"
)

const (
	ordersEndpoint = "/orders"
	marketParam    = "?market="
)

func (c *Client) GetOrders() (*[]model.Order, error) {
	orders := new([]model.Order)
	err := c.get(ordersEndpoint, orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (c *Client) GetOrdersByMarket(marketName string) (*[]model.Order, error) {
	orders := new([]model.Order)
	err := c.get(ordersEndpoint+marketParam+marketName, orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
