package cmd

import (
	"fmt"
	"log"

	"github.com/MarcelArt/ModelCraft/config"
	"github.com/MarcelArt/ModelCraft/database"
	"github.com/MarcelArt/ModelCraft/routes"
	"github.com/gofiber/fiber/v2"
)

func Serve() {
	config.SetupENV()
	database.ConnectDB()

	app := fiber.New()

	routes.SetupRoutes(app)

	log.Printf("Listening to http://localhost:%s", config.Env.PORT)
	log.Fatal(app.Listen(fmt.Sprintf(":%s", config.Env.PORT)))
}
