package mover

import (
	"os"
	"path/filepath"
)

func MoveFileToFolder(filePath, folderName string) error {
	dir := filepath.Dir(filePath)
	processedDir := filepath.Join(dir, folderName)
	if err := os.MkdirAll(processedDir, 0755); err != nil {
		return err
	}
	newFilePath := filepath.Join(processedDir, filepath.Base(filePath))
	return os.Rename(filePath, newFilePath)
}
