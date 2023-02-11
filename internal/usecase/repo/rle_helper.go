package repo

import (
	"strconv"
	"strings"
	"unicode"
)

func encode(input string, index int, c chan RleChan) {
	var result strings.Builder
	for len(input) > 0 {
		firstLetter := input[0]
		inputLength := len(input)
		input = strings.TrimLeft(input, string(firstLetter))
		if counter := inputLength - len(input); counter > 1 {
			result.WriteString(strconv.Itoa(counter))
		}
		result.WriteString(string(firstLetter))
	}
	c <- RleChan{index: index, value: result.String()}
}
func decode(input string, index int, c chan RleChan) {
	var result strings.Builder
	for len(input) > 0 {
		letterIndex := strings.IndexFunc(input, func(r rune) bool { return !unicode.IsDigit(r) })
		multiply := 1
		if letterIndex != 0 {
			multiply, _ = strconv.Atoi(input[:letterIndex])
		}
		result.WriteString(strings.Repeat(string(input[letterIndex]), multiply))
		input = input[letterIndex+1:]
	}
	c <- RleChan{index: index, value: result.String()}
}
