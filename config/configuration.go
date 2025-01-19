package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Configuration struct {
	ServiceName  string
	ServiceHost  string
	ServicePort  string
	SecretKey    string
	APIVersion   string
	APP_ENV      string
	DbConfig     *DBConfig
	TwilioConfig *TwilioConfig
	WAConfig     *WhatsappConfig
	SMPTPConfig  *SMPTPConfig
	Midtrans     *Midtrans
}

type DBConfig struct {
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
}

type TwilioConfig struct {
	AccountSID  string
	AuthToken   string
	PhoneNumber string
	BaseURL     string
}

type WhatsappConfig struct {
	AccountID   string
	AccessToken string
	PhoneNumber string
	BaseURL     string
	APIVersion  string
}

type SMPTPConfig struct {
	Host     string
	Port     string
	Sender   string
	Email    string
	Password string
}

type Midtrans struct {
	ServerKey  string
	ClientKey  string
	MerchantID string
	BaseURL    string
	APIVersion string
}

func NewConfiguration() *Configuration {
	return getConfig()
}

func getConfig() *Configuration {
	config := godotenv.Load()
	if err := config; err != nil {
		fmt.Println("Error loading .env file:", err)
		panic(err)
	}

	appENV := os.Getenv("APP_ENV")
	if appENV == "" {
		appENV = "debug"
	}

	return &Configuration{
		ServiceName: os.Getenv("SERVICE_NAME"),
		ServiceHost: os.Getenv("SERVICE_HOST"),
		ServicePort: os.Getenv("SERVICE_PORT"),
		SecretKey:   os.Getenv("SECRET_KEY"),
		APIVersion:  os.Getenv("API_VERSION"),
		APP_ENV:     appENV,
		DbConfig: &DBConfig{
			DbHost:     os.Getenv("DB_HOST"),
			DbPort:     os.Getenv("DB_PORT"),
			DbUser:     os.Getenv("DB_USER"),
			DbPassword: os.Getenv("DB_PASSWORD"),
			DbName:     os.Getenv("DB_NAME"),
		},
		TwilioConfig: &TwilioConfig{
			AccountSID:  os.Getenv("TWILIO_ACCOUNT_SID"),
			AuthToken:   os.Getenv("TWILIO_AUTH_TOKEN"),
			PhoneNumber: os.Getenv("TWILIO_PHONE_NUMBER"),
			BaseURL:     os.Getenv("TWILIO_BASE_URL"),
		},
		WAConfig: &WhatsappConfig{
			AccountID:   os.Getenv("WHATSAPP_ACCOUNT_ID"),
			AccessToken: os.Getenv("WHATSAPP_ACCESS_TOKEN"),
			PhoneNumber: os.Getenv("WHATSAPP_PHONE_NUMBER"),
			BaseURL:     os.Getenv("WHATSAPP_BASE_URL"),
			APIVersion:  os.Getenv("WHATSAPP_API_VERSION"),
		},
		SMPTPConfig: &SMPTPConfig{
			Host:     os.Getenv("SMTP_HOST"),
			Port:     os.Getenv("SMTP_PORT"),
			Sender:   os.Getenv("SMTP_SENDER"),
			Email:    os.Getenv("SMTP_EMAIL"),
			Password: os.Getenv("SMTP_PASSWORD"),
		},
		Midtrans: &Midtrans{
			ServerKey:  os.Getenv("MIDTRANS_SERVER_KEY"),
			ClientKey:  os.Getenv("MIDTRANS_CLIENT_KEY"),
			MerchantID: os.Getenv("MIDTRANS_MERCHANT_ID"),
			BaseURL:    os.Getenv("MIDTRANS_BASE_URL"),
			APIVersion: os.Getenv("MIDTRANS_API_VERSION"),
		},
	}
}
