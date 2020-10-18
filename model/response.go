package model

// Response represents every FTX API response.
// Response.Result is an interface specific to the request and should be
// type-casted before use.
type Response struct {
	Success bool        `json:"success"`
	Error   string      `json:"error,omitempty"`
	Result  interface{} `json:"result,omitempty"`
}
