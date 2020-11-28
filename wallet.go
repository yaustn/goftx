package goftx

const (
	walletRoute      = "/wallet"
	balancesEndpoint = "/balances"
	coinsEndpoint    = "/coins"
)

type Balance struct {
	Coin  string  `json:"coin"`
	Free  float64 `json:"free"`
	Total float64 `json:"total"`
}

type Coin struct {
	ID               string   `json:"id"`
	Name             string   `json:"name"`
	CanDeposit       bool     `json:"canDeposit"`
	CanWithdraw      bool     `json:"canWithdraw"`
	CanConvert       bool     `json:"canConvert"`
	HasTag           bool     `json:"hasTag"`
	Collateral       bool     `json:"collateral"`
	CollateralWeight float64  `json:"collateralWeight"`
	CreditTo         string   `json:"creditTo"`
	Methods          []string `json:"methods"`
	SpotMargin       bool     `json:"spotMargin"`
	USDFungible      bool     `json:"usdFungible"`
	Bep2Asset        string   `json:"bep2Asset"`
	Fiat             bool     `json:"fiat"`
	IsToken          bool     `json:"isToken"`
	ERC20Contract    string   `json:"erc20Contract"`
	TRC20Contract    string   `json:"trc20Contract"`
}

func (c *Client) GetBalances() ([]Balance, error) {
	var balances []Balance
	err := c.get(walletRoute+balancesEndpoint, &balances)
	if err != nil {
		return nil, err
	}

	return balances, nil
}

func (c *Client) GetCoins() ([]Coin, error) {
	var coins []Coin
	err := c.get(walletRoute+coinsEndpoint, &coins)
	if err != nil {
		return nil, err
	}

	return coins, nil
}
