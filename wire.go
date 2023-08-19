//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/huy-nguyenquoc/stably/api/handlers"
	"github.com/huy-nguyenquoc/stably/services"
)

var FeeServiceSet = wire.NewSet()

// InitializeService initializes the Service.
func InitializeFeeHandler() (handlers.FeeHandler, error) {
	wire.Build(
		services.NewFiatService,
		services.NewCryptoService,
		services.NewLiquidityFeeService,
		services.NewFeeService,
		services.NewFlowFeeService,
		handlers.NewFeeHandler,
		wire.Bind(new(services.FlowFeeService), new(*services.FlowFeeServiceImpl)),
		wire.Bind(new(services.FiatFeeService), new(*services.FiatFeeServiceImpl)),
		wire.Bind(new(services.CryptoFeeService), new(*services.CryptoFeeServiceImpl)),
		wire.Bind(new(services.LiquidityFeeService), new(*services.LiquidityFeeServiceImpl)),
		wire.Bind(new(services.FeeService), new(*services.FeeServiceImpl)),
		wire.Bind(new(handlers.FeeHandler), new(*handlers.FeeHandlerImpl)),
	)
	return nil, nil
}
