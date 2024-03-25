package database

import (
	"final-project/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	user     = "postgres"
	password = "root"
	dbPort   = "5432"
	dbName   = "final-project"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user = %s password = %s dbname = %s port = %s sslmode=disable", host, user, password, dbName, dbPort)
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database : ", err)
	}

	db.Debug().AutoMigrate(models.User{}, models.Comment{}, models.Photo{}, models.SocialMedia{})
}

func GetDB() *gorm.DB {
	return db
}
