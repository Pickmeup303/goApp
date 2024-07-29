package steganography

func stringToBinary(str string) []uint8 {
	var result []uint8
	for _, char := range str {
		for i := 7; i >= 0; i-- {
			bit := (char >> uint(i)) & 1
			result = append(result, uint8(bit))
		}
	}
	return result
}

func binaryToString(bits []uint8) string {
	var result []byte
	for i := 0; i < len(bits); i += 8 {
		var byteVal uint8
		for j := 0; j < 8; j++ {
			byteVal = (byteVal << 1) | bits[i+j]
		}
		result = append(result, byteVal)
	}
	return string(result)
}
