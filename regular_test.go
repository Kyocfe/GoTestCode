package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringToUpper_IncludedSpaces(t *testing.T) {
	res, err := stringToUpper("testsuc cessful")
	assert.Equal(t, err, errors.New("string input cannot accept some spaces"))
	assert.Equal(t, res, "")
}

func TestStringToUpper_StringLengthZero(t *testing.T) {
	res, err := stringToUpper("")
	assert.Equal(t, err, errors.New("string input length must longer than 0"))
	assert.Equal(t, res, "")
}

func TestStringToUpper_Success(t *testing.T) {
	res, err := stringToUpper("testsuccessful")
	assert.Equal(t, err, nil)
	assert.Equal(t, res, "TESTSUCCESSFUL")
}
