package handlers

import (
	"github.com/gin-gonic/gin"
	model "github.com/huy-nguyenquoc/stably/domains"
	"github.com/huy-nguyenquoc/stably/services"
)

type FeeHandler interface {
	Get(c *gin.Context)
}

type FeeHandlerImpl struct {
	feeService services.FeeService
}

func NewFeeHandler(feeService services.FeeService) *FeeHandlerImpl {
	return &FeeHandlerImpl{feeService}
}

func (f *FeeHandlerImpl) Get(c *gin.Context) {
	var transferQuery *model.Transaction

	// Bind query parameters to the TransferQuery struct
	if err := c.ShouldBindQuery(&transferQuery); err != nil {
		c.JSON(400, gin.H{"error": "Invalid query parameters"})
		return
	}
	customerTier := c.Request.Header.Get("X-Customer-Tier")
	if customerTier == "" {
		customerTier = "1"
	}
	tier := &model.Customer{Tier: customerTier}
	amount, name, _ := f.feeService.Calculate(transferQuery, tier)
	// Process the transferQuery object or perform necessary operations
	c.JSON(200, gin.H{
		"fee":      amount.ToAmountString(),
		"asset":    amount.Currency,
		"provider": name,
	})
}
