package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Range struct {
	Low, High int
}

func main() {
	ranges, ids := LoadInput("input")

	count := CountFreshIngredients(ids, ranges)
	fmt.Printf("Counted %d fresh ingredients\n", count)

	count = CountTotalFreshIDs(ranges)
	fmt.Printf("%d total fresh ingredients\n", count)

	fmt.Println("Done")

	return
}

func CountFreshIngredients(ids []int, ranges []*Range) (count int) {
	for _, id := range ids {
		if IDFitsAnyRange(id, ranges) {
			count++
		}
	}

	return
}

func CountTotalFreshIDs(ranges []*Range) (count int) {
	// countedIds := map[int]int{}
	// var countedIds []int
	ranges = MergeRanges(ranges)
	for _, r := range ranges {
		count += r.High - r.Low + 1
	}

	return count
	// return len(countedIds)
}

func MergeRanges(inRanges []*Range) (outRanges []*Range) {
	for len(inRanges) > 0 {
		r1 := inRanges[0]
		inRanges = slices.Delete(inRanges, 0, 1)

		wasAdded := false
		for i, r2 := range inRanges {
			if r1.IsIncluded(r2) {
				break
			}

			if r1.LeftOverlaps(r2) {
				newRange := &Range{Low: r1.Low, High: r2.High}
				outRanges = append(outRanges, newRange)
				inRanges = slices.Delete(inRanges, i, i+1)
				wasAdded = true
				break
			}

			if r2.LeftOverlaps(r1) {
				newRange := &Range{Low: r2.Low, High: r1.High}
				outRanges = append(outRanges, newRange)
				inRanges = slices.Delete(inRanges, i, i+1)
				wasAdded = true
				break
			}
		}

		if !wasAdded {
			outRanges = append(outRanges, r1)
		}
	}

	return
}

func (r1 *Range) IsIncluded(r2 *Range) bool {
	return r2.Low <= r1.Low && r1.High <= r2.High
}

func (r1 *Range) LeftOverlaps(r2 *Range) bool {
	return r1.Low <= r2.Low &&
		r2.Low <= r1.High &&
		r1.High <= r2.High
}

func IDFitsAnyRange(id int, ranges []*Range) bool {
	for _, r := range ranges {
		if r.Low <= id && id <= r.High {
			return true
		}
	}

	return false
}

func LoadInput(path string) (ranges []*Range, ids []int) {
	b, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	strInput := strings.Split(string(b), "\n\n")
	strRanges := strings.Split(strInput[0], "\n")
	strIds := strings.Split(strInput[1], "\n")

	for _, r := range strRanges {
		if r == "" {
			continue
		}

		strValues := strings.Split(r, "-")
		low, err := strconv.Atoi(strValues[0])
		if err != nil {
			panic(err)
		}

		high, err := strconv.Atoi(strValues[1])
		if err != nil {
			panic(err)
		}

		ranges = append(ranges, &Range{Low: low, High: high})
	}

	for _, strId := range strIds {
		if strId == "" {
			continue
		}

		id, err := strconv.Atoi(strId)
		if err != nil {
			panic(err)
		}

		ids = append(ids, id)
	}

	return
}
