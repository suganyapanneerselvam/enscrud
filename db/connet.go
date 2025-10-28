package db

import (
	"github.com/EnsurityTechnologies/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func OpenDB(log logger.Logger) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("crud.db"), &gorm.Config{})
	if err != nil {
		log.Error("Failed to connect database:", "err", err)
		return nil, err
	}
	log.Info("DB Opened")
	return db, nil
}
