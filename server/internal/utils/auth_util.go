package utils

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func CheckUserType(context *fiber.Ctx, role string) (err error) {
	userType := context.Params("role")
	err = nil
	if userType != role {
		err = errors.New("Unauthorized to access this resource")
		return err
	}
	return err
}

func MatchUserTypeToUID(context *fiber.Ctx, userID string) (err error) {
	userType := context.Params("role")
	uid := context.Params("id")

	err = nil
	if userID != uid && userType != "USER" {
		err = errors.New("Unauthorized to access this resource")
		return err
	}
	err = CheckUserType(context, userType)
	return err
}
