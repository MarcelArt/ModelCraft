package scaffold

import (
	"log"
	"os"
	"strings"
)

const routeTemplate = `
package api_routes

import (
	"${moduleName}/database"
	api_handlers "${moduleName}/handlers/api"
	"${moduleName}/middlewares"
	"${moduleName}/repositories"
	"github.com/gofiber/fiber/v2"
)

func Setup${modelName}Routes(api fiber.Router, auth *middlewares.AuthMiddleware) {
	h := api_handlers.New${modelName}Handler(repositories.New${modelName}Repo(database.GetDB()))

	g := api.Group("/${handlerRoute}")
	g.Get("/", auth.ProtectedAPI, h.Read)
	g.Get("/:id", auth.ProtectedAPI, h.GetByID)
	g.Post("/", auth.ProtectedAPI, h.Create)
	g.Put("/:id", auth.ProtectedAPI, h.Update)
	g.Delete("/:id", auth.ProtectedAPI, h.Delete)
}
`

func ScaffoldRoute(modelName string, handlerRoute string) {
	filename := "routes/api/" + ToSeparateByCharLowered(modelName, '_') + ".route.go"
	log.Printf("Generating route file: %s", filename)

	newRoute := strings.ReplaceAll(routeTemplate, "${modelName}", modelName)
	newRoute = strings.ReplaceAll(newRoute, "${handlerRoute}", handlerRoute)
	newRoute = strings.ReplaceAll(newRoute, "${moduleName}", moduleName)

	if err := os.WriteFile(filename, []byte(newRoute), 0644); err != nil {
		panic("Failed writing route file")
	}
}
