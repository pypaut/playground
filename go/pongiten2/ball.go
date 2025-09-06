package main

import "github.com/hajimehoshi/ebiten/v2"

type Ball struct {
	Img *ebiten.Image
	Opt *ebiten.DrawImageOptions

	PosX, PosY float64
	DirX, DirY float64
}

func (b *Ball) Collides(p *Player) bool {
	tooHigh := b.PosY+BallSize < p.PosY
	tooLow := b.PosY > p.PosY+p.Height
	tooRight := b.PosX > p.PosX+p.Width
	tooLeft := b.PosX+BallSize < p.PosX

	if tooHigh || tooLow || tooRight || tooLeft {
		return false
	}

	return true
}
