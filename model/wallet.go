package model

// Balance represents an FTX account's wallet balance.
// See https://docs.ftx.com/#get-balances
type Balance struct {
	Coin  string  `json:"coin"`
	Free  float64 `json:"free"`
	Total float64 `json:"total"`
}
