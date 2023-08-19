package services

import (
	"errors"
	model "github.com/huy-nguyenquoc/stably/domains"
	"github.com/huy-nguyenquoc/stably/domains/amount"
)

type FlowFeeService interface {
	Calculate(tradeAmount *amount.CurrencyAmount, fromNetwork string, toNetwork string, customer *model.Customer) (*amount.CurrencyAmount, error)
}

type FlowFeeServiceImpl struct {
	fiat   map[string]bool
	crypto map[string]bool
}

func NewFlowFeeService() *FlowFeeServiceImpl {
	fiat := map[string]bool{
		"ACH":  true,
		"Wire": true,
		"Card": true,
	}
	crypto := map[string]bool{
		"ethereum": true,
		"bitcoin":  true,
	}
	return &FlowFeeServiceImpl{fiat, crypto}
}

func (t *FlowFeeServiceImpl) Calculate(tradeAmount *amount.CurrencyAmount, fromNetwork string, toNetwork string, customer *model.Customer) (*amount.CurrencyAmount, error) {
	fromType, errFromType := t.getNetworkType(fromNetwork)
	if errFromType != nil {
		return nil, errors.New("invalid.network.from")
	}
	toType, errToType := t.getNetworkType(toNetwork)
	if errToType != nil {
		return nil, errors.New("invalid.network.to")
	}
	if fromType == "fiat" && toType == "crypto" {
		return tradeAmount.Percent("0.5"), nil
	}

	if fromType == "crypto" && toType == "crypto" {
		return tradeAmount.Percent("0.2"), nil
	}
	if fromType == "crypto" && toType == "fiat" {
		return tradeAmount.Percent("0.3"), nil
	}
	return tradeAmount, nil
}

func (t *FlowFeeServiceImpl) getNetworkType(network string) (string, error) {
	if t.fiat[network] {
		return "fiat", nil
	}
	if t.crypto[network] {
		return "crypto", nil
	}

	return "", errors.New("network not supported")
}
