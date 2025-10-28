package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(sqlite.Open("crud.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}
	log.Println("Database connected")
}
