package compute

import "testing"

func TestFindFirstDigit(t *testing.T) {
	lines := []string{
		"kld3lasdk93",
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
		"10394203",
	}
	expected := []string{
		"3",
		"2",
		"8",
		"1",
		"2",
		"4",
		"1",
		"7",
		"1",
	}

	for i, line := range lines {
		result, err := findFirstDigit(line)
		if err != nil {
			t.Errorf("error during findFirstDigit(): %v", err)
		}

		if result != expected[i] {
			t.Errorf("findFirstDigit(%s) = %s; want %s", line, result, expected[i])
		}
	}
}

func TestFindLastDigit(t *testing.T) {
	lines := []string{
		"kld3lasdk94",
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
		"10394203",
	}
	expected := []string{
		"4",
		"9",
		"3",
		"3",
		"4",
		"2",
		"4",
		"6",
		"3",
	}

	for i, line := range lines {
		result, err := findLastDigit(line)
		if err != nil {
			t.Errorf("error during findLastDigit(): %v", err)
		}

		if result != expected[i] {
			t.Errorf("findLastDigit(%s) = %s; want %s", line, result, expected[i])
		}
	}
}

func TestComputeTwoDigitStrings(t *testing.T) {
	lines := []string{
		"kld3lasdk94",
		"10394203",
		"oifjoiwe3092358kj2jih35kljkjkjdhksehf3kjdvksdujnfksundfgkdnsifug9",
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}

	result, err := computeTwoDigitStrings(lines)
	if err != nil {
		t.Errorf("error during computeTwoDigitStrings(): %s", err)
	}

	expected := []string{
		"34",
		"13",
		"39",
		"29",
		"83",
		"13",
		"24",
		"42",
		"14",
		"76",
	}

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("computeTwoDigitStrings(%v) = %v; want %v", lines, result, expected)
		}
	}
}

func TestConvertDigitsToInt(t *testing.T) {
	twoDigitNumbers := []string{
		"34",
		"13",
		"39",
	}

	result, err := convertDigitsToInt(twoDigitNumbers)
	if err != nil {
		t.Errorf("error during computeTwoDigitStrings(): %s", err)
	}

	expected := []int{
		34,
		13,
		39,
	}

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("convertDigitsToInt(%v) = %v; want %v", twoDigitNumbers, result, expected)
		}
	}
}
