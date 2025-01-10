package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDBConfig() *DBConfig {
	return NewConfiguration().DbConfig
}

func DBConnection(config *DBConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DbHost,
		config.DbPort,
		config.DbUser,
		config.DbPassword,
		config.DbName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect to database", err)
		panic(err)
	}

	log.Println("Connected to database")
	return db
}

func CloseConnection(db *gorm.DB) {

	if db != nil {
		db, _ := db.DB()
		db.Close()
	}

}
