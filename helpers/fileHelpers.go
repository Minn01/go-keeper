package helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func GetFilesFromFolder(folderPath string) []os.DirEntry {
	homeDir := GetHomeDir()
	files, err := os.ReadDir(filepath.Join(homeDir, folderPath))

	if err != nil {
		fmt.Println(err)
	}

	return files
}

func IsFileExpired(fileInfo os.FileInfo, expirationDays int) bool {
	var fileAge time.Duration = time.Since(fileInfo.ModTime())
	// the expiration duration comes from the json config file
	return fileAge.Hours() > float64(24*expirationDays)
}
