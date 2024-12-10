package d03

import (
	"regexp"
	"strconv"
	"strings"
)

func ScanMemoryWithDos(input string) (sum int) {
	enabledParts := strings.Split(input, `do()`)
	for _, enabledP := range enabledParts {
		withoutDont := strings.Split(enabledP, `don't()`)[0]
		instructions := ExtractMulInstructions(withoutDont)
		for _, instruction := range instructions {
			sum += ExecuteInstruction(instruction)
		}
	}

	return
}

func ScanMemory(input string) (sum int) {
	instructions := ExtractMulInstructions(input)
	for _, instruction := range instructions {
		sum += ExecuteInstruction(instruction)
	}

	return
}

func ExtractMulInstructions(input string) []string {
	// Extract all correct "mul(x,y)"
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	return re.FindAllString(input, -1)
}

func ExecuteInstruction(instruction string) int {
	re := regexp.MustCompile(`\d{1,3}`)
	strNumbers := re.FindAllString(instruction, -1)
	x, err := strconv.Atoi(strNumbers[0])
	if err != nil {
		panic(err)
	}

	y, err := strconv.Atoi(strNumbers[1])
	if err != nil {
		panic(err)
	}

	return x * y
}
