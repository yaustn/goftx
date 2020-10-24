package model

type Order struct {
	ID            int64   `json:"id"`
	Market        string  `json:"market"`
	Type          string  `json:"type"`
	Side          string  `json:"side"`
	Price         float64 `json:"price"`
	Size          float64 `json:"size"`
	FilledSize    float64 `json:"filledSize"`
	RemainingSize float64 `json:"remainingSize"`
	AvgFillPrice  float64 `json:"avgFillPrice"`
	Status        string  `json:"status"` // new, open, or closed
	CreatedAt     string  `json:"createdAt"`
	ReduceOnly    bool    `json:"reduceOnly"`
	IOC           bool    `json:"ioc"`
	PostOnly      bool    `json:"postOnly"`
	ClientID      string  `json:"clientId"`
}
