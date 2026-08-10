package main

import (
	"fmt"
	"os"

	"go-keeper/commands"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: gokeeper <command>")
		return
	}

	switch os.Args[1] {
	case "downloads":
		commands.CleanDownloads()
	case "desktop":
		commands.FilterDesktop()
	default:
		fmt.Println("Unknown command")
	}
}
