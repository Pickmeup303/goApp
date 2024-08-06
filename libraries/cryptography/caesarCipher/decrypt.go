package caesarCipher

import (
	"strings"
)

func (e *elements) Decrypt() (string, error) {
	e.filter()
	if err := e.validation(); err != nil {
		return "", err
	}
	e.decryptMissingChar()
	e.decryptShifter()
	e.decryptSubstitution()
	return e.finalResult, nil
}

func (e *elements) decryptMissingChar() {
	arr := strings.Split(e.text, string(emmit))
	secretEmmit := []rune(arr[1])
	runeText := []rune(arr[0])
	e.tableOne = append(e.tableOne, secretEmmit...)
	e.text = string(runeText)
}

func (e *elements) decryptShifter() {
	mapChar := make(map[rune]rune)
	for i, r := range e.filterResult {
		index := (i + e.keyShifter) % len(e.filterResult)
		mapChar[r] = e.filterResult[index]
	}

	for _, r := range e.tableOne {
		if newChar, ok := mapChar[r]; ok {
			e.tableTwo = append(e.tableTwo, newChar)
		} else {
			e.tableTwo = append(e.tableTwo, e.modify(r))
		}
	}
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
