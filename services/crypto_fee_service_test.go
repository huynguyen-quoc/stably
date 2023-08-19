package services

import (
	model "github.com/huy-nguyenquoc/stably/domains"
	"github.com/huy-nguyenquoc/stably/domains/amount"
	"testing"
)

func TestCryptoFeeService(t *testing.T) {
	service := NewCryptoService()
	t.Run("testing with ethereum network fee", func(t *testing.T) {
		tradeAmount := amount.NewCurrencyAmount("20", "USD")
		customer := &model.Customer{
			Tier: "1",
		}
		result, _ := service.Calculate(tradeAmount, "ethereum", customer)
		expectedAmount := "10.00"
		if result.ToAmountString() != expectedAmount {
			t.Errorf("Expected %s, but got %s", expectedAmount, result.ToAmountString())
		}
	})

}
