package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID            string    `json:"id" gorm:"primaryKey"`
	Full_name     string    `json:"full_name"`
	First_name    *string   `json:"first_name"`
	Middle_name   *string   `json:"middle_name"`
	Last_name     *string   `json:"last_name"`
	Date_of_birth time.Time `json:"date_of_birth"`
	Email         *string   `json:"email" validate:"email,required"`
	Phone         *string   `json:"phone" validate:"phone,"`
	Password      *string   `json:"password" validate:"required,min=6"`
	Total_points  *float64  `json:"total_points"`
	Token         *string   `json:"token"`
	RefreshToken  *string   `json:"refresh_token"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func MigrateUser(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}
