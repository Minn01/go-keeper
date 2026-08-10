package helpers

import (
	"mime"
	"os"
	"path/filepath"
	"strings"
)

var installerExtensions = map[string]bool{
	".dmg": true,
	".pkg": true,
	".app": true,
	".iso": true,
	".zip": true,
	".rar": true,
	".7z":  true,
	".tar": true,
	".gz":  true,
}
var codeExtensions = map[string]bool{
	".go":   true,
	".js":   true,
	".ts":   true,
	".tsx":  true,
	".jsx":  true,
	".py":   true,
	".java": true,
	".cs":   true,
	".cpp":  true,
	".c":    true,
	".h":    true,
	".hpp":  true,
	".rs":   true,
	".php":  true,
	".sql":  true,
	".json": true,
	".yaml": true,
	".yml":  true,
	".toml": true,
}

func GetFileExtension(file os.DirEntry) string {
	return filepath.Ext(file.Name())
}

func CheckFileIsImage(file os.DirEntry) bool {
	ext := GetFileExtension(file)
	mimeType := mime.TypeByExtension(ext)

	return strings.HasPrefix(mimeType, "image/")
}

func ClassifyFileExtension(file os.DirEntry) string {
	var sortFolder string
	ext := GetFileExtension(file)
	mimeType := mime.TypeByExtension(ext)

	if installerExtensions[ext] {
		return "Installers"
	}

	if codeExtensions[ext] {
		return "Code Related"
	}

	switch {
	case strings.HasPrefix(mimeType, "image/"),
		strings.HasPrefix(mimeType, "video/"),
		strings.HasPrefix(mimeType, "audio/"):
		sortFolder = "Media Related"

	case strings.HasPrefix(mimeType, "text/"),
		mimeType == "application/pdf":
		sortFolder = "Documents Related"

	default:
		sortFolder = "Others"
	}

	return sortFolder
}
