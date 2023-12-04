package compute

import (
	"errors"
	"fmt"
	"strconv"
)

var literals = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var digits = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

func Compute(lines []string) (sum int, err error) {
	twoDigitsNumbers, err := computeTwoDigitStrings(lines)
	if err != nil {
		return 0, err
	}

	// Convert to int
	convertedNumbers, err := convertDigitsToInt(twoDigitsNumbers)
	if err != nil {
		return 0, err
	}

	// Compute sum
	for _, number := range convertedNumbers {
		sum += number
	}

	return sum, nil
}

func convertDigitsToInt(lines []string) (intLines []int, err error) {
	for _, digit := range lines {
		newInt, err := strconv.Atoi(digit)
		if err != nil {
			return []int{}, err
		}

		intLines = append(intLines, newInt)
	}

	return intLines, err
}

func computeTwoDigitStrings(lines []string) (twoDigits []string, err error) {
	for _, line := range lines {
		if line == "" {
			continue
		}

		// Find first digit
		firstDigit, err := findFirstDigit(line)
		if err != nil {
			return []string{}, err
		}

		// Find last digit
		lastDigit, err := findLastDigit(line)
		if err != nil {
			return []string{}, err
		}

		twoDigits = append(twoDigits, firstDigit+lastDigit)
	}

	return twoDigits, nil
}

func findFirstDigit(line string) (digit string, err error) {
	for indexLine := range line {
		// Check if literal digit
		for indexLiterals, lit := range literals {
			if indexLine+len(lit) <= len(line) && line[indexLine:indexLine+len(lit)] == lit {
				return digits[indexLiterals], nil
			}
		}

		// Check if digit
		if _, err := strconv.ParseInt(string(line[indexLine]), 10, 64); err == nil {
			return string(line[indexLine]), nil
		}
	}

	return "", errors.New(fmt.Sprintf("findFirstDigit: no digits found in %s", line))
}

func findLastDigit(line string) (digit string, err error) {
	for indexLine := len(line) - 1; indexLine >= 0; indexLine-- {

		// Check if literal digit
		for indexLiterals, lit := range literals {
			if indexLine+len(lit) <= len(line) && line[indexLine:indexLine+len(lit)] == lit {
				return digits[indexLiterals], nil
			}
		}

		// Check if digit
		if _, err := strconv.ParseInt(string(line[indexLine]), 10, 64); err == nil {
			return string(line[indexLine]), nil
		}
	}

	return "", errors.New(fmt.Sprintf("findLastDigit(): no digits found in %s", line))
}
