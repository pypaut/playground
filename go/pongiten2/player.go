package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	Img *ebiten.Image
	Opt *ebiten.DrawImageOptions

	PosX, PosY float64

	UpKey   ebiten.Key
	DownKey ebiten.Key
}

func (p *Player) Update() error {
	if ebiten.IsKeyPressed(p.DownKey) {
		p.PosY += PlayerSpeed
	}

	if ebiten.IsKeyPressed(p.UpKey) {
		p.PosY -= PlayerSpeed
	}

	p.PosY = clamp(p.PosY, 0, float64(WinH-p.Img.Bounds().Size().Y))

	p.Opt.GeoM.Reset()
	p.Opt.GeoM.Translate(p.PosX, p.PosY)

	return nil
}

func clamp(value float64, min float64, max float64) float64 {
	if value < min {
		return min
	}

	if value > max {
		return max
	}

	return value
}
