package handlers

import (
	"github.com/deigo96/e-wallet.git/app/controllers/auth"
	"github.com/deigo96/e-wallet.git/app/controllers/transaction"
	"github.com/deigo96/e-wallet.git/app/controllers/users"
	"github.com/deigo96/e-wallet.git/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewHandler(config *config.Configuration, db *gorm.DB, router *gin.RouterGroup) {
	// repository := repository.NewRepository(db)
	// services := services.NewService(*repository, config)
	// controller := controllers.NewController(services, *config)

	userController := users.NewUserController(db, *config)
	authController := auth.NewAuthController(db, *config)
	transactionController := transaction.NewTransactionController(db, *config)

	NewUserHandler(userController, router, config)
	NewAuthHandler(authController, router, config)
	NewTransactionHandler(transactionController, router, config, db)
}
