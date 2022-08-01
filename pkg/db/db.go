package db

import (
	"log"
	"todocrud/pkg/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL := "postgres://postgres:secret@localhost:5432/postgres"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Departments{})
	db.AutoMigrate(&models.Doctor{})
	db.AutoMigrate(&models.Guests{})

	return db
}
