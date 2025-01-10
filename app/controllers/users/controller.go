package users

import (
	"net/http"

	"github.com/deigo96/e-wallet.git/app/services/users"
	"github.com/deigo96/e-wallet.git/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller struct {
	userServices users.UserService
	config       config.Configuration
}

func NewUserController(db *gorm.DB, config config.Configuration) Controller {
	return Controller{userServices: users.NewUserService(&config, db), config: config}
}

func (controller Controller) GetUsersHandler(c *gin.Context) {
	users, err := controller.userServices.GetAllUsers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}
