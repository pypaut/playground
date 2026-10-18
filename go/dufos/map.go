package main

import "math"

type Point struct {
	X, Y float64
}

type Map struct {
	WalkableTiles []Point
	Size          Point
}

var myMap = &Map{
	WalkableTiles: []Point{
		{X: 0, Y: 0},
		{X: 4, Y: 5},
		{X: 5, Y: 5},
		{X: 5, Y: 6},
		{X: 10, Y: 0},
		{X: 0, Y: 10},
		{X: 0, Y: 11},
		{X: 0, Y: 12},
		{X: 0, Y: 13},
		{X: 0, Y: 14},
		{X: 0, Y: 15},
	},
	Size: Point{X: 50, Y: 50},
}

func (p *Point) Normalize() {
	norm := math.Sqrt(p.X*p.X + p.Y*p.Y)
	if norm != 0 {
		p.X /= norm
		p.Y /= norm
	}
}
