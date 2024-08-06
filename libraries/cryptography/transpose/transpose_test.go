package transpose_test

import (
	"kaffein/libraries/cryptography/caesarCipher"
	"kaffein/libraries/cryptography/transpose"
	"math/rand"
	"testing"
	"time"
)

func generateRandom(lengthText, lengthAlphabet, maxKey int) (string, string, int) {
	rand.Seed(time.Now().UnixNano())

	var letters []rune
	for i := 32; i <= 126; i++ {
		letters = append(letters, rune(i))
	}

	textRunes := make([]rune, lengthText)
	for i := range textRunes {
		textRunes[i] = letters[rand.Intn(len(letters))]
	}
	text := string(textRunes)

	alphabetRunes := make([]rune, lengthAlphabet)
	for i := range alphabetRunes {
		alphabetRunes[i] = letters[rand.Intn(len(letters))]
	}
	alphabet := string(alphabetRunes)

	key := rand.Intn(maxKey + 1)

	return text, alphabet, key
}

func TestNewTranspose(t *testing.T) {
	text, alphabet, key := generateRandom(1000, 140, 30)

	testCases := []struct {
		name     string
		text     string
		length   int
		alphabet string
		key      int
	}{
		{
			name:   "simple transpose",
			text:   "Hello World",
			length: 6,
		},
		{
			name:   "less than length",
			text:   "Hello",
			length: 6,
		},
		{
			name:     "transpose caesar cipher",
			text:     text,
			length:   6,
			alphabet: alphabet,
			key:      key,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ts := transpose.NewTranspose(tt.text, tt.length).Encrypt()
			ts2 := transpose.NewTranspose(ts, tt.length).Decrypt()
			if ts2 != tt.text {
				t.Errorf("want %s, got %s", tt.text, ts2)
			}

			if tt.name == "transpose caesar cipher" {
				cc, err := caesarCipher.NewCaesarCipher(tt.text, tt.alphabet, tt.key).Encrypt()
				if err != nil {
					t.Errorf("unexpected error during Caesar Cipher encoding: %v", err)
					return
				}

				ts = transpose.NewTranspose(cc, tt.length).Encrypt()
				ts2 = transpose.NewTranspose(ts, tt.length).Decrypt()

				ccDecoded, err := caesarCipher.NewCaesarCipher(ts2, tt.alphabet, tt.key).Decrypt()
				if err != nil {
					t.Errorf("unexpected error during Caesar Cipher decoding: %v", err)
					return
				}
				if ccDecoded != tt.text {
					t.Errorf("expected decoded '%s', got '%s'", tt.text, ccDecoded)
				}
			}
		})
	}
}
