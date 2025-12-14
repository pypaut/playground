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

func main() {
	ranges, ids := LoadInput("input")
	count := CountFreshIngredients(ids, ranges)

	fmt.Printf("Counted %d fresh ingredients\n", count)

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
