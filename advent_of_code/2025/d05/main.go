package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	Low, High int
}

func (r *Range) String() string {
	return fmt.Sprintf("[%d, %d]", r.Low, r.High)
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

func MergeRanges(inRanges []*Range) (outRanges []*Range)

// func MergeRangesRec(headRange *Range, inRanges []*Range) (outRanges []*Range) {
// 	if len(inRanges) == 0 {
// 		outRanges = append(outRanges, headRange)
// 		return
// 	}
//
// 	// if len(inRanges) == 1 {
// 	// 	r := inRanges[0]
//
// 	// 	if headRange.IsIncluded(r) {
// 	// 		outRanges = append(outRanges, headRange)
// 	// 		return
// 	// 	}
//
// 	// 	if headRange.LeftOverlaps(r) {
// 	// 		newRange := &Range{Low: headRange.Low, High: r.High}
// 	// 		outRanges = append(outRanges, newRange)
// 	// 		return
// 	// 	}
//
// 	// 	if r.LeftOverlaps(headRange) {
// 	// 		newRange := &Range{Low: r.Low, High: headRange.High}
// 	// 		outRanges = append(outRanges, newRange)
// 	// 		return
// 	// 	}
//
// 	// 	outRanges = append(outRanges, headRange, r)
// 	// 	return
// 	// }
//
// 	// len(headRange) > 1
// 	for i, r := range inRanges {
// 		if headRange.IsIncluded(r) {
// 			return MergeRangesRec(inRanges[0], inRanges[1:])
// 		}
//
// 		if r.IsIncluded(headRange) {
// 			inRanges = slices.Delete(inRanges, i, i+1)
// 			return MergeRangesRec(headRange, inRanges)
// 		}
//
// 		if headRange.LeftOverlaps(r) {
// 			newRange := &Range{Low: headRange.Low, High: r.High}
// 			inRanges = slices.Delete(inRanges, i, i+1)
// 			return MergeRangesRec(newRange, inRanges)
// 		}
//
// 		if r.LeftOverlaps(headRange) {
// 			newRange := &Range{Low: r.Low, High: headRange.High}
// 			inRanges = slices.Delete(inRanges, i, i+1)
// 			return MergeRangesRec(newRange, inRanges)
// 		}
// 	}
//
// 	// headRange doesn't fit anywhere, keep it
// 	inRanges = append(inRanges, headRange)
// 	return MergeRangesRec(inRanges[0], inRanges[1:])
// }
//
// func MergeRanges(inRanges []*Range) (outRanges []*Range) {
// 	if len(inRanges) <= 1 {
// 		return inRanges
// 	}
//
// 	return MergeRangesRec(inRanges[0], inRanges[1:])
// }

// func MergeRanges2(inRanges []*Range) (outRanges []*Range) {
// 	for i, r1 := range inRanges {
// 		for j, r2 := range inRanges {
// 			if i == j {
// 				continue
// 			}
//
// 		}
// 	}
// 	return
// }
//
// func MergeRanges(inRanges []*Range) (outRanges []*Range) {
// 	for len(inRanges) > 0 {
// 		r1 := inRanges[0]
// 		inRanges = slices.Delete(inRanges, 0, 1)
//
// 		// wasAdded := false
// 		toAdd := true
// 		for _, r2 := range inRanges {
// 			if r1.IsIncluded(r2) {
// 				toAdd = false
// 				break
// 			}
//
// 			if r1.LeftOverlaps(r2) {
// 				newRange := &Range{Low: r1.Low, High: r2.High}
// 				outRanges = append(outRanges, newRange)
// 				// inRanges = slices.Delete(inRanges, i, i+1)
// 				toAdd = false
// 				break
// 			}
//
// 			if r2.LeftOverlaps(r1) {
// 				newRange := &Range{Low: r2.Low, High: r1.High}
// 				outRanges = append(outRanges, newRange)
// 				// inRanges = slices.Delete(inRanges, i, i+1)
// 				toAdd = false
// 				break
// 			}
// 		}
//
// 		if toAdd {
// 			outRanges = append(outRanges, r1)
// 		}
// 	}
//
// 	return
// }

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
