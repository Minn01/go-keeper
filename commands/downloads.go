package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"go-keeper/helpers"
)

func CleanDownloads() {
	fmt.Println("fetching home directory...")
	homeDir := helpers.GetHomeDir()

	fmt.Println("Creating organize folders...")
	createOrganizeFolders(homeDir)

	for _, entry := range helpers.GetFilesFromFolder("Downloads") {
		if entry.IsDir() {
			continue
		}

		fileInfo, infoErr := entry.Info()

		if infoErr != nil {
			fmt.Println("Error getting file info: ", infoErr)
		}

		if helpers.IsFileExpired(fileInfo) {
			// delete the file
			path := filepath.Join(homeDir, "Downloads", entry.Name())
			deleteErr := os.Remove(path)

			if deleteErr != nil {
				fmt.Println("Error deleting "+path+" : ", deleteErr)
			}
		} else {
			// move file to the appropriate folders
			var sortFolder string = helpers.ClassifyFileExtension(entry)

			oldPath := filepath.Join(homeDir, "Downloads", entry.Name())
			newPath := filepath.Join(homeDir, "Downloads", sortFolder, entry.Name())

			moveErr := os.Rename(oldPath, newPath)

			if moveErr != nil {
				fmt.Println("Error moving "+oldPath+" to "+newPath+" : ", moveErr)
			}
		}

	}
}

func createOrganizeFolders(homeDir string) {
	var folders [5]string = [5]string{
		"Code Related",
		"Documents Related",
		"Installers",
		"Media Related",
		"Others",
	}

	for _, folder := range folders {
		path := filepath.Join(homeDir, "Donloads", folder)
		if err := os.MkdirAll(path, 0755); err != nil {
			fmt.Println("Error creating folder: ", err)
			return
		}
	}
}
