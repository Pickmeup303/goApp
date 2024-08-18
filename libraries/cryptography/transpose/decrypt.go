package transpose

import "strings"

// Decrypt is the transposition function that encodes the text by rearranging characters based on the key
func (t *transpose) Decrypt() string {
	textRune := []rune(t.text)                      // Convert the encoded text to a slice of runes
	rowCount := (len(textRune) + t.col - 1) / t.col // Calculate the number of rows

	matrix := make([][]*rune, rowCount) // Create a 2D slice (matrix) to store the characters
	for i := range matrix {
		matrix[i] = make([]*rune, t.col)
	}

	// Fill the matrix by reading the encoded text in the reverse order it was written
	index := 0
	for col := t.col - 1; col >= 0; col-- {
		for row := rowCount - 1; row >= 0; row-- {
			matrix[row][col] = &textRune[index] // Place the character in the matrix
			index++
		}
	}

	// Read the matrix in the original order to reconstruct the original text
	var result strings.Builder
	for i := range rowCount * t.col {
		row := i / t.col                    // Determine the row index
		col := (t.key + i) % t.col          // Determine the column index based on the key
		if *matrix[row][col] != t.padding { // Ignore padding characters
			result.WriteRune(*matrix[row][col])
		}
	}

	return result.String() // Return the decoded text
}
