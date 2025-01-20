package view_handlers

import (
	"log"
	"strconv"

	"github.com/MarcelArt/app_standard/database"
	"github.com/MarcelArt/app_standard/models"
	"github.com/MarcelArt/app_standard/repositories"
	"github.com/MarcelArt/app_standard/scaffold"
	"github.com/MarcelArt/app_standard/utils"
	"github.com/MarcelArt/app_standard/views/components"
	"github.com/MarcelArt/app_standard/views/dev_tools"
	"github.com/gofiber/fiber/v2"
)

type TableHandler struct {
	repo repositories.ITableRepo
}

func NewTableHandler(repo repositories.ITableRepo) *TableHandler {
	return &TableHandler{
		repo: repo,
	}
}

func (h *TableHandler) Index(c *fiber.Ctx) error {
	tables, err := h.repo.GetTables()
	if err != nil {
		return utils.Render(c, dev_tools.Index([]string{}))
	}

	return utils.Render(c, dev_tools.Index(tables))
}

func (h *TableHandler) DropAll(c *fiber.Ctx) error {
	if err := database.DropDB(); err != nil {
		return utils.Render(c, components.Toast(err.Error(), "error"))
	}

	return utils.Render(c, components.Toast("Database Droped", "success"))
}

func (h *TableHandler) MigrateModels(c *fiber.Ctx) error {
	if err := database.MigrateDB(); err != nil {
		return utils.Render(c, components.Toast(err.Error(), "error"))
	}

	return utils.Render(c, components.Toast("Database Droped", "success"))
}

func (h *TableHandler) CreateView(c *fiber.Ctx) error {
	return utils.Render(c, dev_tools.Create())
}

func (h *TableHandler) AddField(c *fiber.Ctx) error {
	i := c.Params("i")
	index, err := strconv.Atoi(i)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return utils.Render(c, components.ModelBuilderForm(index))
}

func (h *TableHandler) Create(c *fiber.Ctx) error {
	var model models.ModelBuilderRequest
	if err := c.BodyParser(&model); err != nil {
		log.Println(err.Error())
		return c.SendStatus(fiber.StatusBadRequest)
	}

	modelName := model.ModelName

	modelCamel := scaffold.ToCamelCase(modelName)
	modelSnake := scaffold.ToSeparateByCharLowered(modelCamel, '_')
	handlerRoute := scaffold.ToSeparateByCharLowered(modelName, '-')
	scaffold.ScaffoldModel(modelName, modelCamel, modelSnake)
	scaffold.ScaffoldRepo(modelName, modelCamel)
	scaffold.ScaffoldHandler(modelName, handlerRoute)
	scaffold.ScaffoldRoute(modelName, handlerRoute)
	log.Println("Successfully scaffolded")

	return c.SendStatus(fiber.StatusOK)
}
