package model

type Transaction struct {
	FromAmount  string `form:"fromAmount" binding:"required"`
	FromNetwork string `form:"fromNetwork" binding:"required"`
	FromAsset   string `form:"fromAsset" binding:"required"`
	ToNetwork   string `form:"toNetwork" binding:"required"`
	ToAsset     string `form:"toAsset" binding:"required"`
	FeeAsset    string `form:"feeAsset" binding:"required"`
}
