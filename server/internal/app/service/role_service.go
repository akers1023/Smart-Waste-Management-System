package service

import (
	"context"
	"errors"

	"github.com/akers1023/Smart-Waste-Management-System/internal/app/models"
	"github.com/akers1023/Smart-Waste-Management-System/internal/app/repository"
)

// RoleService là một dịch vụ chung sử dụng Role.
type RoleService struct {
	RoleRepo *repository.RoleRepository
}

func NewRoleService(RoleRepo *repository.RoleRepository) *RoleService {
	return &RoleService{RoleRepo: RoleRepo}
}

func (s *RoleService) GetUserRole(ctx context.Context, roleType string) (models.Role, error) {
	// Ví dụ:
	switch roleType {
	case "Owner":
		return &models.Owner{}, nil
	case "Admin":
		return &models.Admin{}, nil
	case "Staff":
		return &models.Staff{}, nil
	default:
		return nil, errors.New("Vai trò không hợp lệ")
	}
}
