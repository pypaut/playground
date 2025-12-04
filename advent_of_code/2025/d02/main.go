package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	count := ComputePart1("input")
	fmt.Println(count)

	count = ComputePart2("input")
	fmt.Println(count)
}

func ComputePart1(path string) (count int) {
	// Open file
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	strInput := string(bytes)
	strInput = strings.Trim(strInput, "\n")

	// Iterate on each range
	for r := range strings.SplitSeq(strInput, ",") {
		limitLeft, limitRight := getLimits(r)

		// Check each ID in the current range
		for id := limitLeft; id <= limitRight; id++ {
			idStr := strconv.Itoa(id)
			idLen := len(idStr)
			if idLen%2 != 0 {
				continue
			}

			idLeft := idStr[:idLen/2]
			idRight := idStr[idLen/2:]

			if idLeft == idRight {
				count += id
			}
		}
	}

	return count
}

func ComputePart2(path string) (count int) {
	// Open file
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	strInput := string(bytes)
	strInput = strings.Trim(strInput, "\n")

	// Iterate on each range
	for r := range strings.SplitSeq(strInput, ",") {
		limitLeft, limitRight := getLimits(r)

		// Check each ID in the current range
		for id := limitLeft; id <= limitRight; id++ {
			count += computeSingleId(id)
		}
	}

	return count
}

func computeSingleId(id int) (count int) {
	idStr := strconv.Itoa(id)
	idLen := len(idStr)

	// Check who can divide idLen
	for i := 2; i <= idLen; i++ {
		if idLen%i != 0 {
			continue
		}

		if patternRepeats(idStr, i) {
			count += id
			continue
		}
	}

	return count
}

func patternRepeats(id string, nbPats int) (repeats bool) {
	patLen := len(id) / nbPats
	patterns := []string{}
	for k := range nbPats {
		patterns = append(patterns, id[k*patLen:(k+1)*patLen])
	}

	for _, p := range patterns {
		if p != patterns[0] {
			return false
		}
	}

	return true
}

func getLimits(rng string) (leftLimit, rightLimit int) {
	limits := strings.Split(rng, "-")
	if len(limits) != 2 {
		panic(errors.New("should only have 2 limits"))
	}

	leftLimit, err := strconv.Atoi(limits[0])
	if err != nil {
		panic(err)
	}

	rightLimit, err = strconv.Atoi(limits[1])
	if err != nil {
		panic(err)
	}

	return
}
