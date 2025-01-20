package utils

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func StatusCodeByError(err error) int {
	if err == nil {
		return fiber.StatusOK
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.StatusNotFound

	}

	return fiber.StatusInternalServerError
}
