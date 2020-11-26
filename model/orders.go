package model

type Order struct {
	ID            int64   `json:"id"`
	CreatedAt     string  `json:"createdAt"`
	Market        string  `json:"market"`
	Side          string  `json:"side"`
	Type          string  `json:"type"`
	Price         float64 `json:"price"`
	Size          float64 `json:"size"`
	FilledSize    float64 `json:"filledSize"`
	RemainingSize float64 `json:"remainingSize"`
	AvgFillPrice  float64 `json:"avgFillPrice"`
	Status        string  `json:"status"` // new, open, or closed
	ReduceOnly    bool    `json:"reduceOnly"`
	IOC           bool    `json:"ioc"`
	PostOnly      bool    `json:"postOnly"`
	ClientID      string  `json:"clientId"`
}

// PlaceOrderRequest model to place a new order
//
// ReduceOnly will only close out current positions
// IOC (immediate-or-cancel) orders will only take
// PostOnly orders will only make
type PlaceOrderRequest struct {
	Market     string  `json:"market"`
	Side       string  `json:"side"`  // "buy" or "sell"
	Type       string  `json:"type"`  // "limit" or "market"
	Price      float64 `json:"price"` // nil for market orders
	Size       float64 `json:"size"`
	ReduceOnly bool    `json:"reduceOnly,omitempty"` // optional - default is false
	IOC        bool    `json:"ioc,omitempty"`        // optional - default is false
	PostOnly   bool    `json:"postOnly,omitempty"`   // optional - default is false
	ClientID   string  `json:"clientId,omitempty"`   // optional
}

type CancelOrderRequest struct {
	OrderID int64 `json:"id"`
}
