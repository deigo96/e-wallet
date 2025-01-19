package transaction

import (
	"net/http"

	"github.com/deigo96/e-wallet.git/app/constant"
	"github.com/deigo96/e-wallet.git/app/error"
	"github.com/deigo96/e-wallet.git/app/models"
	"github.com/deigo96/e-wallet.git/app/services/transactions"
	"github.com/deigo96/e-wallet.git/config"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Controller struct {
	transactionService transactions.TransactionService
	config             config.Configuration
}

func NewTransactionController(db *gorm.DB, config config.Configuration) Controller {
	return Controller{
		transactionService: transactions.NewTransactionService(&config, db),
		config:             config,
	}
}

func (controller *Controller) Topup(c *gin.Context) {
	var req models.TransactionRequest

	if err := c.BindJSON(&req); err != nil {
		error.ErrorResponse(err, c)
		return
	}

	validate := validator.New()

	if err := validate.Struct(req); err != nil {
		error.ErrorResponse(err, c)
		return
	}

	req.TransactionType = constant.TransactionTopup

	if err := controller.transactionService.Transaction(c, &req); err != nil {
		error.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Transaction successful",
	})
}
