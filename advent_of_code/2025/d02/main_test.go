package main

import (
	"reflect"
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

func TestInvalidIds(t *testing.T) {
	cases := []struct {
		rng      string
		expected []int
	}{
		{
			rng:      "11-22",
			expected: []int{11, 22},
		},
		{
			rng:      "95-115",
			expected: []int{99, 111},
		},
		{
			rng:      "998-1012",
			expected: []int{999, 1010},
		},
		{
			rng:      "1188511880-1188511890",
			expected: []int{1188511885},
		},
	}

	for _, c := range cases {
		got := invalidIds(c.rng)
		if !reflect.DeepEqual(got, c.expected) {
			t.Fatalf("(rng: %s) got %v, expected %v", c.rng, got, c.expected)
		}
	}
}

func TestPatternRepeats(t *testing.T) {
	cases := []struct {
		id      string
		nbPats  int
		repeats bool
	}{
		{
			id:      "123123",
			nbPats:  2,
			repeats: true,
		},
		{
			id:      "11",
			nbPats:  2,
			repeats: true,
		},
		{
			id:      "1",
			nbPats:  1,
			repeats: false,
		},
		{
			id:      "1",
			nbPats:  2,
			repeats: false,
		},
		{
			id:      "1",
			nbPats:  3,
			repeats: false,
		},
	}

	for _, c := range cases {
		got := patternRepeats(c.id, c.nbPats)
		if got != c.repeats {
			t.Fatalf("(id: %s, nbPats: %d) got %t, expected %t", c.id, c.nbPats, got, c.repeats)
		}
	}
}
