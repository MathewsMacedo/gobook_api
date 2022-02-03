package db

import (
	"fmt"
	"log"
	"os"

	"github.com/mathewsmacedo/go_api/app/models"
	"github.com/mathewsmacedo/go_api/config/environment"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var connection *gorm.DB

func Init() {
	viper.SetConfigName("database")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}

	env := environment.Current()

	host := viper.GetString(env + ".host")
	port := viper.GetString(env + ".port")
	database := viper.GetString(env + ".database")
	password := viper.GetString(env + ".password")
	user := viper.GetString(env + ".user")

	dbinfo := fmt.Sprintf("host=%s port=%s dbname=%s sslmode=disable",
		host,
		port,
		database,
	)

	if user != "" {
		dbinfo += " user=" + user
	}

	if password != "" {
		password += " user=" + password
	}

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
