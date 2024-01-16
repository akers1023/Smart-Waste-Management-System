package models

import (
	"time"

	"gorm.io/gorm"
)

type TrashBins struct {
	ID         string    `json:"id" gorm:"primaryKey"`
	TrashLevel *int      `json:"trashLevel"` // chua biet
	Location   *string   `json:"location"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func MigrateTrashBins(db *gorm.DB) error {
	return db.AutoMigrate(&TrashBins{})
}
