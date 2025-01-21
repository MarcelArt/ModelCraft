package api_handlers

import (
	"github.com/MarcelArt/ModelCraft/models"
	"github.com/MarcelArt/ModelCraft/repositories"
	"github.com/MarcelArt/ModelCraft/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type BaseCrudHandler[TModel any, TDto models.IDTO, TPage any] struct {
	repo      repositories.IBaseCrudRepo[TModel, TDto, TPage]
	validator *validator.Validate
}

func (h *BaseCrudHandler[TModel, TDto, TPage]) Create(c *fiber.Ctx) error {
	var dto TDto

	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.NewJSONResponse(err, ""))
	}

	if err := h.validator.Struct(dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.NewJSONResponse(err, ""))
	}

	id, err := h.repo.Create(dto)
	if err != nil {
		return c.Status(utils.StatusCodeByError(err)).JSON(models.NewJSONResponse(err, ""))
	}

	return c.Status(fiber.StatusCreated).JSON(models.NewJSONResponse(fiber.Map{"ID": id}, "Created successfully"))
}

// Read retrieves a list of resources
// @Summary Get a list of resources
// @Description Get a list of resources
// @Tags CRUDs
// @Accept json
// @Produce json
// @Param resources path string true "Resource route" Enums(user, authorized-device)
// @Param page query int false "Page"
// @Param size query int false "Size"
// @Param sort query string false "Sort"
// @Param filters query string false "Filter"
// @Success 200 {array} interface{}
// @Router /{resources} [get]
func (h *BaseCrudHandler[TModel, TDto, TPage]) Read(c *fiber.Ctx) error {
	var dest []TPage

	page := h.repo.Read(c, dest)

	return c.Status(fiber.StatusOK).JSON(page)
}

func (h *BaseCrudHandler[TModel, TDto, TPage]) Update(c *fiber.Ctx) error {
	id := c.Params("id")

	var dto TDto

	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.NewJSONResponse(err, ""))
	}

	if err := h.validator.Struct(dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.NewJSONResponse(err, ""))
	}

	if err := h.repo.Update(id, &dto); err != nil {
		return c.Status(utils.StatusCodeByError(err)).JSON(models.NewJSONResponse(err, ""))
	}

	return c.Status(fiber.StatusOK).JSON(dto)
}

func (h *BaseCrudHandler[TModel, TDto, TPage]) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	model, err := h.repo.Delete(id)

	if err != nil {
		return c.Status(utils.StatusCodeByError(err)).JSON(models.NewJSONResponse(err, ""))
	}

	return c.Status(fiber.StatusOK).JSON(models.NewJSONResponse(model, "Deleted successfully"))
}

func (h *BaseCrudHandler[TModel, TDto, TPage]) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")

	model, err := h.repo.GetByID(id)

	if err != nil {
		return c.Status(utils.StatusCodeByError(err)).JSON(models.NewJSONResponse(err, ""))
	}

	return c.Status(fiber.StatusOK).JSON(models.NewJSONResponse(model, "Data retrieved successfully"))
}
