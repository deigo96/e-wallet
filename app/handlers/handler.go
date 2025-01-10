package handlers

import (
	"github.com/deigo96/e-wallet.git/app/controllers/users"
	"github.com/deigo96/e-wallet.git/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewHandler(config *config.Configuration, db *gorm.DB, router *gin.RouterGroup) {
	// repository := repository.NewRepository(db)
	// services := services.NewService(*repository, config)
	// controller := controllers.NewController(services, *config)

	controller := users.NewUserController(db, *config)

	NewUserHandler(controller, router)
}
