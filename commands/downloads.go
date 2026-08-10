package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"go-keeper/helpers"
)

func CleanDownloads() {
	var homeDir string = helpers.GetHomeDir()

	fmt.Println("Creating organize folders...")
	var sortFolders [5]string = createOrganizeFolders(homeDir)

	// moves files from Downloads folder to the sort folders
	fmt.Println("Sorting files...")
	moveFilesToSortFolders()

	// delete the files in the sort folders
	fmt.Println("Deleting files in the sort folders")
	cleanSortFolders(homeDir, sortFolders)
}

func createOrganizeFolders(homeDir string) [5]string {
	var sortFolders [5]string = [5]string{
		"Code Related",
		"Documents Related",
		"Installers",
		"Media Related",
		"Others",
	}

	for _, folder := range sortFolders {
		path := filepath.Join(homeDir, "Downloads", folder)
		if err := os.MkdirAll(path, 0755); err != nil {
			fmt.Println("Error creating folder: ", err)
			return sortFolders
		}
	}

	return sortFolders
}

func cleanSortFolders(homeDir string, sortFolders [5]string) {
	for _, dir := range sortFolders {
		for _, entry := range helpers.GetFilesFromFolder(filepath.Join("Downloads", dir)) {
			fileInfo, infoErr := entry.Info()

			if infoErr != nil {
				fmt.Println("Error getting file info: ", infoErr)
				continue
			}

			if helpers.IsFileExpired(fileInfo) {
				// delete the file if expired
				path := filepath.Join(homeDir, "Downloads", dir, entry.Name())
				deleteErr := os.Remove(path)

				if deleteErr != nil {
					fmt.Println("Error deleting "+path+" : ", deleteErr)
				}
			}
		}
	}
}

func moveFilesToSortFolders() {
	var homeDir string = helpers.GetHomeDir()

	// dirName = Downloads
	for _, entry := range helpers.GetFilesFromFolder("Downloads") {
		if entry.IsDir() {
			continue
		}

		// get file metadata
		fileInfo, infoErr := entry.Info()

		if infoErr != nil {
			fmt.Println("Error getting file info: ", infoErr)
			continue
		}

		// checks file date
		if helpers.IsFileExpired(fileInfo) {
			// delete the file if expired
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

			if _, err := os.Stat(newPath); err == nil {
				fmt.Println("File already exists, skipping:", newPath)
				continue
			}

			moveErr := os.Rename(oldPath, newPath)

			if moveErr != nil {
				fmt.Println("Error moving "+oldPath+" to "+newPath+" : ", moveErr)
			}
		}

	}
}
