package users

import (
	"net/http"

	"github.com/deigo96/e-wallet.git/app/error"
	"github.com/deigo96/e-wallet.git/app/models"
	"github.com/deigo96/e-wallet.git/app/services/users"
	"github.com/deigo96/e-wallet.git/config"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Controller struct {
	userServices users.UserService
	config       config.Configuration
}

func NewUserController(db *gorm.DB, config config.Configuration) Controller {
	return Controller{
		userServices: users.NewUserService(&config, db),
		config:       config,
	}
}

var validate *validator.Validate

func (controller Controller) CreateUser(c *gin.Context) {
	user := &models.CreateUserRequest{}

	if err := c.BindJSON(user); err != nil {
		error.ErrorResponse(err, c)
		return
	}

	validate = validator.New()

	if err := validate.Struct(user); err != nil {
		error.ErrorResponse(err, c)
		return
	}

	if err := controller.userServices.CreateUser(c, user); err != nil {
		error.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
	})
}

func (controller Controller) GetUsersHandler(c *gin.Context) {
	users, err := controller.userServices.GetAllUsers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, users)
}
