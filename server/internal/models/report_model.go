package models

import (
	"time"

	"gorm.io/gorm"
)

type Report struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func MigrateReport(db *gorm.DB) error {	
	return db.AutoMigrate(&Report{})
}
