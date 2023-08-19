package services

import (
	"errors"
	"fmt"
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
	flowFeeService      FlowFeeService
}

func NewFeeService(cryptoFeeService CryptoFeeService, fiatFeeService FiatFeeService, liquidityFeeService LiquidityFeeService, flowFeeService FlowFeeService) *FeeServiceImpl {
	return &FeeServiceImpl{cryptoFeeService, fiatFeeService, liquidityFeeService, flowFeeService}
}

func (t *FeeServiceImpl) Calculate(transaction *model.Transaction, customer *model.Customer) (*amount.CurrencyAmount, string, error) {
	// Calculate the base fee for the transaction.
	tradeAmount := amount.NewCurrencyAmount(transaction.FromAmount, transaction.FromAsset)
	cryptoFeeResult, _ := t.cryptoFeeService.Calculate(tradeAmount, transaction.ToNetwork, customer)
	fmt.Printf("calculated crypto currency %s", cryptoFeeResult.ToAmountString())
	fiatFeeResult, _ := t.fiatFeeService.Calculate(tradeAmount, transaction.FromNetwork, customer)
	resultWithFiatAndCryptoFee := tradeAmount.Subtract(cryptoFeeResult).Subtract(fiatFeeResult)
	liquidityFeeResult, name, _ := t.liquidityFeeService.Calculate(resultWithFiatAndCryptoFee)
	resultSumLiquidityWithFiat, errWithSum := liquidityFeeResult.Add(fiatFeeResult)
	fmt.Printf("calculated crypto currency 2 %s", resultSumLiquidityWithFiat.ToAmountString())
	if errWithSum != nil {
		return nil, "", errors.New("invalid.amount")
	}
	resultSumWithLiquidity, errWithTotalSum := resultSumLiquidityWithFiat.Add(cryptoFeeResult)
	if errWithTotalSum != nil {
		return nil, "", errors.New("invalid.amount")
	}
	//resultSumWithFlow, errWithFlow := t.flowFeeService.Calculate(tradeAmount, transaction.FromNetwork, transaction.ToNetwork, customer)
	//if errWithFlow != nil {
	//	return nil, "", errors.New("invalid.amount")
	//}
	totalSum := resultSumWithLiquidity
	//fmt.Printf("resultSum %s", resultSumWithLiquidity.)
	//if errWithTotal != nil {
	//	return nil, "", errors.New("invalid.amount")
	//}
	// Return the lowest fee and the name of the liquidity provider.
	return totalSum, name, nil
}
