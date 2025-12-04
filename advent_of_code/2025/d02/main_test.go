package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputePart1(t *testing.T) {
	res := ComputePart1("sample")
	expected := 1227775554

	assert.Equal(t, expected, res)
}

func TestComputePart2(t *testing.T) {
	res := ComputePart2("sample")
	expected := 4174379265

	assert.Equal(t, expected, res)
}
