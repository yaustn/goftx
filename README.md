# goftx [![GoDoc](https://godoc.org/github.com/yaustn/goftx?status.svg)](https://pkg.go.dev/github.com/yaustn/goftx) [![Go Report Card](https://goreportcard.com/badge/github.com/yaustn/goftx)](https://goreportcard.com/report/github.com/yaustn/goftx) 
A lightweight Golang implementation of the ![FTX REST API specification](https://docs.ftx.com/#overview).

## Usage

Add the latest version to your go.mod
```
require github.com/yaustn/goftx v1.0.0
```

Example REST API Call:
```
import "github.com/yaustn/goftx"

func main() {
    client := goftx.NewClient(<API Key>, <API Secret>)

    orders, err := client.GetOrders()
	if err != nil {
        // Handle errors
	}

	// Process orders
}
```

## REST API

### Subaccounts
- [ ] GET /subaccounts
- [ ] POST /subaccounts
- [ ] POST /subaccounts/update_name
- [ ] DELETE /subaccounts
- [ ] GET /subaccounts/{nickname}/balances
- [ ] POST /subaccounts/transfer

### Markets
- [x] GET /markets
- [x] GET /market{market_name}
- [x] GET /markets/{market_name}/orderbook?depth={depth}
- [x] GET /markets/{market_name}/trades?limit={limit}&start_time={start_time}&end_time={end_time}
- [x] GET /markets/{market_name}/candles?resolution={resolution}&limit={limit}&start_time={start_time}&end_time={end_time}

### Wallet
- [x] GET /wallet/coins
- [x] GET /wallet/balances
- [ ] GET /wallet/all_balances 
- [ ] GET /wallet/deposit_address/{coin}?method={method}
- [ ] GET /wallet/deposits
- [ ] GET /wallet/withdrawals
- [ ] POST /wallet/withdrawals
- [ ] GET /wallet/saved_addresses
- [ ] POST /wallet/saved_addresses
- [ ] POST /wallet/saved_addresses

### Orders
- [x] GET /orders
- [x] GET /orders?market={market}
- [x] GET /orders/history?market={market}
- [ ] GET /conditional_orders?market={market}
- [ ] GET /conditional_orders/{conditional_order_id}/triggers
- [ ] GET /conditional_orders/history?market={market}
- [x] POST /orders
- [ ] POST /conditional_orders
- [ ] POST /orders/{order_id}/modify
- [ ] POST /orders/by_client_id/{client_order_id}/modify
- [ ] POST /conditional_orders/{order_id}/modify
- [ ] GET /orders/{order_id}
- [ ] GET /orders/by_client_id/{client_order_id}
- [x] DELETE /orders/{order_id}
- [ ] DELETE /orders/by_client_id/{client_order_id}
- [ ] DELETE /conditional_orders/{id}
- [x] DELETE /orders

## todo
Currently, the endpoints are only the minimally necessary REST endpoints to get a basic algorithmic market maker/trader stood up. There is plenty left to add, namely: 

- Add support for Subaccounts
- Add support for Websockets
