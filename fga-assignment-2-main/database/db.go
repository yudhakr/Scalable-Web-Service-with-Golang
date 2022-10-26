package database

import (
	"assignment2/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// grom db
var (
	host     = "kesavan.db.elephantsql.com"
	user     = "cmjnftwz"
	password = "nQxQYEIsnQsQ5RgArOtfhTespPvMNnSH"
	dbPort   = "5432"
	dbName   = "cmjnftwz"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, dbPort)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("can't connect database", err)
	}

	db.Debug().AutoMigrate(models.Order{}, models.Item{})
}

func GetDB() *gorm.DB {
	return db
}
