package structs

type Object struct {
	PosX, PosY    float64
	DirX, DirY    float64
	Width, Height float64
}

type Direction struct {
	X, Y float64
}

const (
	WinW = 1920
	WinH = 1080
)
