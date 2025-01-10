package handlers

import (
	"github.com/deigo96/e-wallet.git/app/controllers/auth"
	"github.com/gin-gonic/gin"
)

func NewAuthHandler(controller auth.Controller, router *gin.RouterGroup) {
	authRoute := router.Group("/auth")

	authRoute.POST("/login", controller.Login)
}
