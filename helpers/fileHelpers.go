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

func IsFileExpired(fileInfo os.FileInfo) bool {
	var fileAge time.Duration = time.Since(fileInfo.ModTime())
	// TODO: the expiration time should be loaded from the config JSON file, not hardcoded
	return fileAge.Hours() > 24*30 // 30 days
}
