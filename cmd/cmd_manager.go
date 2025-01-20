package cmd

import (
	"fmt"
	"os"
)

func Manager(args []string) {
	argsLength := len(args)
	fn := args[1]
	switch fn {
	case "scaffold":
		if argsLength < 3 {
			fmt.Println("Please input model name")
			os.Exit(0)
		}
		Scaffolder(args[2])
	case "migrate":
		if argsLength < 3 {
			fmt.Println("Please input up or down")
			os.Exit(0)
		}
		Migrate(args[2])
	default:
		fmt.Println("Unknown command")
		os.Exit(0)
	}
}
