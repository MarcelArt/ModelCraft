package api_handlers

import (
	"errors"
	"strconv"
	"time"

	"github.com/MarcelArt/app_standard/models"
	"github.com/MarcelArt/app_standard/repositories"
	"github.com/MarcelArt/app_standard/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserHandler struct {
	BaseCrudHandler[models.User, models.UserDTO, models.UserPage]
	repo   repositories.IUserRepo
	adRepo repositories.IAuthorizedDeviceRepo
}

func NewUserHandler(repo repositories.IUserRepo, adRepo repositories.IAuthorizedDeviceRepo) *UserHandler {
	return &UserHandler{
		BaseCrudHandler: BaseCrudHandler[models.User, models.UserDTO, models.UserPage]{
			repo:      repo,
			validator: validator.New(validator.WithRequiredStructEnabled()),
		},
		repo:   repo,
		adRepo: adRepo,
	}
}

// Create creates a new user
// @Summary Create a new user
// @Description Create a new user
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param User body models.UserDTO true "User data"
// @Success 201 {object} models.UserDTO
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /user [post]
func (h *UserHandler) Create(c *fiber.Ctx) error {
	var user models.UserDTO
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if err := h.BaseCrudHandler.validator.Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	user.Salt = utils.RandString(10)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password+user.Salt), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(utils.StatusCodeByError(err)).JSON(err.Error())
	}
	user.Password = string(hashedPassword)

	id, err := h.repo.Create(user)
	if err != nil {
		return c.Status(utils.StatusCodeByError(err)).JSON(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"ID": id})
}

// Read retrieves a list of users
// @Summary Get a list of users
// @Description Get a list of users
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query int false "Page"
// @Param size query int false "Size"
// @Param sort query string false "Sort"
// @Param filters query string false "Filter"
// @Success 200 {array} models.UserPage
// @Router /user [get]
func (h *UserHandler) Read(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Read(c)
}

// Update updates an existing user
// @Summary Update an existing user
// @Description Update an existing user
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "User ID"
// @Param User body models.UserDTO true "User data"
// @Success 200 {object} models.UserDTO
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /user/{id} [put]
func (h *UserHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.UserDTO
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	// if err := h.BaseCrudHandler.validator.Struct(user); err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	// }

	if user.Password != "" {
		user.Salt = utils.RandString(10)
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password+user.Salt), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(utils.StatusCodeByError(err)).JSON(err.Error())
		}
		user.Password = string(hashedPassword)
	}

	if err := h.repo.Update(id, &user); err != nil {
		return c.Status(utils.StatusCodeByError(err)).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

// Delete deletes an existing user
// @Summary Delete an existing user
// @Description Delete an existing user
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Failure 500 {object} string
// @Router /user/{id} [delete]
func (h *UserHandler) Delete(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Delete(c)
}

// GetByID retrieves a user by ID
// @Summary Get a user by ID
// @Description Get a user by ID
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Failure 500 {object} string
// @Router /user/{id} [get]
func (h *UserHandler) GetByID(c *fiber.Ctx) error {
	return h.BaseCrudHandler.GetByID(c)
}

// Login is a function to login
// @Summary Login User
// @Description Login User
// @Tags User
// @Accept json
// @Produce json
// @Param input body models.LoginInput true "Login"
// @Success 200 {object} models.LoginResponse
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /user/login [post]
func (h *UserHandler) Login(c *fiber.Ctx) error {
	var user models.LoginInput
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if err := h.BaseCrudHandler.validator.Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	userDB, err := h.repo.GetByUsernameOrEmail(user.Username)
	if err != nil {
		return c.Status(utils.StatusCodeByError(err)).JSON(err.Error())
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(user.Password+userDB.Salt))
	if err != nil {
		return c.Status(utils.StatusCodeByError(err)).JSON(err.Error())
	}

	accessToken, refreshToken, err := utils.GenerateTokenPair(userDB, user.IsRemember)
	if err != nil {
		return c.Status(utils.StatusCodeByError(err)).JSON(err.Error())
	}

	_, err = h.adRepo.Create(models.AuthorizedDeviceDTO{
		RefreshToken: refreshToken,
		UserAgent:    c.Get("User-Agent"),
		Ip:           c.IP(),
		UserID:       userDB.ID,
	})
	if err != nil {
		c.Status(utils.StatusCodeByError(err)).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(models.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

// Refresh is a function to refresh expired access token
// @Summary Refreshes Tokens for User
// @Description Refreshes Tokens for User
// @Tags User
// @Accept json
// @Produce json
// @Param input body models.RefreshInput true "RefreshTokens"
// @Success 200 {object} models.LoginResponse
// @Failure 400 {string} string
// @Failure 401 {string} string
// @Failure 500 {string} string
// @Router /user/refresh [post]
func (h *UserHandler) Refresh(c *fiber.Ctx) error {
	var input models.RefreshInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if err := h.BaseCrudHandler.validator.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	claims, err := utils.ParseToken(input.RefreshToken)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err.Error())
	}

	isRemember := claims["isRemember"].(bool)
	var userID string

	device, err := h.adRepo.GetByRefreshToken(input.RefreshToken)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}
		userID = utils.ClaimsNumberToString(claims["userId"])
	}

	if userID == "" {
		userID = strconv.Itoa(int(device.ID))
	}

	user, err := h.repo.GetByID(userID)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err.Error())
	}

	accessToken, refreshToken, err := utils.GenerateTokenPair(
		models.UserDTO{
			Username: user.Username,
			Email:    user.Email,
			DTO: models.DTO{
				ID: user.ID,
			},
		},
		isRemember,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	device.UpdatedAt = time.Now()
	device.RefreshToken = refreshToken
	device.Ip = c.IP()
	device.UserAgent = c.Get("User-Agent")

	if device.ID == 0 {
		_, err := h.adRepo.Create(device)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}
	} else {
		if err := h.adRepo.Update(strconv.Itoa(int(device.ID)), &device); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}
	}

	return c.Status(fiber.StatusOK).JSON(models.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}
