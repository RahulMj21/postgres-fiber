package models

import (
	"time"

	"gorm.io/gorm"
)

type Books struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title     *string   `json:"title"`
	Author    *string   `json:"author"`
	Publisher *string   `json:"publisher"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func MigrateBooks(db *gorm.DB) error {
	err := db.AutoMigrate(&Books{})
	return err
}
