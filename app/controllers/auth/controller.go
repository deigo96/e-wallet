package auth

import (
	"net/http"

	"github.com/deigo96/e-wallet.git/app/error"
	customError "github.com/deigo96/e-wallet.git/app/error"
	"github.com/deigo96/e-wallet.git/app/models"
	"github.com/deigo96/e-wallet.git/app/services/auth"
	"github.com/deigo96/e-wallet.git/app/utils"
	"github.com/deigo96/e-wallet.git/config"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Controller struct {
	authServices auth.AuthService
	config       config.Configuration
}

func NewAuthController(db *gorm.DB, config config.Configuration) Controller {
	return Controller{
		authServices: auth.NewAuthService(&config, db),
		config:       config,
	}
}

var validate *validator.Validate

func (controller Controller) Login(c *gin.Context) {
	request := &models.LoginRequest{}

	if err := c.BindJSON(request); err != nil {
		error.ErrorResponse(err, c)
		return
	}

	validate = validator.New()

	if err := validate.Struct(request); err != nil {
		error.ErrorResponse(err, c)
		return
	}

	response, err := controller.authServices.Login(c, request)
	if err != nil {
		error.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (controller Controller) SendOTP(c *gin.Context) {
	req := &models.SendOTPRequest{}

	if err := c.BindJSON(req); err != nil {
		error.ErrorResponse(err, c)
		return
	}

	validate = validator.New()

	if err := validate.Struct(req); err != nil {
		error.ErrorResponse(err, c)
		return
	}

	if !utils.ValidPhone(req.Phone) {
		error.ErrorResponse(customError.ErrInvalidPhone, c)
		return
	}

	_, err := controller.authServices.SendOTP(c, utils.RefactorPhoneNumber(req.Phone))
	if err != nil {
		error.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OTP sent successfully",
	})
}

func (controller Controller) ValidateOTP(c *gin.Context) {
	req := &models.ValidateOTPRequest{}

	if err := c.BindJSON(req); err != nil {
		error.ErrorResponse(err, c)
		return
	}

	validate = validator.New()

	if err := validate.Struct(req); err != nil {
		error.ErrorResponse(err, c)
		return
	}

	if err := controller.authServices.ValidateOTP(c, req); err != nil {
		error.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OTP validated successfully",
	})
}
