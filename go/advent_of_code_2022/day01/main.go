package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file := LoadFile("input")
	defer file.Close()

	elves := []int{0}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			elves = append(elves, 0)
			continue
		}

		cal, err := strconv.Atoi(line)
		Check(err)

		elves[len(elves)-1] += cal
	}

	err := scanner.Err()
	Check(err)

	// Solutions
	max, _ := Max(elves)
	fmt.Printf("Most calories carried by a single elf: %d\n", max)
	fmt.Printf(
		"Total calories carried by the three elves carrying the most calories: %d\n",
		SumOfThreeMaxes(elves),
	)
	return
}

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func LoadFile(filename string) *os.File {
	file, err := os.Open(filename)
	Check(err)

	return file
}

func Max(slice []int) (int, int) {
	if len(slice) < 1 {
		return -1, -1
	}

	index := 0
	max := slice[0]

	for i, e := range slice {
		if e > max {
			max = e
			index = i
		}
	}

	return max, index
}

func SumOfThreeMaxes(slice []int) int {
	if len(slice) < 1 {
		return -1
	}

	sliceCopy := make([]int, len(slice))
	copy(sliceCopy, slice)

	max1, index := Max(sliceCopy)
	sliceCopy[index] = 0

	max2, index := Max(sliceCopy)
	sliceCopy[index] = 0

	max3, _ := Max(sliceCopy)

	return max1 + max2 + max3
}
