package transpose

import "strings"

// kriptografi transposisi kolom
// karna pembacaanya berdasarkan kolom yang disepakati

type transpose struct {
	text    string
	col     int
	padding rune
}

const Padding = '*'

func NewTranspose(text string, colWidth int) *transpose {
	return &transpose{
		text:    text,
		col:     colWidth,
		padding: Padding,
	}
}

func (t *transpose) Encode() string {
	textRune := []rune(t.text)
	rowCount := (len(textRune) / t.col) + 1

	matrix := make([][]*rune, rowCount)
	for i := range matrix {
		matrix[i] = make([]*rune, t.col)
	}

	for i := 0; i < rowCount*t.col; i++ {
		row := i / t.col
		col := i % t.col
		if i < len(textRune) {
			matrix[row][col] = &textRune[i]
		} else {
			matrix[row][col] = &t.padding
		}
	}

	var result strings.Builder
	for col := t.col - 1; col >= 0; col-- {
		for row := rowCount - 1; row >= 0; row-- {
			result.WriteRune(*matrix[row][col])
		}
	}
	return result.String()
}
