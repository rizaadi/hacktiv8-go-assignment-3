package database

import (
	"fmt"
	"hacktiv8-go-assignment-3/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	HOST     = "localhost"
	USER     = "postgres"
	PASSWORD = "postgres"
	PORT     = "5432"
	DATABASE = "postgres"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", HOST, USER, PASSWORD, DATABASE, PORT)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("conectedd to db")
	db.Debug().AutoMigrate(models.Model{})
}
func GetDB() *gorm.DB {	
	return db
}
