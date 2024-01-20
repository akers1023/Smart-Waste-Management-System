package handlers

import (
	"net/http"

	"github.com/akers1023/Smart-Waste-Management-System/internal/app/models"
	"github.com/akers1023/Smart-Waste-Management-System/internal/app/service"
	"github.com/akers1023/Smart-Waste-Management-System/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

func (uh *UserHandler) RegisterUser(ctx *fiber.Ctx) error {
	user := models.User{}

	// Đọc dữ liệu từ request body và chuyển đổi thành đối tượng User
	err := ctx.BodyParser(&user)
	if err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusUnprocessableEntity, "Invalid request payload")
	}

	// Gọi hàm đăng ký user từ UserService
	if err := uh.UserService.RegisterUser(ctx, user); err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusInternalServerError, "Failed to register user")
	}

	// Trả về thành công nếu không có lỗi
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "user has been added"})
}

// Signup by phone number (Owner only)

// Signup (Update account) by code (Staff)
// Signin by phone number
// Delete Account (Admin)
// Update Account
// View information account
// View all accounts (Admin, Owner)