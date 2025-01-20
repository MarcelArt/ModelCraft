package api_handlers

import (
	"github.com/MarcelArt/app_standard/models"
	"github.com/MarcelArt/app_standard/repositories"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AuthorizedDeviceHandler struct {
	BaseCrudHandler[models.AuthorizedDevice, models.AuthorizedDeviceDTO, models.AuthorizedDevicePage]
	repo repositories.IAuthorizedDeviceRepo
}

func NewAuthorizedDeviceHandler(repo repositories.IAuthorizedDeviceRepo) *AuthorizedDeviceHandler {
	return &AuthorizedDeviceHandler{
		BaseCrudHandler: BaseCrudHandler[models.AuthorizedDevice, models.AuthorizedDeviceDTO, models.AuthorizedDevicePage]{
			repo:      repo,
			validator: validator.New(validator.WithRequiredStructEnabled()),
		},
		repo: repo,
	}
}

// Create creates a new authorized device
// @Summary Create a new authorized device
// @Description Create a new authorized device
// @Tags AuthorizedDevice
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param AuthorizedDevice body models.AuthorizedDeviceDTO true "AuthorizedDevice data"
// @Success 201 {object} models.AuthorizedDeviceDTO
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /authorized-device [post]
func (h *AuthorizedDeviceHandler) Create(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Create(c)
}

// Read retrieves a list of authorized devices
// @Summary Get a list of authorized devices
// @Description Get a list of authorized devices
// @Tags AuthorizedDevice
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query int false "Page"
// @Param size query int false "Size"
// @Param sort query string false "Sort"
// @Param filters query string false "Filter"
// @Success 200 {array} models.AuthorizedDevicePage
// @Router /authorized-device [get]
func (h *AuthorizedDeviceHandler) Read(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Read(c)
}

// Update updates an existing authorized device
// @Summary Update an existing authorized device
// @Description Update an existing authorized device
// @Tags AuthorizedDevice
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "AuthorizedDevice ID"
// @Param AuthorizedDevice body models.AuthorizedDeviceDTO true "AuthorizedDevice data"
// @Success 200 {object} models.AuthorizedDeviceDTO
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /authorized-device/{id} [put]
func (h *AuthorizedDeviceHandler) Update(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Update(c)
}

// Delete deletes an existing authorized device
// @Summary Delete an existing authorized device
// @Description Delete an existing authorized device
// @Tags AuthorizedDevice
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "AuthorizedDevice ID"
// @Success 200 {object} models.AuthorizedDevice
// @Failure 500 {object} string
// @Router /authorized-device/{id} [delete]
func (h *AuthorizedDeviceHandler) Delete(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Delete(c)
}

// GetByID retrieves a authorized device by ID
// @Summary Get a authorized device by ID
// @Description Get a authorized device by ID
// @Tags AuthorizedDevice
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "AuthorizedDevice ID"
// @Success 200 {object} models.AuthorizedDevice
// @Failure 500 {object} string
// @Router /authorized-device/{id} [get]
func (h *AuthorizedDeviceHandler) GetByID(c *fiber.Ctx) error {
	return h.BaseCrudHandler.GetByID(c)
}
