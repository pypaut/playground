package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	maxJoltage := ComputePart1("input")
	fmt.Println(maxJoltage)
	return
}

func ComputePart1(path string) (maxJoltage int) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		maxJoltage += GetMaxJoltage(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}

func GetMaxJoltage(bank string) int {
	intBank := []int{}
	for _, c := range bank {
		intVal, err := strconv.Atoi(string(c))
		if err != nil {
			panic(err)
		}

		intBank = append(intBank, intVal)
	}

	// Find max diggit and its indice
	maxValue, maxIndex := MaxWithIndice(intBank[:len(intBank)-1])

	// Find max diggit, placed after this indice
	secondMaxValue, _ := MaxWithIndice(intBank[maxIndex+1:])

	return maxValue*10 + secondMaxValue
}

func MaxWithIndice(s []int) (value, index int) {
	if len(s) == 0 {
		return
	}

	value = s[0]
	for i, v := range s {
		if v > value {
			value = v
			index = i
		}
	}

	return
}
