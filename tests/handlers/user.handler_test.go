package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	api_handlers "github.com/MarcelArt/ModelCraft/handlers/api"
	"github.com/MarcelArt/ModelCraft/mocks/repositories"
	"github.com/MarcelArt/ModelCraft/mocks/services"
	"github.com/MarcelArt/ModelCraft/models"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUser(t *testing.T) {
	app := fiber.New()

	userRepoMock := new(repositories.IUserRepo)
	authDeviceRepoMock := new(repositories.IAuthorizedDeviceRepo)
	mailServiceMock := new(services.IMailService)

	userHandler := api_handlers.NewUserHandler(userRepoMock, authDeviceRepoMock, mailServiceMock)

	app.Post("api/user", userHandler.Create)

	t.Run("should return 201 when user created", func(t *testing.T) {
		mockInput := models.UserDTO{
			Username: "string",
			Password: "stringest",
			Email:    "string@yopmail.com",
		}

		body, err := json.Marshal(mockInput)
		assert.NoError(t, err)

		userRepoMock.On("Create", mock.Anything).
			Once().
			Return(uint(1), nil)

		mailServiceMock.On("SendMail", mock.Anything).
			Once().
			Return(nil)

		req := httptest.NewRequest("POST", "/api/user", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		res, err := app.Test(req)
		assert.NoError(t, err)

		var result models.JSONResponse
		err = json.NewDecoder(res.Body).Decode(&result)
		assert.NoError(t, err)

		assert.Equal(t, fiber.StatusCreated, res.StatusCode)

		userRepoMock.AssertExpectations(t)
	})

}
