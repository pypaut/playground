package compute

import (
	"errors"
	"fmt"
	"strconv"
)

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
	for _, character := range line {
		if _, err := strconv.ParseInt(string(character), 10, 64); err == nil {
			return string(character), nil
		}
	}

	return "", errors.New(fmt.Sprintf("findFirstDigit: no digits found in %s", line))
}

func findLastDigit(line string) (digit string, err error) {
	for index := len(line) - 1; index >= 0; index-- {
		character := line[index]
		if _, err := strconv.ParseInt(string(character), 10, 64); err == nil {
			return string(character), nil
		}
	}

	return "", errors.New(fmt.Sprintf("findLastDigit(): no digits found in %s", line))
}
