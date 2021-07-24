package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var parameters = []struct {
	in  string
	out string
	err error
}{
	{"testsuccessful", "TESTSUCCESSFUL", nil},
	{"TESTSUCCESSFUL", "TESTSUCCESSFUL", nil},
	{"TEStsuccessFUL", "TESTSUCCESSFUL", nil},
	{"test successful", "", errors.New("string input cannot accept some spaces")},
	{"", "", errors.New("string input length must longer than 0")},
}

func TestStringToUpperByParameters(t *testing.T) {
	for _, testCase := range parameters {
		res, err := stringToUpper(testCase.in)
		assert.Equal(t, testCase.err, err)
		assert.Equal(t, testCase.out, res)
	}
}
