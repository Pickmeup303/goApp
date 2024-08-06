package cryptography_handler

import (
	"kaffein/internal/dto"
	"kaffein/libraries/common"
	"kaffein/libraries/steganography"
	"net/http"
	"strconv"
)

var pageViewDecrypt = "decrypt"

// HandlerDecrypt handles the GET and POST requests for the index page
func HandlerDecrypt(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"assets": assets,
	}

	switch r.Method {
	case http.MethodGet:
		renderTemplate(w, data, pageViewDecrypt)
	case http.MethodPost:
		handleDecrypt(w, r, data)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleDecrypt(w http.ResponseWriter, r *http.Request, data map[string]interface{}) {
	file, handler, err := r.FormFile("file")
	if err != nil {
		handler = nil
	} else {
		defer file.Close()
	}

	keyShifter, _ := strconv.Atoi(r.Form.Get("keyShifter"))
	reqInput := &dto.RequestDecryptInput{
		Alphabet: r.Form.Get("keyAlphabet"),
		Key:      keyShifter,
		File:     handler,
	}

	vErrors := validation.ValidateInputVideo(reqInput, handler)
	data["formValues"] = reqInput

	if len(vErrors) > 0 {
		data["validationError"] = vErrors
		renderTemplate(w, data, pageViewDecrypt)
		return
	}

	// save object file to directory path
	pathVideo, err := common.SaveFileToDirectory(file, "temp", handler.Filename)
	if err != nil {
		handleError(w, data, "Failed to save file: "+err.Error(), pageViewDecrypt)
		return
	}

	// extracted hidden message
	videoExtracted := steganography.NewVideoSteganoGraphy()
	extracted, err := videoExtracted.Decode(pathVideo)
	if err != nil {
		handleError(w, data, "Extracted error: "+err.Error(), pageViewDecrypt)
		return
	}

	// decrypt message
	plainText, err := common.WrapperCaesarDecrypt(extracted, reqInput.Alphabet, reqInput.Key)
	if err != nil {
		handleError(w, data, "Decode error: "+err.Error(), pageViewDecrypt)
		return
	}

	data["plainText"] = plainText
	renderTemplate(w, data, pageViewDecrypt)
}
