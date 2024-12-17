package main

import (
	"d04/internal/d04"
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")
	var fixedLines []string

	for _, line := range lines {
		if line != "" {
			fixedLines = append(fixedLines, line)
		}
	}

	fmt.Printf("Number of occurrences of XMAS: %d\n", d04.CountXmas(fixedLines))
	return
}
