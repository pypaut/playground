package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	Img *ebiten.Image
	Opt *ebiten.DrawImageOptions

	PosX, PosY    float64
	Width, Height float64

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

	p.PosY = clamp(p.PosY, 0, WinH-p.Height)

	p.Opt.GeoM.Reset()
	p.Opt.GeoM.Scale(PlayerScale, PlayerScale)
	p.Opt.GeoM.Translate(p.PosX, p.PosY)

	return nil
}

func (p *Player) Reset() {
	p.PosY = float64(WinH-p.Height) / 2
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
