package cmd

import (
	"fmt"
	"log"

	"github.com/MarcelArt/app_standard/config"
	"github.com/MarcelArt/app_standard/database"
	"github.com/MarcelArt/app_standard/routes"
	"github.com/gofiber/fiber/v2"
)

func Serve() {
	database.ConnectDB()

	app := fiber.New()

	routes.SetupRoutes(app)

	log.Printf("Listening to http://localhost:%s", config.Env.PORT)
	log.Fatal(app.Listen(fmt.Sprintf(":%s", config.Env.PORT)))
}
