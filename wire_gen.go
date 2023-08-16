// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/huy-nguyenquoc/stably/api/handlers"
	"github.com/huy-nguyenquoc/stably/services"
)

// Injectors from wire.go:

// InitializeService initializes the Service.
func InitializeFeeHandler() (handlers.FeeHandler, error) {
	cryptoFeeServiceImpl := services.NewCryptoService()
	fiatFeeServiceImpl := services.NewFiatService()
	liquidityFeeServiceImpl := services.NewLiquidityFeeService()
	feeServiceImpl := services.NewFeeService(cryptoFeeServiceImpl, fiatFeeServiceImpl, liquidityFeeServiceImpl)
	feeHandlerImpl := handlers.NewFeeHandler(feeServiceImpl)
	return feeHandlerImpl, nil
}

// wire.go:

var FeeServiceSet = wire.NewSet()