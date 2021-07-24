package main

import (
	"errors"
	"strings"
)

func stringToUpper(input string) (error, string) {
	if len(input) == 0 {
		return errors.New("string input length must longer than 0"), ""
	}

	runes := []rune(input)

	for _, value := range runes {
		if string(value) == " " {
			return errors.New("string input cannot accept some spaces"), ""
		}
	}

	return nil, strings.ToUpper(input)
}
