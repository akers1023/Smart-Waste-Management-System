package models

import "gorm.io/gorm"

type Roles struct {
	ID       string  `json:"id" gorm:"primaryKey"`
	RoleName *string `json:"roleName"`
}

func MigateRoles(db *gorm.DB) error {
	return db.AutoMigrate(&Roles{})
}
