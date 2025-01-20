package scaffold

import (
	"log"
	"os"
	"strings"
)

const handlerTemplate = `
package api_handlers

import (
	"${moduleName}/models"
	"${moduleName}/repositories"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ${modelName}Handler struct {
	BaseCrudHandler[models.${modelName}, models.${modelName}DTO, models.${modelName}Page]
	repo repositories.I${modelName}Repo
}

func New${modelName}Handler(repo repositories.I${modelName}Repo) *${modelName}Handler {
	return &${modelName}Handler{
		BaseCrudHandler: BaseCrudHandler[models.${modelName}, models.${modelName}DTO, models.${modelName}Page]{
			repo: repo,
			validator: validator.New(validator.WithRequiredStructEnabled()),
		},
		repo: repo,
	}
}

// Create creates a new ${swagLower}
// @Summary Create a new ${swagLower}
// @Description Create a new ${swagLower}
// @Tags ${modelName}
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param ${modelName} body models.${modelName}DTO true "${modelName} data"
// @Success 201 {object} models.${modelName}DTO
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /${handlerRoute} [post]
func (h *${modelName}Handler) Create(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Create(c)
}

// Read retrieves a list of ${swagPlural}
// @Summary Get a list of ${swagPlural}
// @Description Get a list of ${swagPlural}
// @Tags ${modelName}
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query int false "Page"
// @Param size query int false "Size"
// @Param sort query string false "Sort"
// @Param filters query string false "Filter"
// @Success 200 {array} models.${modelName}Page
// @Router /${handlerRoute} [get]
func (h *${modelName}Handler) Read(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Read(c)
}

// Update updates an existing ${swagLower}
// @Summary Update an existing ${swagLower}
// @Description Update an existing ${swagLower}
// @Tags ${modelName}
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "${modelName} ID"
// @Param ${modelName} body models.${modelName}DTO true "${modelName} data"
// @Success 200 {object} models.${modelName}DTO
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /${handlerRoute}/{id} [put]
func (h *${modelName}Handler) Update(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Update(c)
}

// Delete deletes an existing ${swagLower}
// @Summary Delete an existing ${swagLower}
// @Description Delete an existing ${swagLower}
// @Tags ${modelName}
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "${modelName} ID"
// @Success 200 {object} models.${modelName}
// @Failure 500 {object} string
// @Router /${handlerRoute}/{id} [delete]
func (h *${modelName}Handler) Delete(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Delete(c)
}

// GetByID retrieves a ${swagLower} by ID
// @Summary Get a ${swagLower} by ID
// @Description Get a ${swagLower} by ID
// @Tags ${modelName}
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "${modelName} ID"
// @Success 200 {object} models.${modelName}
// @Failure 500 {object} string
// @Router /${handlerRoute}/{id} [get]
func (h *${modelName}Handler) GetByID(c *fiber.Ctx) error {
	return h.BaseCrudHandler.GetByID(c)
}
`

func ScaffoldHandler(modelName string, handlerRoute string) {
	filename := "handlers/api/" + ToSeparateByCharLowered(modelName, '_') + ".handler.go"
	log.Printf("Generating handler file: %s", filename)

	swagLower := ToSeparateByCharLowered(modelName, ' ')
	swagPlural := PluralizeWord(swagLower)
	newHandler := handlerTemplate
	newHandler = strings.ReplaceAll(newHandler, "${modelName}", modelName)
	newHandler = strings.ReplaceAll(newHandler, "${swagLower}", swagLower)
	newHandler = strings.ReplaceAll(newHandler, "${swagPlural}", swagPlural)
	newHandler = strings.ReplaceAll(newHandler, "${handlerRoute}", handlerRoute)
	newHandler = strings.ReplaceAll(newHandler, "${moduleName}", moduleName)

	if err := os.WriteFile(filename, []byte(newHandler), 0644); err != nil {
		panic("Failed writing handler file")
	}
}
