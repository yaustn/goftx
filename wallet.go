package goftx

import (
	"encoding/json"
	"log"

	"github.com/yaustn/goftx/model"
)

const (
	walletRoute      = "/wallet"
	balancesEndpoint = "/balances"
)

func (c *Client) GetBalances() (balances *[]model.Balance, err error) {
	respBytes, err := c.get(walletRoute + balancesEndpoint)
	if err != nil {
		log.Printf("[ERROR] Failed GET %s%s request: %+v", walletRoute, balancesEndpoint, err)
		return balances, err
	}

	log.Printf("GetbalanceResponse: %s", string(respBytes))

	response := new(model.Response)
	response.Result = new([]model.Balance)

	err = json.Unmarshal(respBytes, response)
	if err != nil {
		log.Printf("[ERROR] Failed to unmarshal GET %s%s response: %+v", walletRoute, balancesEndpoint, err)
		return balances, err
	}

	balances = response.Result.(*[]model.Balance)

	return balances, nil
}
