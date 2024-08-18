package cryptography_handler

import (
	"fmt"
	"io/fs"
	"kaffein/internal/dto"
	"kaffein/libraries/common"
	"kaffein/libraries/steganography"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	pageView   = "index"
	baseDir    = "temp"
	validation = common.NewValidation()
	assets     = common.Assets(nil)
)

// HandlerEncrypt handles the GET and POST requests for the index page
func HandlerEncrypt(w http.ResponseWriter, r *http.Request) {
	// add assets
	//registerAssets := map[string][]string{
	//	"js": {
	//		"/resources/js/blob.js",
	//	},
	//}

	//assets = common.Assets(registerAssets)
	data := map[string]interface{}{
		"assets": assets,
	}

	switch r.Method {
	case http.MethodGet:
		renderTemplate(w, data, pageView)
	case http.MethodPost:
		handlePost(w, r, data)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handlePost(w http.ResponseWriter, r *http.Request, data map[string]interface{}) {
	file, handler, err := r.FormFile("file")
	if err != nil {
		file = nil
	} else {
		defer file.Close()
	}

	keyShifter, _ := strconv.Atoi(r.Form.Get("keyShifter"))
	keyTranspose, _ := strconv.Atoi(r.Form.Get("keyTranspose"))
	reqInput := &dto.RequestEncryptInput{
		Alphabet:     r.Form.Get("keyAlphabet"),
		KeyShifter:   keyShifter,
		KeyTranspose: keyTranspose,
		Message:      r.Form.Get("message"),
		File:         handler,
	}

	vErrors := validation.ValidateInputVideo(reqInput, handler)
	data["formValues"] = reqInput

	if len(vErrors) > 0 {
		data["validationError"] = vErrors
		renderTemplate(w, data, pageView)
		return
	}

	// Save file to directory
	pathVideo, err := common.SaveFileToDirectory(file, baseDir, handler.Filename)
	if err != nil {
		handleError(w, data, "Failed to save file: "+err.Error(), pageView)
		return
	}

	// Encrypt message using Caesar cipher
	encrypt, err := common.WrapperCaesarEncrypt(reqInput.Message, reqInput.Alphabet, reqInput.KeyShifter, reqInput.KeyTranspose)
	if err != nil {
		handleError(w, data, "Encryption error: "+err.Error(), pageView)
		return
	}

	// Embed message into video frames
	videoEmbedded := steganography.NewVideoSteganoGraphy()
	outputPath, outputFileName, err := videoEmbedded.Encode(pathVideo, baseDir, handler.Filename, "FFV1", "avi", encrypt)
	if err != nil {
		handleError(w, data, "Failed to encode video: "+err.Error(), pageView)
		return
	}

	if err = common.SendFile(w, outputPath, outputFileName); err != nil {
		handleError(w, data, "Failed to send file: "+err.Error(), pageView)
		return
	}

	// Delete the file after success
	go func() {
		err := filepath.Walk(baseDir, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.TrimSuffix(info.Name(), filepath.Ext(info.Name())) == strings.TrimSuffix(handler.Filename, filepath.Ext(handler.Filename)) {
				if err := os.Remove(path); err != nil {
					fmt.Printf("Failed to delete %s: %v\n", path, err)
				} else {
					fmt.Printf("Deleted %s\n", path)
				}
			}
			return nil
		})
		if err != nil {
			fmt.Printf("Error walking the path %v: %v\n", baseDir, err)
		}
	}()
}

func renderTemplate(w http.ResponseWriter, data map[string]interface{}, view string) {
	if w.Header().Get("Content-Type") == "" {
		if err := common.LoadTemplate(w, data, view); err != nil {
			http.Error(w, "Template loading error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

func handleError(w http.ResponseWriter, data map[string]interface{}, message string, view string) {
	vErrors := make(map[string]string)
	vErrors["Alert"] = message
	data["validationError"] = vErrors
	renderTemplate(w, data, view)
}
