package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSamplePart1(t *testing.T) {
	m := LoadMatrix("sample")
	expected := 13
	result := CountAccessibleRolls(m)
	assert.Equal(t, expected, result)
}

func TestSamplePart2(t *testing.T) {
	m := LoadMatrix("sample")
	expected := 43
	result := RemoveRolls(m)
	assert.Equal(t, expected, result)
}
