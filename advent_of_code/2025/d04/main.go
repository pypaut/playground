package main

import "fmt"

func main() {
	m := LoadMatrix("input")
	fmt.Printf("Part 1: %d\n", CountAccessibleRolls(m))
	fmt.Printf("Part 2: %d\n", RemoveRolls(m))
}

func CountAccessibleRolls(m *Matrix) (count int) {
	for i := range m.Height {
		for j := range m.Width {
			if m.Get(i, j) == 1 && m.CountAdjascentRolls(i, j) < 4 {
				count++
			}
		}
	}

	return
}

func RemoveRolls(m *Matrix) (count int) {
	for {
		currentCount := 0

		for i := range m.Height {
			for j := range m.Width {
				if m.Get(i, j) == 1 && m.CountAdjascentRolls(i, j) < 4 {
					currentCount++
					m.Set(i, j, 0)
				}
			}
		}

		if currentCount == 0 {
			break
		}

		count += currentCount
	}

	return
}
