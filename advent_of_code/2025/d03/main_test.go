package main

import (
	"testing"

	"github.com/go-jose/go-jose/v4/testutils/assert"
)

func TestGetMaxJoltage(t *testing.T) {
	cases := []struct {
		bank       string
		maxJoltage int
	}{
		{
			bank:       "987654321111111",
			maxJoltage: 98,
		},
		{
			bank:       "811111111111119",
			maxJoltage: 89,
		},
		{
			bank:       "234234234234278",
			maxJoltage: 78,
		},
		{
			bank:       "818181911112111",
			maxJoltage: 92,
		},
	}

	for _, c := range cases {
		joltage := GetMaxJoltage(c.bank)
		assert.Equal(t, joltage, c.maxJoltage)
	}
}

func TestSample(t *testing.T) {
	expectedJoltage := 357
	totalJoltage := ComputePart1("sample")
	assert.Equal(t, totalJoltage, expectedJoltage)
}
