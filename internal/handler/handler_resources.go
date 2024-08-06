package handler

import (
	"net/http"
	"os"
	"path/filepath"
)

// ResourcesHandler serves static files from the "static" directory
func ResourcesHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Construct the full File path by joining "static" with the request URL path
		filePath := filepath.Join("static", r.URL.Path)

		// Check if the path is a directory
		if info, err := os.Stat(filePath); err == nil {
			if info.IsDir() {
				// Return 404 for directories
				http.NotFound(w, r)
				return
			}
		} else if os.IsNotExist(err) {
			// Return 404 if the File does not exist
			http.NotFound(w, r)
			return
		}

		// Serve the File if it exists
		http.ServeFile(w, r, filePath)
	})
}
