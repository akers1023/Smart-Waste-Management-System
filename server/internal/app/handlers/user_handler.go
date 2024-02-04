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

func (uh *UserHandler) Register(ctx *fiber.Ctx) error {
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

// Signup (Update account) by code (Staff)
func (uh *UserHandler) Signup(ctx *fiber.Ctx) error {
	user := models.User{}

	foundUser, err := uh.UserService.LoginUser(ctx, user)
	if err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusBadRequest, "Request Login failed")
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"massage": "Register Successfully!",
		"user":    foundUser,
	})
}
func (uh *UserHandler) Login(ctx *fiber.Ctx) error {
	user := models.User{}

	err := ctx.BodyParser(&user)
	if err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusUnprocessableEntity, "Invalid request payload")
	}

	foundUser, err := uh.UserService.LoginUser(ctx, user)
	if err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusBadRequest, "Request Login failed")
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"full_name":     "Hello " + *foundUser.FirstName + *foundUser.LastName,
		"message":       "Login successfully",
		"token":         foundUser.Token,
		"refresh_token": foundUser.RefreshToken})
}

// View information account by uid
func (uh *UserHandler) ViewInfomationUser(ctx *fiber.Ctx) error {
	userID := ctx.Params("id")

	user, err := uh.UserService.GetUserByID(ctx, userID)
	if err != nil {
		return utils.HandleErrorResponse(ctx, http.StatusInternalServerError, "Failed to get user by ID")
	}

	// Trả về thành công nếu không có lỗi
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "user information",
		"data":    user,
	})
}

// Signup by phone number (Owner only)

// Signup (Update account) by code (Staff)
// Signin by phone number
// Delete Account (Admin)
// Update Account

// View all accounts (Admin, Owner)
