package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	maxJoltage := ComputePart1("input")
	fmt.Println(maxJoltage)

	maxJoltage = ComputePart2("input")
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

func ComputePart2(path string) (maxJoltage int) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		maxJoltage += GetMaxJoltagePart2(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}

func GetMaxJoltagePart2(bank string) (maxJoltage int) {
	intBank := []int{}
	for _, c := range bank {
		intVal, err := strconv.Atoi(string(c))
		if err != nil {
			panic(err)
		}

		intBank = append(intBank, intVal)
	}

	l := len(bank)
	i := -1
	m := intBank[0]

	for counter := 11; counter >= 0; counter-- {
		left := i + 1
		right := l - counter
		m, i = MaxWithIndice(intBank, left, right)
		maxJoltage += m * int(math.Pow10(counter))
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
	maxValue, maxIndex := MaxWithIndice(intBank, 0, len(intBank)-1)

	// Find max diggit, placed after this indice
	secondMaxValue, _ := MaxWithIndice(intBank, maxIndex+1, len(intBank))

	return maxValue*10 + secondMaxValue
}

func MaxWithIndice(s []int, l int, r int) (value, index int) {
	if len(s) == 0 || l >= r || l >= len(s) || r > len(s) {
		panic(errors.New("wrong max call"))
	}

	index = l
	value = s[index]
	for i := l; i < r; i++ {
		if s[i] > value {
			value = s[i]
			index = i
		}
	}

	return value, index
}
