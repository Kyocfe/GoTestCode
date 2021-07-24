package main

import (
	"errors"
	"strings"
)

func isIncludedSpaces(value string) bool {
	return value == " "
}

func stringLengthZero(value string) bool {
	return len(value) == 0
}

func stringToUpper(input string) (string, error) {
	if stringLengthZero(input) {
		return "", errors.New("string input length must longer than 0")
	}

	runes := []rune(input)

	for _, char := range runes {
		if isIncludedSpaces(string(char)) {
			return "", errors.New("string input cannot accept some spaces")
		}
	}

	return strings.ToUpper(input), nil
}
