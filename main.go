package main

import (
	"os"

	"github.com/MarcelArt/ModelCraft/cmd"
	_ "github.com/MarcelArt/ModelCraft/docs"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
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
