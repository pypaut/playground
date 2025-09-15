package main

import "github.com/hajimehoshi/ebiten/v2"

type BallUpdateOutput int

const (
	BallUpdateOutputNone BallUpdateOutput = iota
	BallUpdateOutputGameOver
)

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

func (b *Ball) Update(p1, p2 *Player) BallUpdateOutput {
	b.PosX += b.DirX * BallSpeed
	b.PosY += b.DirY * BallSpeed

	// Bottom/top walls collision
	if b.PosY < 0 || b.PosY+b.Size > WinH {
		b.PosY = clamp(b.PosY, 0, WinH-b.Size)
		b.DirY *= -1
	}

	// Left/right walls collision
	if b.PosX < 0 || b.PosX+b.Size > WinW {
		return BallUpdateOutputGameOver
	}

	oldPosX := b.PosX
	oldPosY := b.PosY

	// Check player collision
	if b.Collides(p1) || b.Collides(p2) {
		b.PosX = oldPosX
		b.PosY = oldPosY
		b.DirX *= -1
	}

	return BallUpdateOutputNone
}

func (b *Ball) Reset() {
	b.PosX = (WinW - b.Size) / 2
	b.PosY = (WinH - b.Size) / 2
	b.DirX = 0
	b.DirY = 0
	b.Opt.GeoM.Reset()
	b.Opt.GeoM.Translate(b.PosX, b.PosY)
}
