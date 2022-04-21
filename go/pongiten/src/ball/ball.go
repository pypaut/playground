package ball

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"image/color"
)

type Ball struct {
	PosX  float64
	PosY  float64
	DirX  float64
	DirY  float64
	Size  float64
	Speed float64
	Color color.Color
}

func NewBall(screenWidth, screenHeight float64) *Ball {
	return &Ball{
		PosX:  screenWidth/2 - 5,
		PosY:  screenHeight/2 - 5,
		DirX:  0.0,
		DirY:  0.0,
		Size:  10,
		Speed: 10,
		Color: color.RGBA{150, 0, 150, 255},
	}
}

func (b *Ball) Update(screenWidth, screenHeight float64, isRunning, hasStarted bool) (bool, bool) {
	return false, true
}

func (b *Ball) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, b.PosX, b.PosY, b.Size, b.Size, b.Color)
}
