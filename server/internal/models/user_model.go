package models

import (
	"time"

	"gorm.io/gorm"
)

// username va email bo sung xac thuc dang nhap sau
// xem lai category nen dung string hay dung pointer string
type User struct {
	ID           string    `json:"id" gorm:"primaryKey"`
	FullName     string    `json:"full_name"`
	FirstName    *string   `json:"first_name" validate:"required,min=2,max=100"`
	MiddleName   *string   `json:"middle_name" validate:"required,min=2,max=100"`
	LastName     *string   `json:"last_name" validate:"required,min=2,max=100"`
	Gender       *string   `json:"gender" validate:"required,eq=male|eq=female"`
	DateOfBirth  time.Time `json:"date_of_birth"`
	Nationality  *string   `json:"nationality" validate:"required"`
	CIN          *[12]byte `json:"cin" validate:"required,min=12"`
	POO          *string   `json:"poo" validate:"required"`
	POR          *string   `json:"por" validate:"required"`
	Email        *string   `json:"email"`    // validate: email
	Username     *string   `json:"username"` // validate: user
	Phone        *string   `json:"phone,omitempty" validate:"omitempty,phone"`
	Password     *string   `json:"password" validate:"required,min=6"`
	Category     *string   `json:"category" validate:"required,eq=fulltime,eq=parttime"`
	Token        *string   `json:"token"`
	RefreshToken *string   `json:"refresh_token"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func MigrateUser(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}
