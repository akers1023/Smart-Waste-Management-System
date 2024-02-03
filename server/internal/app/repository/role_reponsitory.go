package repository

import "gorm.io/gorm"

type RoleRepository struct {
	Repository
}

// NewUserRepository tạo một thể hiện mới của UserRepository
func NewRoleRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{Repository: Repository{DB: db}}
}
