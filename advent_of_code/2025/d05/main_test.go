package main

import (
	"reflect"
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestSample(t *testing.T) {
	ranges, ids := LoadInput("sample")

	expectedCount := 3
	gotCount := CountFreshIngredients(ids, ranges)
	assert.Equal(t, expectedCount, gotCount)

	expectedCount = 14
	gotCount = CountTotalFreshIDs(ranges)
	assert.Equal(t, expectedCount, gotCount)
}

func TestMergeRanges(t *testing.T) {
	cases := []struct {
		inRanges  []*Range
		outRanges []*Range
	}{
		{
			inRanges: []*Range{
				&Range{Low: 3, High: 5},
			},
			outRanges: []*Range{
				&Range{Low: 3, High: 5},
			},
		},
		{
			inRanges: []*Range{
				&Range{Low: 3, High: 5},
				&Range{Low: 10, High: 14},
			},
			outRanges: []*Range{
				&Range{Low: 3, High: 5},
				&Range{Low: 10, High: 14},
			},
		},
		{
			inRanges: []*Range{
				&Range{Low: 10, High: 14},
				&Range{Low: 12, High: 18},
			},
			outRanges: []*Range{
				&Range{Low: 10, High: 18},
			},
		},
		{
			inRanges: []*Range{
				&Range{Low: 10, High: 18},
				&Range{Low: 16, High: 20},
			},
			outRanges: []*Range{
				&Range{Low: 10, High: 20},
			},
		},
		{
			inRanges: []*Range{
				&Range{Low: 3, High: 5},
				&Range{Low: 10, High: 14},
				&Range{Low: 16, High: 20},
				&Range{Low: 12, High: 18},
			},
			outRanges: []*Range{
				&Range{Low: 3, High: 5},
				&Range{Low: 10, High: 20},
			},
		},
	}

	for _, c := range cases {
		result := MergeRanges(c.inRanges)
		if !reflect.DeepEqual(result, c.outRanges) {
			t.Fatalf("expected %v, got %v", c.outRanges, result)
		}
	}
}
