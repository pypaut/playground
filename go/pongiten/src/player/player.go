package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"image/color"
)

type Player struct {
	PosX    float64
	PosY    float64
	Height  float64
	Width   float64
	Speed   float64
	Color   color.Color
	UpKey   ebiten.Key
	DownKey ebiten.Key
}

func NewPlayer(posX, posY, width, height, speed float64, clr color.Color, controls int) *Player {
	var upKey ebiten.Key
	var downKey ebiten.Key

	if controls == 1 {
		upKey = ebiten.KeyW
		downKey = ebiten.KeyS
	} else {
		upKey = ebiten.KeyI
		downKey = ebiten.KeyK
	}

	return &Player{
		PosX:    posX,
		PosY:    posY,
		Height:  height,
		Width:   width,
		Speed:   speed,
		Color:   clr,
		UpKey:   upKey,
		DownKey: downKey,
	}
}

func (p *Player) Update(boundWidth, boundHeight float64) {
	if ebiten.IsKeyPressed(p.UpKey) && p.PosY > 0 {
		p.PosY -= p.Speed
	}

	if ebiten.IsKeyPressed(p.DownKey) && p.PosY+p.Height < boundHeight {
		p.PosY += p.Speed
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, p.PosX, p.PosY, p.Width, p.Height, p.Color)
}
