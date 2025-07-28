package config

import (
	"fmt"
	"log"
	"os"

	"github.com/Sudhir0302/secure_vault/services/storage/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Load() {
	godotenv.Load("../.env.storage")
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	ssl := os.Getenv("SSL_MODE")
	timezone := os.Getenv("TIME_ZONE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=%s TimeZone=%s", host, user, pass, dbname, ssl, timezone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("db not connected", err)
	}
	DB = db
	fmt.Println("db connected")
	db.AutoMigrate(&models.Storage{})
}
