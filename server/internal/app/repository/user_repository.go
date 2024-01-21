package repository

import (
	"fmt"
	"net/http"
	"time"

	"github.com/akers1023/Smart-Waste-Management-System/internal/app/models"
	"github.com/akers1023/Smart-Waste-Management-System/internal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CRUD User Account
// UserRepository đại diện cho repository của User
type UserRepository struct {
	Repository
}

// NewUserRepository tạo một thể hiện mới của UserRepository
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{Repository: Repository{DB: db}}
}

// Create Account by phone number (Owner only)
func (ur *UserRepository) CreateUser(ctx *fiber.Ctx, user models.User) error {
	// Kiểm tra nếu vai trò là Owner
	if ownerRole, ok := user.Role.(*models.Owner); ok {
		name := ownerRole.GetName()
		fmt.Printf(name)

		user.ID = uuid.New().String()

		validationErr := validator.New().Struct(user)
		if validationErr != nil {
			return utils.HandleErrorResponse(ctx, http.StatusBadRequest, "Validation failed")
		}

		token, refreshToken, _ := utils.GenerateAllTokens(*user.Phone, user.ID, *user.FirstName, *user.MiddleName, *user.LastName)
		user.Token = &token
		user.RefreshToken = &refreshToken

		password := utils.HashPassword(*user.Password)
		user.Password = &password

		user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		// err = NewUserRepository(ur.DB.Create(&user))
		err := ur.DB.Create(&user).Error
		if err != nil {
			return utils.HandleErrorResponse(ctx, http.StatusInternalServerError, "Failed to create user")
		}
		return nil
	} else {
		return utils.HandleErrorResponse(ctx, http.StatusBadRequest, "Validation failed")
	}
}

// Signup (Update account) by code (Staff)
// Signin by phone number
// Delete Account (Admin)
// Update Account
// View all accounts (Admin, Owner)

// View information account by uid
func (ur *UserRepository) GetUserByID(ctx *fiber.Ctx, userID string) (*models.User, error) {
	var user models.User

	return &user, nil
}
