package steganography

import (
	"errors"
	"gocv.io/x/gocv"
	"os"
	"os/exec"
)

func WriteVideo(file, message, outputFile string) error {
	video, err := gocv.VideoCaptureFile(file)
	if err != nil {
		return err
	}
	defer video.Close()

	if !video.IsOpened() {
		return errors.New("video tidak terbuka")
	}

	// Membaca frame pertama untuk mendapatkan ukuran gambar
	frame := gocv.NewMat()
	defer frame.Close()

	if !video.Read(&frame) {
		return errors.New("gagal membaca frame pertama")
	}

	tempAudioFile := "temp/audio/temp.mp3"
	if err = extractAudio(file, tempAudioFile); err != nil {
		return err
	}
	defer os.Remove(tempAudioFile)
	// Inisialisasi VideoWriter
	writer, err := gocv.VideoWriterFile(outputFile, "FFV1", video.Get(gocv.VideoCaptureFPS), frame.Cols(), frame.Rows(), true)
	if err != nil {
		return err
	}
	defer writer.Close()

	// Konversi pesan menjadi biner
	messageBinary := stringToBinary(message)
	lengthMessage := len(messageBinary)

	index := 0
	for {
		if ok := video.Read(&frame); !ok {
			break
		}
		if frame.Empty() {
			continue
		}

		for x := 0; x < frame.Rows(); x++ {
			for y := 0; y < frame.Cols(); y++ {
				if index >= lengthMessage {
					break
				}

				// Mendapatkan nilai pixel
				pixel := frame.GetVecbAt(x, y)
				b, g, r := pixel[0], pixel[1], pixel[2]

				// Menyisipkan bit pesan ke dalam pixel
				if index < lengthMessage {
					b = (b & 0xFE) | messageBinary[index]
					index++
				}
				if index < lengthMessage {
					g = (g & 0xFE) | messageBinary[index]
					index++
				}
				if index < lengthMessage {
					r = (r & 0xFE) | messageBinary[index]
					index++
				}

				// Menetapkan nilai pixel yang baru
				frame.SetUCharAt(x, y*3, b)
				frame.SetUCharAt(x, y*3+1, g)
				frame.SetUCharAt(x, y*3+2, r)
			}
			if index >= lengthMessage {
				break
			}
		}

		// Menulis frame ke video keluaran
		writer.Write(frame)
	}

	if err = mergeAudio(outputFile, tempAudioFile, outputFile); err != nil {
		return err
	}

	return nil
}

func extractAudio(inputVideo, outputAudio string) error {
	// Check if the directory "temp/audio" exists, and if not, create it.
	if _, err := os.Stat("temp/audio"); os.IsNotExist(err) {
		err = os.MkdirAll("temp/audio", 0777)
		if err != nil {
			return err
		}
	}

	// Use ffmpeg to extract audio from the input video.
	command := exec.Command("ffmpeg", "-i", inputVideo, "-q:a", "0", "-map", "a", outputAudio)
	err := command.Run()
	if err != nil {
		return err
	}
	return nil
}

func mergeAudio(inputVideo, inputAudio, outputVideo string) error {
	command := exec.Command("ffmpeg", "-i", inputVideo, "-i", inputAudio, "-c", "copy", "-map", "0:v:0", "-map", "1:a:0", outputVideo)
	err := command.Run()
	if err != nil {
		return err
	}
	return nil
}
