package main

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestSample(t *testing.T) {
	ranges, ids := LoadInput("sample")
	expectedCount := 3
	gotCount := CountFreshIngredients(ids, ranges)
	assert.Equal(t, expectedCount, gotCount)
}
