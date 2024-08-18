package transpose

import "strings"

// kriptografi transposisi kolom
// karna pembacaanya berdasarkan kolom yang disepakati

type transpose struct {
	text    string
	col     int
	key     int
	padding rune
}

const Padding = '\u00D7' //Unit Separator

func NewTranspose(text string, colWidth, key int) *transpose {
	return &transpose{
		text:    text,
		col:     colWidth,
		key:     key,
		padding: Padding,
	}
}

// Encrypt is the transposition function that encodes the text by rearranging characters based on the key
func (t *transpose) Encrypt() string {
	textRune := []rune(t.text)              // Convert the text to a slice of runes (supports Unicode characters)
	rowCount := (len(textRune) / t.col) + 1 // Calculate the number of rows needed for the matrix
	matrix := make([][]*rune, rowCount)     // Create a 2D slice (matrix) to store the characters

	// Initialize each row in the matrix
	for i := range matrix {
		matrix[i] = make([]*rune, t.col)
	}

	// Fill the matrix with characters from text, based on the key
	for i := 0; i < rowCount*t.col; i++ {
		row := i / t.col           // Determine the row index
		col := (t.key + i) % t.col // Determine the column index based on the key
		if i < len(textRune) {
			matrix[row][col] = &textRune[i] // Place the character in the calculated position
		} else {
			matrix[row][col] = &t.padding // Fill remaining empty spaces with 'padding'
		}
	}

	// Read the matrix column by column, starting from the last column, to create the encoded text
	var result strings.Builder
	for col := t.col - 1; col >= 0; col-- {
		for row := rowCount - 1; row >= 0; row-- {
			result.WriteRune(*matrix[row][col])
		}
	}
	return result.String() // Return the encoded text
}
