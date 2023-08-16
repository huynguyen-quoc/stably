package services

import (
	"errors"
	model "github.com/huy-nguyenquoc/stably/domains"
	"github.com/huy-nguyenquoc/stably/domains/amount"
)

type FiatFeeService interface {
	Calculate(tradeAmount *amount.CurrencyAmount, fiatNetwork string, customer *model.Customer) (*amount.CurrencyAmount, error)
}
type FiatFeeServiceImpl struct {
	feeMap map[string]func(txnAmount *amount.CurrencyAmount, customer *model.Customer) *amount.CurrencyAmount
}

func NewFiatService() *FiatFeeServiceImpl {
	feeMap := map[string]func(txnAmount *amount.CurrencyAmount, customer *model.Customer) *amount.CurrencyAmount{
		"ACH": func(txnAmount *amount.CurrencyAmount, customer *model.Customer) *amount.CurrencyAmount {
			feeFixed := amount.NewCurrencyAmount("2", "USD")
			feePercent := txnAmount.Percent("1")
			discountPercent := tierDiscountFiat[customer.Tier]("ACH")

			if feeFixed.Cmp(feePercent) < 0 {
				return feePercent.Subtract(feePercent.Percent(discountPercent))
			} else {
				return feeFixed.Subtract(feeFixed.Percent(discountPercent))
			}
		},
		"Wire": func(txnAmount *amount.CurrencyAmount, customer *model.Customer) *amount.CurrencyAmount {
			discountPercent := tierDiscountFiat[customer.Tier]("Wire")
			fixedFee := amount.NewCurrencyAmount("25", "USD")
			discountFee := fixedFee.Percent(discountPercent)
			return fixedFee.Subtract(discountFee)
		},
		"Card": func(txnAmount *amount.CurrencyAmount, customer *model.Customer) *amount.CurrencyAmount {
			discountPercent := tierDiscountFiat[customer.Tier]("Card")
			fixedFee := txnAmount.Percent("3")
			discountFee := fixedFee.Percent(discountPercent)
			return fixedFee.Subtract(discountFee)
		},
		// Add more currencies and their scale factors as needed
	}

	return &FiatFeeServiceImpl{
		feeMap,
	}
}

func (f *FiatFeeServiceImpl) Calculate(tradeAmount *amount.CurrencyAmount, fiatNetwork string, customer *model.Customer) (*amount.CurrencyAmount, error) {
	feeCalculator := f.feeMap[fiatNetwork]
	if feeCalculator == nil {
		return nil, errors.New("unsupported fiat network")
	}
	return feeCalculator(tradeAmount, customer), nil
}
