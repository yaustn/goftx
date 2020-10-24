package goftx

import (
	"encoding/json"

	"github.com/yaustn/goftx/model"
)

const (
	ordersEndpoint = "/orders"
)

func (c *Client) GetOrders() (orders *[]model.Order, err error) {
	respBytes, err := c.get(ordersEndpoint)
	if err != nil {
		return orders, err
	}

	response := new(model.Response)
	response.Result = new([]model.Order)

	err = json.Unmarshal(respBytes, response)
	if err != nil {
		return orders, err
	}

	orders = response.Result.(*[]model.Order)

	return orders, nil
}

func (c *Client) GetOrdersByMarket(marketName string) (orders *[]model.Order, err error) {
	respBytes, err := c.get(ordersEndpoint +
		"/?market=" + marketName)
	if err != nil {
		return orders, err
	}

	response := new(model.Response)
	response.Result = new([]model.Order)

	err = json.Unmarshal(respBytes, response)
	if err != nil {
		return orders, err
	}

	orders = response.Result.(*[]model.Order)

	return orders, nil
}
