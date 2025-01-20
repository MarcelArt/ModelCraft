package cmd

import (
	"os"

	"github.com/MarcelArt/app_standard/database"
)

func Migrate(arg string) {
	switch arg {
	case "up":
		database.ConnectDB()
		database.MigrateDB()
	case "down":
		database.ConnectDB()
		database.DropDB()
	default:
		println("Unknown command")
		os.Exit(0)
	}
}
