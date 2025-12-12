package main

import (
	"os"
	"strings"
)

type Matrix struct {
	Height  int
	Width   int
	Content []int
}

func LoadMatrix(path string) *Matrix {
	b, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	Conversion := map[rune]int{'.': 0, '@': 1}

	strMat := string(b)
	lines := strings.Split(strMat, "\n")

	m := Matrix{
		Height:  0,
		Width:   len(lines[0]),
		Content: []int{},
	}

	for _, l := range lines {
		if l == "" {
			continue
		}

		m.Height++

		for _, c := range l {
			m.Content = append(m.Content, Conversion[c])
		}
	}

	return &m
}

func (m *Matrix) Get(i, j int) int {
	return m.Content[i*m.Width+j]
}

func (m *Matrix) Set(i, j, val int) {
	m.Content[i*m.Width+j] = val
}

func (m *Matrix) CountAdjascentRolls(i, j int) (rollsCount int) {
	coordinatesAroundRoll := []struct{ I, J int }{
		{
			I: i - 1, J: j - 1,
		},
		{
			I: i - 1, J: j,
		},
		{
			I: i - 1, J: j + 1,
		},
		{
			I: i, J: j + 1,
		},
		{
			I: i + 1, J: j + 1,
		},
		{
			I: i + 1, J: j,
		},
		{
			I: i + 1, J: j - 1,
		},
		{
			I: i, J: j - 1,
		},
	}

	for _, c := range coordinatesAroundRoll {
		if c.I < 0 || c.I >= m.Height || c.J < 0 || c.J >= m.Width {
			continue
		}

		rollsCount += m.Get(c.I, c.J)
	}

	return
}
