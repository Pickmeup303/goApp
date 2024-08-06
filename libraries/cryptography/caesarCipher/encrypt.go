package caesarCipher

import (
	"errors"
	"strings"
)

const emmit = '\u0003' //End-of-text character

type elements struct {
	text       string
	alphabet   string
	keyShifter int
	seen       map[rune]bool
	results
}

type results struct {
	filterResult      []rune
	missingCharResult []rune
	tableOne          []rune
	tableTwo          []rune
	finalResult       string
}

func NewCaesarCipher(text, shifterAlphabet string, keyShifter int) *elements {
	return &elements{
		text:       text,
		alphabet:   shifterAlphabet,
		keyShifter: keyShifter,
		seen:       make(map[rune]bool),
		results:    results{},
	}
}

func (e *elements) validation() error {
	if e.text == "" {
		return errors.New("text is empty")
	}
	if len(e.alphabet) < 140 {
		return errors.New("alphabet must be at least 140 chars or more")
	}
	if e.keyShifter <= 0 {
		return errors.New("key shifter must be greater than zero")
	}
	if e.keyShifter == len(e.filterResult) {
		return errors.New("invalid key shifter, try another key")
	}
	return nil
}
func (e *elements) Encrypt() (string, error) {
	e.filter()
	if err := e.validation(); err != nil {
		return "", err
	}
	e.missingChar()
	e.shifter()
	e.substitution()
	return e.finalResult, nil
}

func (e *elements) filter() {
	for _, i2 := range e.alphabet {
		if !e.seen[i2] {
			e.seen[i2] = true
			e.filterResult = append(e.filterResult, i2)
			e.tableOne = append(e.tableOne, i2)
		}
	}
}

func (e *elements) missingChar() {
	for _, i2 := range e.filterResult {
		e.seen[i2] = true
	}

	for _, i2 := range e.text {
		if !e.seen[i2] {
			e.seen[i2] = true
			e.missingCharResult = append(e.missingCharResult, i2)
			e.tableOne = append(e.tableOne, i2)
		}
	}
}

func (e *elements) shifter() {
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

func (e *elements) modify(char rune) rune {
	for _, i2 := range e.tableOne {
		e.seen[i2] = true
	}
	for _, i2 := range e.tableTwo {
		e.seen[i2] = true
	}
	for _, v := range e.alphabet {
		e.seen[v] = true
	}

	for i := rune(32); i <= 126; i++ {
		if !e.seen[i] && i != char {
			return i
		}
	}
	return char
}

func (e *elements) substitution() {
	var result strings.Builder
	charMap := make(map[rune]rune)
	for i, r := range e.tableOne {
		charMap[r] = e.tableTwo[i]
	}

	for _, char := range e.text {
		if newChar, ok := charMap[char]; ok {
			result.WriteRune(newChar)
		} else {
			result.WriteRune(char)
		}
	}

	result.WriteRune(emmit)
	result.WriteString(string(e.missingCharResult))
	e.finalResult = result.String()
}
