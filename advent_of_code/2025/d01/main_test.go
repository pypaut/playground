package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeNextPos(t *testing.T) {
	cases := []struct {
		currentPos             int
		clickDir               ClickDir
		nbClicks               int
		expectedNewPos         int
		expectedPassagesByZero int
	}{
		{
			0, "L", 1, 99, 0,
		},
		{
			0, "R", 1, 1, 0,
		},
		{
			99, "R", 1, 0, 1,
		},
		{
			50, "R", 600, 50, 6,
		},
		{
			50, "L", 600, 50, 6,
		},
		{
			50, "L", 644, 6, 6,
		},
		{
			50, "L", 594, 56, 6,
		},
		{
			50, "L", 68, 82, 1,
		},
		{
			82, "L", 30, 52, 0,
		},
		{
			52, "R", 48, 0, 1,
		},
		{
			0, "L", 5, 95, 0,
		},
		{
			95, "R", 60, 55, 1,
		},
		{
			55, "L", 55, 0, 1,
		},
		{
			55, "L", 1, 54, 0,
		},
		{
			54, "L", 99, 55, 1,
		},
		{
			55, "R", 14, 69, 0,
		},
		{
			69, "L", 82, 87, 1,
		},
	}

	for _, c := range cases {
		nextPos, nbPassages := ComputeNextPos(c.currentPos, c.clickDir, c.nbClicks)
		if nextPos != c.expectedNewPos {
			t.Fatalf("(%d,%s%d) new pos: expected %d, got %d",
				c.currentPos, c.clickDir, c.nbClicks, c.expectedNewPos, nextPos,
			)
		} else if nbPassages != c.expectedPassagesByZero {
			t.Fatalf("(%d,%s%d) nb passages: expected %d, got %d",
				c.currentPos, c.clickDir, c.nbClicks, c.expectedPassagesByZero, nbPassages,
			)
		}
	}
}

func TestSampleInput(t *testing.T) {
	landedOnZero, passedByZero := Compute("sample")
	assert.Equal(t, 3, landedOnZero)
	assert.Equal(t, 6, passedByZero)
}
