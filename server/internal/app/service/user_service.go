package service

import (
	"errors"
	"log"
	"net/http"

	"github.com/akers1023/Smart-Waste-Management-System/internal/app/models"
	"github.com/akers1023/Smart-Waste-Management-System/internal/app/repository"
	"github.com/akers1023/Smart-Waste-Management-System/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type UserService struct {
	UserRepo *repository.UserRepository
}

func NewUserService(UserRepo *repository.UserRepository) *UserService {
	return &UserService{UserRepo: UserRepo}
}

func (us *UserService) RegisterUser(ctx *fiber.Ctx, user models.User) error {
	err := us.UserRepo.CreateUser(ctx, user)
	if err != nil {
		log.Println("Error registering user:", err)
		return errors.New("failed to register user")
	}
	return nil
}

func (us *UserService) Update(ctx *fiber.Ctx, user models.User) (*models.User, error) {
	foundUser, err := us.UserRepo.UpdateUser(ctx, user)
	if err != nil {
		return nil, utils.HandleErrorResponse(ctx, http.StatusBadRequest, "Request Login failed")
	}

	return foundUser, nil
}

func (us *UserService) LoginUser(ctx *fiber.Ctx, user models.User) (*models.User, error) {
	foundUser, err := us.UserRepo.SigninByPhoneNumber(ctx, user)
	if err != nil {
		return nil, utils.HandleErrorResponse(ctx, http.StatusBadRequest, "Request Login failed")
	}
	return foundUser, nil
}

func (us *UserService) GetUserByID(ctx *fiber.Ctx, userID string) (*models.User, error) {
	if err := utils.MatchUserTypeToUID(ctx, userID); err != nil {
		return nil, utils.HandleErrorResponse(ctx, http.StatusBadRequest, "Request Get failed")
	}
	user, err := us.UserRepo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, utils.HandleErrorResponse(ctx, http.StatusBadRequest, "Request abc failed")
	}
	return user, nil
}

// func (us *UserService) GetUserByID(ctx context.Context, userID string) (*models.User, error) {
// 	// // Gọi hàm từ repository để lấy thông tin user từ cơ sở dữ liệu
// 	// user, err := us.UserRepo.GetUserByID(ctx, userID)
// 	// if err != nil {
// 	// 	log.Println("Error getting user by ID:", err)
// 	// 	return nil, errors.New("failed to get user by ID")
// 	// }

// 	// // Thêm các logic khác nếu cần

// 	return user, nil
// }
