package handlers

import (
	"github.com/deigo96/e-wallet.git/app/controllers/users"
	"github.com/gin-gonic/gin"
)

func NewUserHandler(controller users.Controller, router *gin.RouterGroup) {
	router.GET("/users", controller.GetUsersHandler)
}
