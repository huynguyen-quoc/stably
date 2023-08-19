package services

import (
	model "github.com/huy-nguyenquoc/stably/domains"
	"github.com/huy-nguyenquoc/stably/domains/amount"
	"testing"
)

func TestFlowFeeService(t *testing.T) {
	service := NewFlowFeeService()
	t.Run("testing with fiat->crypto network fee", func(t *testing.T) {
		tradeAmount := amount.NewCurrencyAmount("10", "USD")
		customer := &model.Customer{
			Tier: "1",
		}
		result, _ := service.Calculate(tradeAmount, "ACH", "ethereum", customer)
		expectedAmount := "0.05"
		if result.ToAmountString() != expectedAmount {
			t.Errorf("Expected %s, but got %s", expectedAmount, result.ToAmountString())
		}
	})
	t.Run("testing with crypto->crypto network fee", func(t *testing.T) {
		tradeAmount := amount.NewCurrencyAmount("10", "USD")
		customer := &model.Customer{
			Tier: "1",
		}
		result, _ := service.Calculate(tradeAmount, "bitcoin", "ethereum", customer)
		expectedAmount := "0.02"
		if result.ToAmountString() != expectedAmount {
			t.Errorf("Expected %s, but got %s", expectedAmount, result.ToAmountString())
		}
	})
	t.Run("testing with crypto->fiat network fee", func(t *testing.T) {
		tradeAmount := amount.NewCurrencyAmount("10", "USD")
		customer := &model.Customer{
			Tier: "1",
		}
		result, _ := service.Calculate(tradeAmount, "ethereum", "ACH", customer)
		expectedAmount := "0.03"
		if result.ToAmountString() != expectedAmount {
			t.Errorf("Expected %s, but got %s", expectedAmount, result.ToAmountString())
		}
	})

}
