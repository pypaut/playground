package main

import "github.com/hajimehoshi/ebiten/v2"

type Ball struct {
	Img *ebiten.Image
	Opt *ebiten.DrawImageOptions

	PosX, PosY float64
	DirX, DirY float64

	Size float64
}

func (b *Ball) Collides(p *Player) bool {
	tooHigh := b.PosY+b.Size < p.PosY
	tooLow := b.PosY > p.PosY+p.Height
	tooRight := b.PosX > p.PosX+p.Width
	tooLeft := b.PosX+b.Size < p.PosX

	if tooHigh || tooLow || tooRight || tooLeft {
		return false
	}

	return true
}
