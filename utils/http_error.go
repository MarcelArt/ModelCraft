package utils

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func StatusCodeByError(err error) int {
	if err == nil {
		return fiber.StatusOK
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.StatusNotFound

	}
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return fiber.StatusUnauthorized

	}

	return fiber.StatusInternalServerError
}
