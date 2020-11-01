package goftx

import (
	"github.com/yaustn/goftx/model"
)

const (
	walletRoute      = "/wallet"
	balancesEndpoint = "/balances"
	coinsEndpoint    = "/coins"
)

func (c *Client) GetBalances() (*[]model.Balance, error) {
	balances := new([]model.Balance)
	err := c.get(walletRoute+balancesEndpoint, balances)
	if err != nil {
		return nil, err
	}

	return balances, nil

}

func (c *Client) GetCoins() (*[]model.Coin, error) {
	coins := new([]model.Coin)
	err := c.get(walletRoute+coinsEndpoint, coins)
	if err != nil {
		return nil, err
	}

	return coins, nil

}
