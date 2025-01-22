package api_handlers

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/MarcelArt/ModelCraft/models"
	"github.com/MarcelArt/ModelCraft/repositories"
	"github.com/MarcelArt/ModelCraft/services"
	"github.com/MarcelArt/ModelCraft/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserHandler struct {
	BaseCrudHandler[models.User, models.UserDTO, models.UserPage]
	repo     repositories.IUserRepo
	adRepo   repositories.IAuthorizedDeviceRepo
	mService services.IMailService
}

func NewUserHandler(repo repositories.IUserRepo, adRepo repositories.IAuthorizedDeviceRepo, mService services.IMailService) *UserHandler {
	return &UserHandler{
		BaseCrudHandler: BaseCrudHandler[models.User, models.UserDTO, models.UserPage]{
			repo:      repo,
			validator: validator.New(validator.WithRequiredStructEnabled()),
		},
		repo:     repo,
		adRepo:   adRepo,
		mService: mService,
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
		return c.Status(fiber.StatusBadRequest).JSON(models.NewJSONResponse(err, ""))
	}

	if err := h.BaseCrudHandler.validator.Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.NewJSONResponse(err, ""))
	}

	user.Salt = utils.RandString(10)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password+user.Salt), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(utils.StatusCodeByError(err)).JSON(models.NewJSONResponse(err, ""))
	}
	user.Password = string(hashedPassword)

	id, err := h.repo.Create(user)
	if err != nil {
		return c.Status(utils.StatusCodeByError(err)).JSON(models.NewJSONResponse(err, ""))
	}

	err = h.mService.SendMail(services.Mailer{
		To:      []string{user.Email},
		Subject: "Email Verification",
		Body:    fmt.Sprintf("Please verify your email by clicking this link: <a href='%s/api/user/verify/%d'>Verify</a>", c.BaseURL(), id),
	})
	if err != nil {
		return c.Status(utils.StatusCodeByError(err)).JSON(models.NewJSONResponse(err, ""))
	}

	return c.Status(fiber.StatusCreated).JSON(models.NewJSONResponse(fiber.Map{"ID": id}, "User registered successfully"))
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
		return c.Status(fiber.StatusBadRequest).JSON(models.NewJSONResponse(err, ""))
	}

	// if err := h.BaseCrudHandler.validator.Struct(user); err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	// }

	if user.Password != "" {
		user.Salt = utils.RandString(10)
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password+user.Salt), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(utils.StatusCodeByError(err)).JSON(models.NewJSONResponse(err, ""))
		}
		user.Password = string(hashedPassword)
	}

	if err := h.repo.Update(id, &user); err != nil {
		return c.Status(utils.StatusCodeByError(err)).JSON(models.NewJSONResponse(err, ""))
	}

	return c.Status(fiber.StatusOK).JSON(models.NewJSONResponse(user, "User updated successfully"))
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
		return c.Status(fiber.StatusBadRequest).JSON(models.NewJSONResponse(err, ""))
	}

	if err := h.BaseCrudHandler.validator.Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.NewJSONResponse(err, ""))
	}

	userDB, err := h.repo.GetByUsernameOrEmail(user.Username)
	if err != nil {
		return c.Status(utils.StatusCodeByError(err)).JSON(models.NewJSONResponse(err, "Username or password is incorrect"))
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(user.Password+userDB.Salt))
	if err != nil {
		return c.Status(utils.StatusCodeByError(err)).JSON(models.NewJSONResponse(err, "Username or password is incorrect"))
	}

	accessToken, refreshToken, err := utils.GenerateTokenPair(userDB, user.IsRemember)
	if err != nil {
		return c.Status(utils.StatusCodeByError(err)).JSON(models.NewJSONResponse(err, ""))
	}

	_, err = h.adRepo.Create(models.AuthorizedDeviceDTO{
		RefreshToken: refreshToken,
		UserAgent:    c.Get("User-Agent"),
		Ip:           c.IP(),
		UserID:       userDB.ID,
	})
	if err != nil {
		c.Status(utils.StatusCodeByError(err)).JSON(models.NewJSONResponse(err, ""))
	}

	return c.Status(fiber.StatusOK).JSON(models.NewJSONResponse(models.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, "Login successfully"))
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
		return c.Status(fiber.StatusBadRequest).JSON(models.NewJSONResponse(err, ""))
	}

	if err := h.BaseCrudHandler.validator.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.NewJSONResponse(err, ""))
	}

	claims, err := utils.ParseToken(input.RefreshToken)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(models.NewJSONResponse(err, ""))
	}

	isRemember := claims["isRemember"].(bool)
	var userID string

	device, err := h.adRepo.GetByRefreshToken(input.RefreshToken)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusInternalServerError).JSON(models.NewJSONResponse(err, ""))
		}
		userID = utils.ClaimsNumberToString(claims["userId"])
	}

	if userID == "" {
		userID = strconv.Itoa(int(device.ID))
	}

	user, err := h.repo.GetByID(userID)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(models.NewJSONResponse(err, ""))
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
		return c.Status(fiber.StatusInternalServerError).JSON(models.NewJSONResponse(err, ""))
	}

	device.UpdatedAt = time.Now()
	device.RefreshToken = refreshToken
	device.Ip = c.IP()
	device.UserAgent = c.Get("User-Agent")

	if device.ID == 0 {
		_, err := h.adRepo.Create(device)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(models.NewJSONResponse(err, ""))
		}
	} else {
		if err := h.adRepo.Update(strconv.Itoa(int(device.ID)), &device); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(models.NewJSONResponse(err, ""))
		}
	}

	return c.Status(fiber.StatusOK).JSON(models.NewJSONResponse(models.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, "Tokens refreshed successfully"))
}

func (h *UserHandler) Verify(c *fiber.Ctx) error {
	id := c.Params("id")
	err := h.repo.Verify(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.NewJSONResponse(err, ""))
	}

	log.Println("User verified")
	return c.Redirect("/")
}
