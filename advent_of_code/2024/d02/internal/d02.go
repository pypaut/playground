package d02

import (
	"math"
	"os"
	"strconv"
	"strings"
)

func ParseReports(inputFile string) (reports [][]int, err error) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		var intLevels []int
		strLevels := strings.Split(line, " ")

		for _, strLevel := range strLevels {
			if strLevel == "" {
				continue
			}

			intValue, err := strconv.Atoi(strLevel)
			if err != nil {
				return nil, err
			}

			intLevels = append(intLevels, intValue)
		}

		// Avoid adding empty levels
		if len(intLevels) > 0 {
			reports = append(reports, intLevels)
		}
	}

	return
}

func IsReportSafe(report []int) bool {
	if !IsStrictlyMonotonous(report) {
		return false
	}

	// Check levels differ by at most three
	if !DiffersByMaxThree(report) {
		return false
	}

	return true
}

func BuildSubReport(report []int, indexToRemove int) []int {
	subReport1 := make([]int, indexToRemove)
	_ = copy(subReport1, report[:indexToRemove])

	subReport2 := make([]int, len(report)-indexToRemove-1)
	_ = copy(subReport2, report[indexToRemove+1:])

	return append(subReport1, subReport2...)
}

func IsStrictlyMonotonousWithDampener(report []int) (isMonotonous bool, unsafeLevelIndices []int) {
	if IsStrictlyMonotonous(report) {
		return true, nil
	}

	i := 0

	for i < len(report) {
		subReport := BuildSubReport(report, i)

		if IsStrictlyMonotonous(subReport) {
			unsafeLevelIndices = append(unsafeLevelIndices, i)
		}

		i++
	}

	if len(unsafeLevelIndices) > 0 {
		return true, unsafeLevelIndices
	}

	return false, nil
}

func IsReportSafeWithDampener(report []int) (bool, int) {
	isMonotonous, unsafeLevelIndex := IsStrictlyMonotonousWithDampener(report)
	if !isMonotonous {
		return false, unsafeLevelIndex
	}

	if unsafeLevelIndex != -1 {
		subReport := BuildSubReport(report, unsafeLevelIndex)
		return DiffersByMaxThree(subReport), unsafeLevelIndex
	}

	return DiffersByMaxThree(report), unsafeLevelIndex
}

func DiffersByMaxThree(report []int) bool {
	for i, _ := range report {
		if i == len(report)-1 {
			// The whole report was checked
			break
		}

		if math.Abs(float64(report[i]-report[i+1])) > 3 {
			return false
		}
	}

	return true
}

func IsStrictlyMonotonous(report []int) bool {
	coef := 0

	for i, _ := range report {
		if i == len(report)-1 {
			// The whole report was checked
			break
		}

		if report[i] < report[i+1] {
			if coef == -1 {
				return false
			}

			coef = 1
		} else if report[i] > report[i+1] {
			if coef == 1 {
				return false
			}

			coef = -1
		} else {
			// Two levels are equals
			return false
		}
	}

	return true
}

func NumberOfSafeReports(reports [][]int) (number int) {
	for _, report := range reports {
		if IsReportSafe(report) {
			number++
		}
	}

	return
}

func NumberOfSafeReportsWithDampener(reports [][]int) (number int) {
	for _, report := range reports {
		isSafe, _ := IsReportSafeWithDampener(report)
		if isSafe {
			number++
		}
	}

	return
}
