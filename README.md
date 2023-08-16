# Stably

Stably is a REST service to get the fee depend on the amount network from Fiat to crypto network

## Features

- Calculate fee to convert from asset to asset



### Layout

```tree
├── api
│   ├── handlers
│   │   └── fee.go
│   ├── domains
|   |   └── amount
|   |         └── amount.go
|   |         └── currenct_amount.go
│   │   └── customer_tier.go
│   │   └── transaction.go
│   ├── services
│   │   └── crypto_fee_service.go
│   │   └── customer_tier_discount.go
|   |   └── fee_service.go
|   |   └── fiat_fee_service.go
|   |   └── liquidity_provider_fee_service.go
│   ├── main.go
│   ├── routes.go
│   ├── wire.go
```

A brief description of the layout:

* `domains` place for models
* `pkg` places for service business logic
* `handlers` places for all api handlers


## How to Run
1. Install go version 1.18
2. Run `go mod download` to download the dependencies
3. Run `go run github.com/google/wire/cmd/wire` to create DI
4. Run `go run . ` to start server

## Example CURL
```
curl --request GET \
  --url 'http://localhost:8080/v1/fee?fromAmount=100&fromNetwork=ACH&fromAsset=USD&toNetwork=ethereum&toAsset=ETH&feeAsset=USD' \
  --header 'X-Customer-Tier: 1'
