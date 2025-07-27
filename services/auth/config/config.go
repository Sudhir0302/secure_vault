package config

import (
	"fmt"
	"log"
	"os"

	"github.com/Sudhir0302/secure_vault.git/services/auth/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Load() {
	//loads env values from .env file and stores it in the process's env memory (temp memory)
	godotenv.Load("../.env.auth")
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	ssl := os.Getenv("SSL_MODE")
	timezone := os.Getenv("TIME_ZONE")

	//dsn means data source name, which contains all neccessary info to connect to a database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=%s TimeZone=%s", host, user, pass, dbname, ssl, timezone)
	// fmt.Println(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("could'nt connect to db : ", err)
	}
	DB = db
	fmt.Println("db connected")
	db.AutoMigrate(&models.User{})
}
