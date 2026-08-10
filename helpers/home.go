package helpers

import (
	"fmt"
	"os"
)

func GetHomeDir() string {
	if homeDir, err := os.UserHomeDir(); err != nil {
		fmt.Println("Error getting home directory: ", err)
		return ""
	} else {
		return homeDir
	}
}
