package goftx

import (
	"encoding/json"

	"github.com/yaustn/goftx/model"
)

const (
	walletRoute      = "/wallet"
	balancesEndpoint = "/balances"
	coinsEndpoint    = "/coins"
)

func (c *Client) GetBalances() (balances *[]model.Balance, err error) {
	respBytes, err := c.get(walletRoute + balancesEndpoint)
	if err != nil {
		return balances, err
	}

	response := new(model.Response)
	response.Result = new([]model.Balance)

	err = json.Unmarshal(respBytes, response)
	if err != nil {
		return balances, err
	}

	balances = response.Result.(*[]model.Balance)

	return balances, nil
}

func (c *Client) GetCoins() (coins *[]model.Coin, err error) {
	respBytes, err := c.get(walletRoute + coinsEndpoint)
	if err != nil {
		return coins, err
	}

	response := new(model.Response)
	response.Result = new([]model.Coin)

	err = json.Unmarshal(respBytes, response)
	if err != nil {
		return coins, err
	}

	coins = response.Result.(*[]model.Coin)

	return coins, nil
}
