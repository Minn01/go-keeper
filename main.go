package main

import (
	"fmt"
	"os"

	"go-keeper/commands"
	"go-keeper/config"
)

func main() {
	cfg, cfgErr := config.Load()

	if cfgErr != nil {
		fmt.Println("Error loading config:", cfgErr)
		return
	}

	if len(os.Args) < 3 {
		fmt.Println("usage: gokeeper <command>")
		return
	}

	switch os.Args[1] {
	case "downloads":
		commands.CleanDownloads(cfg)
	case "desktop":
		commands.CleanDesktop(cfg)
	case "config":
		switch os.Args[2] {
			case "show":
				// show current configuration
				config.PrintConfig(cfg)
			case "path":
				// show path of config file
				config.PrintConfigFilePath()
		}
	default:
		fmt.Println("Unknown command")
	}
}
