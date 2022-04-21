package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"image/color"
)

type Player struct {
	PosX   float64
	PosY   float64
	Height float64
	Width  float64
	Color  color.Color
}

func (p *Player) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, p.PosX, p.PosY, p.Width, p.Height, p.Color)
}
