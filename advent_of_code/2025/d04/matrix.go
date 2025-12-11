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

func (m *Matrix) Submatrix(i1, i2, j1, j2 int) (sub *Matrix) {
	i1 = Clamp(i1, 0, m.Height-1)
	i2 = Clamp(i2, 0, m.Height-1)
	j1 = Clamp(j1, 0, m.Width-1)
	j2 = Clamp(j2, 0, m.Width-1)

	sub.Height = i2 - i1 - 1
	sub.Width = j2 - j1 - 1

	for i := range sub.Height {
		for j := range sub.Width {
			sub.Content = append(sub.Content, m.Get(i1+i, j1+j))
		}
	}

	return
}

func (m *Matrix) IsAccessible(i, j int) bool {
	spaceCounter := 0

	// I'll try to avoid this case handling using a submatrix technique, see above

	// Left
	if j > 0 {

	}

	// Top
	if i > 0 {

	}

	// Right
	if j < m.Width-1 {

	}

	// Bottom
	if i < m.Height-1 {

	}

	return spaceCounter > 4
}

func Clamp(v, min, max int) int {
	if v < min {
		return min
	}

	if v > max {
		return max
	}

	return v
}
