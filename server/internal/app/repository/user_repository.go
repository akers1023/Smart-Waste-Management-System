package repository

import (
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

// Check Role o service hay repository?
// Create Account by phone number (Owner only)
func (ur *UserRepository) CreateUser(ctx *fiber.Ctx, user models.User) error {
	// Kiểm tra nếu vai trò là Owner
	// if ownerRole, ok := user.Role.(*models.Owner); ok {
	// name := ownerRole.GetName()
	// fmt.Printf(name)

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
	// } else {
	// 	return utils.HandleErrorResponse(ctx, http.StatusBadRequest, "Validation failed")
	// }
}

// Signin by phone number
func (ur *UserRepository) SigninByPhoneNumber(ctx *fiber.Ctx, user models.User) (*models.User, error) {
	foundUser := models.User{}

	if err := ur.DB.Where("phone = ?", user.Phone).First(&foundUser).Error; err != nil {
		return nil, utils.HandleErrorResponse(ctx, http.StatusUnprocessableEntity, err.Error())
	}

	passwordIsValid, msg := utils.VerifyPassword(*user.Password, *foundUser.Password)
	if passwordIsValid != true {
		return nil, utils.HandleErrorResponse(ctx, http.StatusInternalServerError, msg)
	}

	if foundUser.Phone == nil {
		return nil, utils.HandleErrorResponse(ctx, http.StatusInternalServerError, "user not found")
	}

	token, refreshToken, _ := utils.GenerateAllTokens(*foundUser.Phone, foundUser.ID, *foundUser.FirstName, *foundUser.MiddleName, *foundUser.LastName)
	utils.UpdateAllTokens(ur.DB, foundUser.ID, token, refreshToken)

	return &foundUser, nil
}

// View information account by uid
func (ur *UserRepository) GetUserByID(ctx *fiber.Ctx, userID string) (*models.User, error) {
	user := models.User{}
	err := ur.DB.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return nil, utils.HandleErrorResponse(ctx, http.StatusInternalServerError, "user not found")
	}
	return &user, nil
}

// Signup (Update account) by code (Staff)
func (ur *UserRepository) UpdateUser(ctx *fiber.Ctx, user models.User) (*models.User, error) {
	updatedUser, err := ur.SigninByPhoneNumber(ctx, user)
	if err != nil {
		return nil, utils.HandleErrorResponse(ctx, http.StatusInternalServerError, "user not found")
	}

	// Save the updated user back to the database
	if err := ur.DB.Save(&user).Error; err != nil {
		return nil, utils.HandleErrorResponse(ctx, http.StatusInternalServerError, "abc failed")
	}

	return updatedUser, nil
}

// Delete Account (Admin)
// Update Account
// View all accounts (Admin, Owner)
