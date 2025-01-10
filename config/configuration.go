package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Configuration struct {
	ServiceName string
	ServiceHost string
	ServicePort string
	SecretKey   string
	DbConfig    *DBConfig
}

type DBConfig struct {
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
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
	}
}
