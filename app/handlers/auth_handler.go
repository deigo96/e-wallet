package handlers

import (
	"github.com/deigo96/e-wallet.git/app/controllers/auth"
	"github.com/deigo96/e-wallet.git/app/middleware"
	"github.com/deigo96/e-wallet.git/config"
	"github.com/gin-gonic/gin"
)

func NewAuthHandler(controller auth.Controller, router *gin.RouterGroup, config *config.Configuration) {
	authRoute := router.Group("/auth")

	authRoute.POST("/login", controller.Login)

	protectedAuthroute := authRoute
	protectedAuthroute.Use(middleware.Authorization(config))

	protectedAuthroute.POST("/send-otp", controller.SendOTP)
	protectedAuthroute.POST("/validate-otp", controller.ValidateOTP)
}
