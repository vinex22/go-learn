package main

import (
	"strings"
)

func ToMorseCode(input string) string {
	morseCode := []string{
		".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---",
		"-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-",
		"..-", "...-", ".--", "-..-", "-.--", "--..", " ",
	}

	alphabet := []string{
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
		"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", " ",
	}

	input = strings.ToUpper(input)
	output := ""

	for _, char := range input {
		index := -1
		for i, letter := range alphabet {
			if string(char) == letter {
				index = i
				break
			}
		}
		if index == -1 {
			continue
		}
		output += morseCode[index]
	}

	return output
}
