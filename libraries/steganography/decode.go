package steganography

import (
	"fmt"
	"gocv.io/x/gocv"
)

func (v *videoSteganography) Decode(targetVideo string) (string, error) {
	video, err := gocv.VideoCaptureFile(targetVideo)
	if err != nil {
		return "", fmt.Errorf("error opening video: %v", err)
	}
	defer video.Close()

	if !video.IsOpened() {
		return "", fmt.Errorf("video isn't opened")
	}

	frame := gocv.NewMat()
	defer frame.Close()

	var bits []uint8

	for {
		if !video.Read(&frame) {
			break // end of video
		}

		if frame.Empty() {
			continue
		}

		rows, cols := frame.Rows(), frame.Cols()

		for y := 0; y < rows; y++ {
			for x := 0; x < cols; x++ {
				pixel := frame.GetVecbAt(y, x)
				blue, green, red := pixel[0], pixel[1], pixel[2]

				// Collect LSBs from each color channel
				bits = append(bits, red&0x01, green&0x01, blue&0x01)
			}
		}

		// Convert collected bits to string and check if we found the emmit character
		if msg, found := binaryToString(bits); found {
			return msg, nil
		}
	}

	// If emmit character was not found, return an empty string
	return "", nil

}
