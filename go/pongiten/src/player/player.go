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
	Speed  float64
	Color  color.Color
}

func (p *Player) Update(boundWidth, boundHeight float64) {
	if ebiten.IsKeyPressed(ebiten.KeyW) && p.PosY > 0 {
		p.PosY -= p.Speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) && p.PosY+p.Height < boundHeight {
		p.PosY += p.Speed
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, p.PosX, p.PosY, p.Width, p.Height, p.Color)
}
