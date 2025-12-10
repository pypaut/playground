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

	expectedJoltage = 3121910778619
	totalJoltage = ComputePart2("sample")
	assert.Equal(t, totalJoltage, expectedJoltage)
}

func TestGetMaxJoltagePart2(t *testing.T) {
	cases := []struct {
		bank       string
		maxJoltage int
	}{
		{
			bank:       "987654321111111",
			maxJoltage: 987654321111,
		},
		{
			bank:       "811111111111119",
			maxJoltage: 811111111119,
		},
		{
			bank:       "234234234234278",
			maxJoltage: 434234234278,
		},
		{
			bank:       "818181911112111",
			maxJoltage: 888911112111,
		},
	}

	for _, c := range cases {
		joltage := GetMaxJoltagePart2(c.bank)
		assert.Equal(t, joltage, c.maxJoltage)
	}
}

func TestMaxWithIndice(t *testing.T) {
	cases := []struct {
		s              []int
		left           int
		right          int
		expectedMax    int
		expectedIndice int
	}{
		{
			s:              []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			left:           0,
			right:          10,
			expectedMax:    9,
			expectedIndice: 9,
		},
		{
			s:              []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			left:           9,
			right:          10,
			expectedMax:    9,
			expectedIndice: 9,
		},
		{
			s:              []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			left:           5,
			right:          6,
			expectedMax:    5,
			expectedIndice: 5,
		},
		{
			s:              []int{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8},
			left:           0,
			right:          4,
			expectedMax:    4,
			expectedIndice: 2,
		},
		{
			s:              []int{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8},
			left:           3,
			right:          5,
			expectedMax:    3,
			expectedIndice: 4,
		},
	}

	for _, c := range cases {
		m, i := MaxWithIndice(c.s, c.left, c.right)
		assert.Equal(t, m, c.expectedMax)
		assert.Equal(t, i, c.expectedIndice)
	}
}
