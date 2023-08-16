package main

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	userGroup := router.Group("v1/fee")
	{
		userHandler, _ := InitializeFeeHandler()
		userGroup.GET("", userHandler.Get)
	}
}
