package main

import (
	"os"

	"github.com/MarcelArt/app_standard/cmd"
	_ "github.com/MarcelArt/app_standard/config"
	_ "github.com/MarcelArt/app_standard/docs"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /api
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	args := os.Args
	argsLength := len(args)
	if argsLength > 1 {
		cmd.Manager(args)
	} else {
		cmd.Serve()
	}

}

// func serve() {
// 	database.ConnectDB()

// 	app := fiber.New()

// 	routes.SetupRoutes(app)

// 	log.Printf("Listening to http://localhost:%s", config.Env.PORT)
// 	log.Fatal(app.Listen(fmt.Sprintf(":%s", config.Env.PORT)))
// }

// func cmdManager(args []string) {
// 	argsLength := len(args)
// 	fn := args[1]
// 	switch fn {
// 	case "scaffold":
// 		if argsLength < 3 {
// 			fmt.Println("Please input model name")
// 			os.Exit(0)
// 		}
// 		scaffolder(args[2])
// 	default:
// 		fmt.Println("Unknown command")
// 		os.Exit(0)
// 	}
// }

// func scaffolder(modelName string) {
// 	modelCamel := scaffold.ToCamelCase(modelName)
// 	modelSnake := scaffold.ToSeparateByCharLowered(modelCamel, '_')
// 	handlerRoute := scaffold.ToSeparateByCharLowered(modelName, '-')
// 	scaffold.ScaffoldModel(modelName, modelCamel, modelSnake)
// 	scaffold.ScaffoldRepo(modelName, modelCamel)
// 	scaffold.ScaffoldHandler(modelName, handlerRoute)
// 	scaffold.ScaffoldRoute(modelName, handlerRoute)
// 	log.Println("Successfully scaffolded")
// }
