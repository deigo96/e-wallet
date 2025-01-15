package handlers

import (
	"github.com/deigo96/e-wallet.git/app/controllers/users"
	"github.com/deigo96/e-wallet.git/app/middleware"
	"github.com/deigo96/e-wallet.git/config"
	"github.com/gin-gonic/gin"
)

func NewUserHandler(controller users.Controller, router *gin.RouterGroup, config *config.Configuration) {
	userRoute := router.Group("/users")
	userRoute.Use(middleware.Authorization(config))

	userRoute.GET("", controller.GetUsersHandler)
	userRoute.POST("/register", controller.CreateUser)
	userRoute.GET("/profile", controller.GetProfile)
	userRoute.POST("/profile", controller.CreateProfile)
}
