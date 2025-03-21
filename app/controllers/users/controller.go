package users

import (
	"log"
	"net/http"

	"github.com/deigo96/e-wallet.git/app/error"
	"github.com/deigo96/e-wallet.git/app/models"
	"github.com/deigo96/e-wallet.git/app/services/profile"
	"github.com/deigo96/e-wallet.git/app/services/users"
	"github.com/deigo96/e-wallet.git/config"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Controller struct {
	userServices    users.UserService
	profileServices profile.ProfileService
	config          config.Configuration
}

func NewUserController(db *gorm.DB, config config.Configuration) *Controller {
	return &Controller{
		userServices:    users.NewUserService(&config, db),
		profileServices: profile.NewProfileService(&config, db),
		config:          config,
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
		error.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusOK, users)
}

func (controller Controller) GetProfile(c *gin.Context) {
	profile, err := controller.profileServices.GetProfile(c)
	if err != nil {
		error.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusOK, profile)
}

func (controller Controller) CreateProfile(c *gin.Context) {
	req := &models.ProfileRequest{}

	if err := c.BindJSON(req); err != nil {
		error.ErrorResponse(err, c)
		return
	}

	validate = validator.New()

	if err := validate.Struct(req); err != nil {
		error.ErrorResponse(err, c)
		return
	}

	profile, err := controller.profileServices.CreateProfile(c, req)
	if err != nil {
		error.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusCreated, profile)
}

func (controller Controller) UpdateProfile(c *gin.Context) {
	req := &models.ProfileRequest{}

	if err := c.BindJSON(req); err != nil {
		error.ErrorResponse(err, c)
		return
	}

	validate = validator.New()

	if err := validate.Struct(req); err != nil {
		error.ErrorResponse(err, c)
		return
	}

	profile, err := controller.profileServices.UpdateProfile(c, req)
	if err != nil {
		error.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusOK, profile)
}

func (controller Controller) VerifyEmail(c *gin.Context) {
	email := c.Param("email")
	token := c.Param("token")

	if err := controller.userServices.VerifyEmail(c, email, token); err != nil {
		log.Println("Error verifying email: " + err.Error())
		error.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Email verified successfully",
	})
}

func (controller Controller) ResendEmailVerification(c *gin.Context) {
	if err := controller.userServices.ResendEmailVerification(c); err != nil {
		error.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Email verification link sent successfully",
	})
}
