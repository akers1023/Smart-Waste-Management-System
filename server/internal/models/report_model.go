package models

import (
	"time"

	"gorm.io/gorm"
)

type Report struct {
	ID             string    `json:"id" gorm:"primaryKey"`
	Description    string    `json:"description"`
	TransactionIDs string    `json:"transaction_ids" gorm:"not null"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func MigrateReport(db *gorm.DB) error {
	return db.AutoMigrate(&Report{})
}
