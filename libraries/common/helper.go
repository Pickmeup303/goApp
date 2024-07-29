package common

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

var DestDir = "storage"

func SaveFileToStorage(file io.Reader, filename string) (string, error) {
	// Ensure the storage directory exists
	err := os.MkdirAll(DestDir, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("Error creating storage directory: %v", err)
	}

	// Create the destination file
	destFilePath := filepath.Join(DestDir, filename)
	destFile, err := os.Create(destFilePath)
	if err != nil {
		return "", fmt.Errorf("Error creating destination file: %v", err)
	}
	defer destFile.Close()

	// Copy the uploaded file to the destination file
	_, err = io.Copy(destFile, file)
	if err != nil {
		return "", fmt.Errorf("Error saving file: %v", err)
	}

	return destFilePath, nil
}

func SendFile(w http.ResponseWriter, filePath, fileName string) error {
	// Set headers to force download
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	w.Header().Set("Content-Type", "application/octet-stream")

	// Open the saved file
	savedFile, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer savedFile.Close()

	// Copy the file content to the response writer
	_, err = io.Copy(w, savedFile)
	if err != nil {
		return fmt.Errorf("error copying file to response: %v", err)
	}

	return nil
}
