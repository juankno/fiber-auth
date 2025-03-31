package database

import (
	"fmt"
	"os"

	"github.com/juankno/fiber-auth/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	DSN := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	var err error
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

	err = DB.AutoMigrate(&model.Book{}, &model.User{})

	if err != nil {
		panic("failed to migrate database schemas")
	}
	fmt.Println("Database Migrated")

}
