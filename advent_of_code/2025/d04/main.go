package main

import "fmt"

func main() {
	m := LoadMatrix("sample")
	fmt.Println(m)
}

func CountAccessibleRolls(m *Matrix) (count int) {
	for i := range m.Height {
		for j := range m.Width {
			if m.IsAccessible(i, j) {
				count++
			}
		}
	}

	return
}
