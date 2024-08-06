package transpose

import "strings"

func (t *transpose) Decrypt() string {
	textRune := []rune(t.text)
	rowCount := (len(textRune) + t.col - 1) / t.col

	matrix := make([][]*rune, rowCount)
	for i := range matrix {
		matrix[i] = make([]*rune, t.col)
	}

	index := 0
	for col := t.col - 1; col >= 0; col-- {
		for row := rowCount - 1; row >= 0; row-- {
			matrix[row][col] = &textRune[index]
			index++
		}
	}

	var result strings.Builder
	for row := 0; row < rowCount; row++ {
		for col := 0; col < t.col; col++ {
			if *matrix[row][col] != t.padding {
				result.WriteRune(*matrix[row][col])
			}
		}
	}
	return result.String()
}
