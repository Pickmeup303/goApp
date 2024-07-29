package caesarCipher

import (
	"strings"
)

func (e *elements) Decode() string {
	e.filter()
	e.decryptMissingChar()
	e.shifter()
	e.decryptSubstitution()
	return e.finalResult
}

func (e *elements) decryptMissingChar() {
	arr := strings.Split(e.text, string(emmit))
	secretEmmit := []rune(arr[1])
	runeText := []rune(arr[0])
	e.tableOne = append(e.tableOne, secretEmmit...)
	e.text = string(runeText)
}

func (e *elements) decryptSubstitution() {
	var result strings.Builder

	charSet := make(map[rune]rune)
	for i, v := range e.tableTwo {
		charSet[v] = e.tableOne[i]
	}

	for _, v := range e.text {
		if newChar, ok := charSet[v]; ok {
			result.WriteRune(newChar)
		} else {
			result.WriteRune(v)
		}
	}
	e.finalResult = result.String()
}
