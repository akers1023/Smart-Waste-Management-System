package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func MigrateTransaction(db *gorm.DB) error {
	return db.AutoMigrate(&Transaction{})

}
