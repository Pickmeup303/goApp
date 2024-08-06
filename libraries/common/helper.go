package common

import (
	"errors"
	"fmt"
	"html/template"
	"io"
	"kaffein/config"
	"net/http"
	"os"
	"path/filepath"
)

func SaveFileToDirectory(file io.Reader, dirName, fileName string) (string, error) {
	if err := CreateDirectory(dirName); err != nil {
		return "", err
	}

	filePath := filepath.Join(dirName, fileName)
	outFile, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer outFile.Close()

	if _, err = io.Copy(outFile, file); err != nil {
		return "", err
	}

	return filePath, nil
}

func CreateDirectory(dirName string) error {
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		if err = os.MkdirAll(dirName, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
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

func Assets(st map[string][]string) map[string][]string {
	cfg, err := config.DefaultConfig()
	if err != nil {
		return map[string][]string{}
	}

	based := map[string][]string{
		"css": {
			cfg.AppURI + "/resources/css/bootstrap.min.css",
			cfg.AppURI + "/resources/css/main.css",
			"https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.0.0-beta3/css/all.min.css",
		},
		"js": {
			cfg.AppURI + "/resources/js/bootstrap.bundle.min.js",
			cfg.AppURI + "/resources/js/main.js",
		},
	}

	for key, paths := range st {
		for _, path := range paths {
			based[key] = append(based[key], cfg.AppURI+path)
		}
	}

	return based
}

func LoadTemplate(w http.ResponseWriter, data interface{}, fileName string) error {
	filePath := filepath.Join("templates", fileName+".html")

	temp, err := template.ParseFiles(filePath)
	if err != nil {
		return errors.New("failed to load template: " + err.Error())
	}
	if err := temp.Execute(w, data); err != nil {
		return errors.New("failed to execute template: " + err.Error())
	}
	return nil
}
