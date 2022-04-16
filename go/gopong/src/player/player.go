package player

import "github.com/veandco/go-sdl2/sdl"

type Player struct {
	Rect sdl.Rect
	Color sdl.Color
	ControlType int
}