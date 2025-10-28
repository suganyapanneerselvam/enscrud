package model

import (
	"github.com/EnsurityTechnologies/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID    uuid.UUID `gorm:"column:ID;primaryKey;" json:"id"`
	Name  string    `gorm:"column:Name" json:"name"`
	Email string    `gorm:"column:Email" json:"email"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}
