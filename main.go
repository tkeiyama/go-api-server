package main

import (
	"api-server/config"
	"api-server/models"
	"api-server/routes"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var err error

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.DB, err = gorm.Open(
		"mysql",
		config.DbURL(config.BuildDBConfig()),
	)
	if err != nil {
		fmt.Println("Status:", err)
	}

	defer config.DB.Close()
	config.DB.AutoMigrate(&models.User{})

	r := routes.SetupRouter()
	r.Run()
}
