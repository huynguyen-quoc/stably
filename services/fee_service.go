package services

import (
	"errors"
	model "github.com/huy-nguyenquoc/stably/domains"
	"github.com/huy-nguyenquoc/stably/domains/amount"
)

type FeeService interface {
	Calculate(transaction *model.Transaction, customer *model.Customer) (*amount.CurrencyAmount, string, error)
}

type FeeServiceImpl struct {
	cryptoFeeService    CryptoFeeService
	fiatFeeService      FiatFeeService
	liquidityFeeService LiquidityFeeService
}

func NewFeeService(cryptoFeeService CryptoFeeService, fiatFeeService FiatFeeService, liquidityFeeService LiquidityFeeService) *FeeServiceImpl {
	return &FeeServiceImpl{cryptoFeeService, fiatFeeService, liquidityFeeService}
}

func (t *FeeServiceImpl) Calculate(transaction *model.Transaction, customer *model.Customer) (*amount.CurrencyAmount, string, error) {
	// Calculate the base fee for the transaction.
	tradeAmount := amount.NewCurrencyAmount(transaction.FromAmount, transaction.FromAsset)
	cryptoFeeResult, _ := t.cryptoFeeService.Calculate(tradeAmount, transaction.ToNetwork, customer)
	fiatFeeResult, _ := t.fiatFeeService.Calculate(tradeAmount, transaction.FromNetwork, customer)
	liquidityFeeResult, name, _ := t.liquidityFeeService.Calculate(tradeAmount.Subtract(cryptoFeeResult).Subtract(fiatFeeResult))
	resultSumLiquidityWithFiat, errWithSum := liquidityFeeResult.Add(fiatFeeResult)
	if errWithSum != nil {
		return nil, "", errors.New("invalid.amount")
	}
	resultSumTotal, errWithTotalSum := resultSumLiquidityWithFiat.Add(cryptoFeeResult)
	if errWithTotalSum != nil {
		return nil, "", errors.New("invalid.amount")
	}
	// Return the lowest fee and the name of the liquidity provider.
	return resultSumTotal, name, nil
}
