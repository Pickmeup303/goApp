package cryptography_handler

import (
	"kaffein/internal/dto"
	"kaffein/libraries/common"
	"kaffein/libraries/steganography"
	"net/http"
	"strconv"
)

var (
	pageView   = "index"
	validation = common.NewValidation()
	assets     = common.Assets(nil)
)

// HandlerEncrypt handles the GET and POST requests for the index page
func HandlerEncrypt(w http.ResponseWriter, r *http.Request) {
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
		handler = nil
	} else {
		defer file.Close()
	}

	keyShifter, _ := strconv.Atoi(r.Form.Get("keyShifter"))
	reqInput := &dto.RequestEncryptInput{
		Alphabet: r.Form.Get("keyAlphabet"),
		Key:      keyShifter,
		Message:  r.Form.Get("message"),
		File:     handler,
	}

	vErrors := validation.ValidateInputVideo(reqInput, handler)
	data["formValues"] = reqInput

	if len(vErrors) > 0 {
		data["validationError"] = vErrors
		renderTemplate(w, data, pageView)
		return
	}

	// save object file to storage
	pathVideo, err := common.SaveFileToDirectory(file, "temp", handler.Filename)
	if err != nil {
		handleError(w, data, "Failed to save file: "+err.Error(), pageView)
		return
	}

	// encrypt message using algorithm caesar chiper
	encrypt, err := common.WrapperCaesarEncrypt(reqInput.Message, reqInput.Alphabet, reqInput.Key)
	if err != nil {
		handleError(w, data, "Encryption error: "+err.Error(), pageView)
		return
	}

	// embedded to frame
	videoEmbedded := steganography.NewVideoSteganoGraphy()
	outputPath, outputFileName, err := videoEmbedded.Encode(pathVideo, "temp", handler.Filename, "FFV1", "avi", encrypt)
	if err != nil {
		handleError(w, data, "Failed to encode video: "+err.Error(), pageView)
		return
	}

	if err = common.SendFile(w, outputPath, outputFileName); err != nil {
		handleError(w, data, "Failed to send file: "+err.Error(), pageView)
	}

}

func renderTemplate(w http.ResponseWriter, data map[string]interface{}, view string) {
	if err := common.LoadTemplate(w, data, view); err != nil {
		http.Error(w, "Template loading error: "+err.Error(), http.StatusInternalServerError)
	}
}

func handleError(w http.ResponseWriter, data map[string]interface{}, message string, view string) {
	vErrors := make(map[string]string)
	vErrors["Alert"] = message
	data["validationError"] = vErrors
	renderTemplate(w, data, view)
}
