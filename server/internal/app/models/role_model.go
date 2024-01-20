package models

import "gorm.io/gorm"

type Role struct {
	ID       string  `json:"id" gorm:"primaryKey"`
	RoleName *string `json:"role_name" validate:"required,eq=OWNER|eq=ADMIN|eq=USER"`
}

func MigateRoles(db *gorm.DB) error {
	return db.AutoMigrate(&Role{})
}