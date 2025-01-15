package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Configuration struct {
	ServiceName  string
	ServiceHost  string
	ServicePort  string
	SecretKey    string
	DbConfig     *DBConfig
	TwilioConfig *TwilioConfig
	WAConfig     *WhatsappConfig
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

func NewConfiguration() *Configuration {
	return getConfig()
}

func getConfig() *Configuration {
	config := godotenv.Load()
	if err := config; err != nil {
		panic(err)
	}

	return &Configuration{
		ServiceName: os.Getenv("SERVICE_NAME"),
		ServiceHost: os.Getenv("SERVICE_HOST"),
		ServicePort: os.Getenv("SERVICE_PORT"),
		SecretKey:   os.Getenv("SECRET_KEY"),
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
	}
}
