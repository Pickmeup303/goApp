package caesarCipher_test

import (
	"kaffein/libraries/cryptography/caesarCipher"
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

func TestValidation(t *testing.T) {
	plainText, alphabet, key := generateRandom(1000, 141, 10)

	testCases := []struct {
		name      string
		plainText string
		alphabet  string
		key       int
		expected  string
	}{
		{
			name:      "validation plain text",
			plainText: "",
			alphabet:  alphabet,
			key:       key,
			expected:  "text is empty",
		},
		{
			name:      "validation alphabet",
			plainText: plainText,
			alphabet:  "short",
			key:       key,
			expected:  "alphabet must be at least 140 chars or more",
		},
		{
			name:      "validation key shifter zero",
			plainText: plainText,
			alphabet:  alphabet,
			key:       0,
			expected:  "key shifter must be greater than zero",
		},
		{
			name:      "validation key shifter equals filterResult length",
			plainText: plainText,
			alphabet:  "abcdefghijklmnopqrstuvwxyz12345678910!@#$%^&*(<>?:\"{}|abcdefghijklmnopqrstuvwxyz12345678910!@#$%^&*(<>?:\"{}|abcdefghijklmnopqrstuvwxyz12345678910!@#$%^&*(<>?:\"{}|",
			key:       53, // Adjust this key according to the length of the filterResult
			expected:  "invalid key shifter, try another key",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test encoding
			_, err := caesarCipher.NewCaesarCipher(tc.plainText, tc.alphabet, tc.key).Encrypt()
			if err != nil {
				if err.Error() != tc.expected {
					t.Errorf("expected error '%s', got '%v'", tc.expected, err)
				}
			}
		})
	}
}

func TestOutputNotEqualPlain(t *testing.T) {
	plainText, alphabet, key := generateRandom(1000, 141, 10)
	testCases := []struct {
		name            string
		plainText       string
		alphabet        string
		key             int
		invalidKey      int
		invalidAlphabet string
	}{
		{
			name:            "The decryption result is not the same as plain text by invalid alphabet",
			plainText:       plainText,
			alphabet:        alphabet,
			key:             key,
			invalidAlphabet: "今日は友達と一緒に公園に行き、たくさんの笑い声と楽しい時間を過ごしました。日差しがとても心地よく、風も爽やかでした。みんなでピクニックを楽しんで、美味しい料理を食べながらおしゃべりしました。カラフルな花々を見てリフレッシュでき、心温まる素晴らしい一日になりました。自然に囲まれて、",
		},
		{
			name:       "The decryption result is not the same as plain text by invalid key",
			plainText:  plainText,
			alphabet:   alphabet,
			key:        key,
			invalidKey: key + 1,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			enc, err := caesarCipher.NewCaesarCipher(tt.plainText, tt.alphabet, tt.key).Encrypt()
			if err != nil {
				t.Fatal(err)
			}

			var dec string
			if tt.invalidAlphabet != "" {
				dec, err = caesarCipher.NewCaesarCipher(enc, tt.invalidAlphabet, tt.key).Decrypt()
			} else {
				dec, err = caesarCipher.NewCaesarCipher(enc, tt.alphabet, tt.invalidKey).Decrypt()
			}

			if err != nil {
				t.Fatal(err)
			}

			if dec == tt.plainText {
				t.Errorf("Test failed: decryption result is the same as plain text. Got %s, expected not %s", dec, tt.plainText)
			}
		})
	}
}

func TestNewCaesarCipher(t *testing.T) {
	plainText, alphabet, key := generateRandom(1000, 141, 10)
	t1, err := caesarCipher.NewCaesarCipher(plainText, alphabet, key).Encrypt()
	if err != nil {
		t.Error(err)
	}
	t2, err := caesarCipher.NewCaesarCipher(t1, alphabet, key).Decrypt()
	if err != nil {
		t.Error(err)
	}
	if t2 != plainText {
		t.Errorf("got %s, want %s", t2, plainText)
	}
}
