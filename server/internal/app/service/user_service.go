package service

import (
	"errors"
	"log"

	"github.com/akers1023/Smart-Waste-Management-System/internal/app/models"
	"github.com/akers1023/Smart-Waste-Management-System/internal/app/repository"
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
