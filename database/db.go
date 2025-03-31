package database

import (
	"fmt"
	"log"
	"os"

	"github.com/juankno/fiber-auth/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	var DSN = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	fmt.Println(DSN)

	var err error

	DB, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connection Opened to Database")

		DB.AutoMigrate(&model.Book{}, &model.User{})
		fmt.Println("Database Migrated")
	}
}
