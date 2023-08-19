package services

import (
	"errors"
	model "github.com/huy-nguyenquoc/stably/domains"
	"github.com/huy-nguyenquoc/stably/domains/amount"
)

type CryptoFeeService interface {
	Calculate(tradeAmount *amount.CurrencyAmount, network string, customer *model.Customer) (*amount.CurrencyAmount, error)
}

type CryptoFeeServiceImpl struct {
	feeMap map[string]func(txnAmount *amount.CurrencyAmount, customer *model.Customer) *amount.CurrencyAmount
}

func NewCryptoService() *CryptoFeeServiceImpl {
	feeMap := map[string]func(txnAmount *amount.CurrencyAmount, customer *model.Customer) *amount.CurrencyAmount{
		"ethereum": func(txnAmount *amount.CurrencyAmount, customer *model.Customer) *amount.CurrencyAmount {
			discountPercent := tierDiscountCrypto[customer.Tier]("ethereum")
			fixedFee := amount.NewCurrencyAmount("10", "USD")
			discountFee := fixedFee.Percent(discountPercent)
			return fixedFee.Subtract(discountFee)
		},
		"bitcoin": func(txnAmount *amount.CurrencyAmount, customer *model.Customer) *amount.CurrencyAmount {
			discountPercent := tierDiscountFiat[customer.Tier]("bitcoin")
			fixedFee := amount.NewCurrencyAmount("15", "USD")
			discountFee := fixedFee.Percent(discountPercent)
			return fixedFee.Subtract(discountFee)
		},
	}
	return &CryptoFeeServiceImpl{feeMap}
}

func (t *CryptoFeeServiceImpl) Calculate(tradeAmount *amount.CurrencyAmount, network string, customer *model.Customer) (*amount.CurrencyAmount, error) {
	feeCalculator := t.feeMap[network]
	if feeCalculator == nil {
		return nil, errors.New("unsupported crypto network")
	}
	return feeCalculator(tradeAmount, customer), nil
}
