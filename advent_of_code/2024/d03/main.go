package main

import (
	"d03/internal/d03"
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	sum := d03.ScanMemory(string(data))
	fmt.Printf("Total: %d\n", sum)

	return
}
