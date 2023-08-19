package services

import (
	"github.com/huy-nguyenquoc/stably/domains/amount"
)

type LiquidityFeeService interface {
	Calculate(tradeAmount *amount.CurrencyAmount) (*amount.CurrencyAmount, string, error)
}
type LiquidityFeeServiceImpl struct {
	feeMap map[string]func(txnAmount *amount.CurrencyAmount) *amount.CurrencyAmount
}

func NewLiquidityFeeService() *LiquidityFeeServiceImpl {
	feeMap := map[string]func(txnAmount *amount.CurrencyAmount) *amount.CurrencyAmount{
		"Duck": func(txnAmount *amount.CurrencyAmount) *amount.CurrencyAmount {
			five := amount.NewCurrencyAmount("5", "USD")
			result, _ := five.Percent("0.1").Add(five)
			return result
		},
		"Goose": func(txnAmount *amount.CurrencyAmount) *amount.CurrencyAmount {
			return txnAmount.Percent("0.3")
		},
		"Fox": func(txnAmount *amount.CurrencyAmount) *amount.CurrencyAmount {
			return txnAmount.Percent("0.5")
		},
	}

	return &LiquidityFeeServiceImpl{
		feeMap,
	}
}

func (f *LiquidityFeeServiceImpl) Calculate(tradeAmount *amount.CurrencyAmount) (*amount.CurrencyAmount, string, error) {
	var lowestFee *amount.CurrencyAmount
	var lowestProvider string
	for name, provider := range f.feeMap {
		fee := provider(tradeAmount)
		if lowestFee == nil || fee.Cmp(lowestFee) < 0 {
			lowestFee = fee
			lowestProvider = name
		}
	}
	return lowestFee, lowestProvider, nil
}
