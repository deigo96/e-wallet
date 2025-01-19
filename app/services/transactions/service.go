package transactions

import (
	"log"
	"time"

	"github.com/deigo96/e-wallet.git/app/constant"
	"github.com/deigo96/e-wallet.git/app/entity"
	customError "github.com/deigo96/e-wallet.git/app/error"
	"github.com/deigo96/e-wallet.git/app/external"
	"github.com/deigo96/e-wallet.git/app/models"
	"github.com/deigo96/e-wallet.git/app/repository/balances"
	"github.com/deigo96/e-wallet.git/app/repository/orders"
	"github.com/deigo96/e-wallet.git/app/repository/profile"
	"github.com/deigo96/e-wallet.git/app/repository/transactions"
	"github.com/deigo96/e-wallet.git/app/repository/users"
	"github.com/deigo96/e-wallet.git/app/utils"
	"github.com/deigo96/e-wallet.git/config"
	"github.com/gin-gonic/gin"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type TransactionService interface {
	Topup(c *gin.Context, req *models.TransactionRequest) error
	Transaction(c *gin.Context, req *models.TransactionRequest) error
}

type transactionsService struct {
	balanceRepository     balances.BalanceRepository
	userRepository        users.UserRepository
	transactionRepository transactions.TransactionRepository
	profileRepository     profile.ProfileRepository
	orderRepository       orders.OrderRepository
	midtrans              *external.Midtrans
	config                *config.Configuration
	db                    *gorm.DB
}

func NewTransactionService(config *config.Configuration, db *gorm.DB) TransactionService {
	return &transactionsService{
		balanceRepository:     balances.NewBalanceRepository(db),
		userRepository:        users.NewUserRepository(db),
		transactionRepository: transactions.NewTransactionRepository(db),
		orderRepository:       orders.NewOrderRepository(db),
		midtrans:              external.NewMidtrans(config),
		profileRepository:     profile.NewProfileRepository(db),
		config:                config,
		db:                    db,
	}
}

func (t *transactionsService) constructTopup(req *models.TransactionRequest) *entity.Transaction {
	return &entity.Transaction{
		UserID:          req.UserID,
		TotalAmount:     decimal.NewFromFloat(req.Amount),
		TransactionType: req.TransactionType,
		OrderID:         utils.GenerateOrderID(),
		Note:            req.Note,
		Status:          constant.TransactionPending,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
}

func (t *transactionsService) constructTopupOrder(orderID string, req *models.TransactionRequest) *entity.Order {
	return &entity.Order{
		OrderID:   orderID,
		Item:      req.TransactionType.GetTransactionType(),
		Qty:       1,
		Amount:    decimal.NewFromFloat(req.Amount),
		Note:      req.Note,
		CreatedAt: time.Now(),
	}
}

func (t *transactionsService) Transaction(c *gin.Context, req *models.TransactionRequest) error {

	if !req.TransactionType.IsValidTransactionType() {
		return customError.ErrInvalidTransactionType
	}

	if err := t.Topup(c, req); err != nil {
		return err
	}

	return nil
}

func (t *transactionsService) Topup(c *gin.Context, req *models.TransactionRequest) (err error) {
	ctxUser := utils.GetContext(c)

	tx := t.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			err = customError.ErrInternalServerError
		}
	}()

	profile, err := t.profileRepository.GetProfile(c, ctxUser.ID)
	if err != nil {
		log.Println("Error getting user: " + err.Error())
		tx.Rollback()
		return err
	}

	req.UserID = ctxUser.ID
	topupReq := t.constructTopup(req)
	orderReq := t.constructTopupOrder(topupReq.OrderID, req)

	_, err = t.transactionRepository.CreateTransaction(c, tx, topupReq)
	if err != nil {
		log.Println("Error creating transaction: " + err.Error())
		tx.Rollback()
		return err
	}

	_, err = t.orderRepository.CreateOrder(c, tx, orderReq)
	if err != nil {
		log.Println("Error creating order: " + err.Error())
		tx.Rollback()
		return err
	}

	client := t.midtrans.Client

	resp, errMidtrans := client.ChargeTransaction(&coreapi.ChargeReq{
		PaymentType: constant.BANK_TRANSFER,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  topupReq.OrderID,
			GrossAmt: int64(topupReq.TotalAmount.IntPart()),
		},
		BankTransfer: &coreapi.BankTransferDetails{
			Bank:     constant.BANK_BCA,
			VaNumber: profile.VANumber,
		},
	})
	if errMidtrans != nil {
		log.Println("Error charging transaction: " + errMidtrans.Message)
		tx.Rollback()
		return err
	}

	if resp.StatusCode != "200" && resp.StatusCode != "201" {
		log.Println("Error charging transaction: " + resp.StatusMessage)
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
