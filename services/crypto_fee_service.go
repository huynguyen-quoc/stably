package services

import (
	model "github.com/huy-nguyenquoc/stably/domains"
	"github.com/huy-nguyenquoc/stably/domains/amount"
)

type CryptoFeeService interface {
	Calculate(tradeAmount *amount.CurrencyAmount, network string, customer *model.Customer) (*amount.CurrencyAmount, error)
}

type CryptoFeeServiceImpl struct {
}

func NewCryptoService() *CryptoFeeServiceImpl {
	return &CryptoFeeServiceImpl{}
}

func (t *CryptoFeeServiceImpl) Calculate(tradeAmount *amount.CurrencyAmount, network string, customer *model.Customer) (*amount.CurrencyAmount, error) {
	discountPercent := tierDiscountCrypto[customer.Tier](network)
	fixedFee := amount.NewCurrencyAmount("10", "USD")
	discountFee := fixedFee.Percent(discountPercent)
	return fixedFee.Subtract(discountFee), nil
}
