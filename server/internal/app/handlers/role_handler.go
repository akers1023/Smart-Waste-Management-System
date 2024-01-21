package handlers

import "github.com/akers1023/Smart-Waste-Management-System/internal/app/service"

// Create a new role
// Delete a role (it su dung co the khoi lam)
// Update permissions role

type RoleHandler struct {
	RoleService *service.RoleService
}

func NewRoleHandler(roleService *service.RoleService) *RoleHandler {
	return &RoleHandler{RoleService: roleService}
}
