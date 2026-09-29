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

	if len(os.Args) < 2 {
		printUsage()
		return
	}

	switch os.Args[1] {
	case "downloads":
		commands.CleanDownloads(cfg)
	case "desktop":
		commands.CleanDesktop(cfg)
	case "config":
		if len(os.Args) < 3 {
			fmt.Println("usage: gokeeper config <show|path>")
			return
		}

		switch os.Args[2] {
		case "show":
			// show current configuration
			config.PrintConfig(cfg)
		case "path":
			// show path of config file
			config.PrintConfigFilePath()
		default:
			fmt.Println("usage: gokeeper config <show|path>")
		}
	default:
		fmt.Println("unknown command:", os.Args[1])
		printUsage()
	}
}

func printUsage() {
	fmt.Println("usage: gokeeper <downloads|desktop|config>")
	fmt.Println("       gokeeper config <show|path>")
}
