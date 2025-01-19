package handlers

import (
	"github.com/deigo96/e-wallet.git/app/controllers/transaction"
	"github.com/deigo96/e-wallet.git/app/middleware"
	"github.com/deigo96/e-wallet.git/config"
	"github.com/gin-gonic/gin"
)

func NewTransactionHandler(controller transaction.Controller, router *gin.RouterGroup, config *config.Configuration) {
	transactionRoute := router.Group("/transactions")

	transactionRoute.Use(middleware.Authorization(config))

	transactionRoute.POST("/topup", controller.Topup)

}
