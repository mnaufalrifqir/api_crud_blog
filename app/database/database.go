package database

import (
	"api_crud_blog/app/model"
	"api_crud_blog/app/utils"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

// ConnectDB is a function to connect to database
func ConnectDB() {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		utils.GetConfig("DB_HOST"),
		utils.GetConfig("DB_USER"),
		utils.GetConfig("DB_PASSWORD"),
		utils.GetConfig("DB_NAME"),
		utils.GetConfig("DB_PORT"),
		utils.GetConfig("DB_SSL_MODE"),
		utils.GetConfig("DB_TIMEZONE"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

// InitialMigration is a function to migrate database
func InitialMigration() {
	err := DB.AutoMigrate(&model.File{}, &model.User{}, &model.Post{}, &model.Category{}, &model.PostCategory{})

	if err != nil {
		log.Printf("Error when migrating the database: %v", err)
	}
}