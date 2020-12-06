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

## todo
Currently, the endpoints are only the minimally necessary REST endpoints to get a basic algorithmic market maker/trader stood up. There is plenty left to add, namely: 

- Add support for Subaccounts
- Add support for websockets
