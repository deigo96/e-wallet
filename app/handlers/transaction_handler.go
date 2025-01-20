package handlers

import (
	"github.com/deigo96/e-wallet.git/app/controllers/transaction"
	"github.com/deigo96/e-wallet.git/app/middleware"
	"github.com/deigo96/e-wallet.git/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewTransactionHandler(controller transaction.Controller, router *gin.RouterGroup, config *config.Configuration, db *gorm.DB) {
	transactionRoute := router.Group("/transactions")

	transactionRoute.Use(middleware.Authorization(config))
	transactionRoute.Use(middleware.TransactionAuthorization(config, db))

	transactionRoute.POST("/topup", controller.Topup)

}
