package model

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
