package models

import "gorm.io/gorm"

type Permission struct {
	ID            string  `json:"id" gorm:"primaryKey"`
	PermissonName *string `json:"permisson_name"`
}

func MigratePermission(db *gorm.DB) error {
	return db.AutoMigrate(&Permission{})
}
