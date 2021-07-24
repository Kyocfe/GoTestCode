package main

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

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

func TestStringToUpper(t *testing.T) {
	for _, testCase := range parameters {
		err, res := stringToUpper(testCase.in)
		assert.Equal(t, testCase.err, err)
		assert.Equal(t, testCase.out, res)
	}
}
