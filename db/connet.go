package db

import (
	"ensweb_crud_demo/model"

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

err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Error("Failed to migrate database:", "err", err)
		return nil, err
	}
	log.Info("Database migration completed successfully!")
	
	return db, nil
}
