package external

import (
	"fmt"
	"log"
	"net/smtp"

	"github.com/deigo96/e-wallet.git/config"
)

type EmailService interface {
	SendEmail(to string, subject string, message string) error
}

type emailService struct {
	config *config.Configuration
}

func NewEmailService(config *config.Configuration) EmailService {
	return &emailService{config: config}
}

func (e *emailService) Auth() smtp.Auth {
	return smtp.PlainAuth(
		"",
		e.config.SMPTPConfig.Email,
		e.config.SMPTPConfig.Password,
		e.config.SMPTPConfig.Host,
	)
}

func (e *emailService) Address() string {
	return fmt.Sprintf("%s:%s", e.config.SMPTPConfig.Host, e.config.SMPTPConfig.Port)
}

func (e *emailService) SendEmail(to string, subject string, message string) error {
	headers := "MIME-Version: 1.0\r\nContent-Type: text/html; charset=\"UTF-8\"\r\n"
	emailBody := headers + "\r\n" + message

	body := "From: " + e.config.SMPTPConfig.Sender + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		emailBody

	err := smtp.SendMail(e.Address(), e.Auth(), e.config.SMPTPConfig.Email, []string{to}, []byte(body))
	if err != nil {
		log.Println("Error sending email: " + err.Error())
		return err
	}

	log.Println("Email sent successfully")

	return nil
}
