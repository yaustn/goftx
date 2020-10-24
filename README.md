# goftx
A Golang implementation of the ![FTX REST API specification](https://docs.ftx.com/#overview).

This package is currently still in alpha. It is the minimally necessary endpoints to get a basic market maker or algorithmic trading bot stood up.

## Usage

Import the package
```
go get -u github.com/yaustn/goftx
```

Example usage
```
go run example/main.go
```

## todo
- Add support for Subaccounts
- Make request creation more generic
- Add additional endpoints
