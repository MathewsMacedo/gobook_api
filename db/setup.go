package db

import (
	"fmt"
	"log"
	"os"

	"github.com/mathewsmacedo/go_api/app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var connection *gorm.DB

func Init() {
	host := getEnv("PG_HOST", "127.0.0.1")
	port := getEnv("PG_PORT", "5432")
	database := getEnv("PG_DB", "postgres")

	dbinfo := fmt.Sprintf("host=%s port=%s dbname=%s sslmode=disable",
		host,
		port,
		database,
	)

	db, err := gorm.Open(postgres.Open(dbinfo), &gorm.Config{})

	if err != nil {
		log.Println("Failed to connect to database")
		log.Fatalln(err)
	}

	log.Println("Database connected")

	db.AutoMigrate(&models.Book{})

	connection = db
}

func Connection() *gorm.DB {
	return connection
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
