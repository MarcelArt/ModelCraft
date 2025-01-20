package middlewares

import (
	"strings"

	"github.com/MarcelArt/app_standard/repositories"
	"github.com/MarcelArt/app_standard/utils"
	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware struct {
	uRepo repositories.IUserRepo
}

func NewAuthMiddleware(uRepo repositories.IUserRepo) *AuthMiddleware {
	return &AuthMiddleware{
		uRepo: uRepo,
	}
}

func (m *AuthMiddleware) ProtectedAPI(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	idToken := strings.TrimSpace(strings.Replace(token, "Bearer", "", 1))
	claims, err := utils.ParseToken(idToken)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err.Error())
	}

	c.Locals("userId", claims["userId"])
	return c.Next()
}
