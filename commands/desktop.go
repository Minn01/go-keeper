package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"go-keeper/helpers"
)

func FilterDesktop() {
	fmt.Println("Filtering desktop...")

	homeDir := helpers.GetHomeDir()
	// create a Screenshots folder
	path := filepath.Join(homeDir, "Desktop", "Screenshots")

	if err := os.MkdirAll(path, 0755); err != nil {
		fmt.Println("Error creating folder: ", err)
		return
	}

	// move all the images/screenshots to the screenshots folder
	for _, entry := range helpers.GetFilesFromFolder("Desktop") {
		// check for the screenshots folder
		if entry.IsDir() {
			continue
		}

		// check for images
		if helpers.CheckFileIsImage(entry) {

			// move media files
			oldPath := filepath.Join(homeDir, "Desktop", entry.Name())
			newPath := filepath.Join(homeDir, "Desktop", "Screenshots", entry.Name())

			moveErr := os.Rename(oldPath, newPath)
			if moveErr != nil {
				fmt.Println("Error moving "+oldPath+" to "+newPath+" : ", moveErr)
				continue
			}
			fmt.Println("Moved file " + oldPath + "\nto " + newPath + "\n")
		} else {
			// move to downloads folder

			oldPath := filepath.Join(homeDir, "Desktop", entry.Name())
			newPath := filepath.Join(homeDir, "Downloads", entry.Name())

			if _, err := os.Stat(newPath); err == nil {
				fmt.Println("File already exists, skipping:", newPath)
				continue
			}

			moveErr := os.Rename(oldPath, newPath)

			if moveErr != nil {
				fmt.Println("Error moving "+oldPath+" to "+newPath+" : ", moveErr)
				continue
			}
			fmt.Println("Moved file " + oldPath + " to the Downloads folder\n")
		}
	}

	// Delete the files if expired and if not keep
	cleanScreenshots(homeDir)
}

func cleanScreenshots(homeDir string) {
	for _, entry := range helpers.GetFilesFromFolder("Desktop/Screenshots") {
		fileInfo, infoErr := entry.Info()

		if infoErr != nil {
			fmt.Println("Error getting file info: ", infoErr)
			continue
		}

		// file expiry check
		if helpers.IsFileExpired(fileInfo) && strings.HasPrefix(entry.Name(), "Screenshot") {
			// delete if file expired
			path := filepath.Join(homeDir, "Desktop", "Screenshots", entry.Name())

			deleteErr := os.Remove(path)

			if deleteErr != nil {
				fmt.Println("Error deleting "+path+" : ", deleteErr)
				continue
			}

			fmt.Println("Deleted file: " + entry.Name())
		}
	}
}
