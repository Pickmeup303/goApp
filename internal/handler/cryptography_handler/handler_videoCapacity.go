package cryptography_handler

import (
	"encoding/json"
	"gocv.io/x/gocv"
	"kaffein/internal/dto"
	"kaffein/libraries/common"
	"net/http"
	"strconv"
)

func HandlerGetCapacity(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})

	file, handler, err := r.FormFile("videoCapacity")
	if err != nil {
		data["validationError"] = "Error retrieving the file"
		respondWithJSON(w, http.StatusBadRequest, data)
		return
	}
	defer file.Close()

	reqInput := &dto.RequestCapacityVideo{File: handler}
	validationErrors := validation.ValidateInputVideo(reqInput, handler)
	if len(validationErrors) > 0 {
		data["validationError"] = validationErrors
		respondWithJSON(w, http.StatusBadRequest, data)
		return
	}

	path, err := common.SaveFileToDirectory(file, "temp", handler.Filename)
	if err != nil {
		data["validationError"] = "Error saving file: " + err.Error()
		respondWithJSON(w, http.StatusBadRequest, data)
		return
	}

	video, err := gocv.VideoCaptureFile(path)
	if err != nil {
		data["validationError"] = "Error opening video file: " + err.Error()
		respondWithJSON(w, http.StatusBadRequest, data)
		return
	}
	defer video.Close()

	if !video.IsOpened() {
		data["validationError"] = "Video isn't opened"
		respondWithJSON(w, http.StatusBadRequest, data)
		return
	}

	frame := gocv.NewMat()
	defer frame.Close()

	if !video.Read(&frame) {
		data["validationError"] = "Error reading video frame"
		respondWithJSON(w, http.StatusBadRequest, data)
		return
	}

	// Calculate the maximum hidden message capacity
	frameCount := int(video.Get(gocv.VideoCaptureFrameCount))
	frameSize := frame.Rows() * frame.Cols() * 3
	maxHiddenMessage := ((frameSize * frameCount) / 8) - 1
	data["capacity"] = strconv.Itoa(maxHiddenMessage)
	respondWithJSON(w, http.StatusOK, data)
}

// Helper function to respond with JSON
func respondWithJSON(w http.ResponseWriter, status int, data map[string]interface{}) {
	response, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Error generating response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}
