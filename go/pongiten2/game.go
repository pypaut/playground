package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Player1 *Player
	Player2 *Player
	Ball    *Ball

	IsRunning bool
}

func NewGame() *Game {
	// Player 1
	p1Img := ebiten.NewImage(15, 100)
	p1Img.Fill(color.White)

	player1 := &Player{
		Img:     p1Img,
		Opt:     &ebiten.DrawImageOptions{},
		PosX:    float64(WinW) / 10,
		PosY:    float64(WinH-p1Img.Bounds().Dy()) / 2,
		DownKey: ebiten.KeyD,
		UpKey:   ebiten.KeyE,
	}

	// Player 2
	p2Img := ebiten.NewImage(15, 100)
	p2Img.Fill(color.White)

	player2 := &Player{
		Img:     p2Img,
		Opt:     &ebiten.DrawImageOptions{},
		PosX:    float64(WinW*9/10 - p2Img.Bounds().Size().X),
		PosY:    float64(WinH-p2Img.Bounds().Size().Y) / 2,
		UpKey:   ebiten.KeyI,
		DownKey: ebiten.KeyK,
	}

	// Ball
	bImg := ebiten.NewImage(BallSize, BallSize)
	bImg.Fill(color.White)

	ball := &Ball{
		Img:  bImg,
		Opt:  &ebiten.DrawImageOptions{},
		PosX: (WinW - BallSize) / 2,
		PosY: (WinH - BallSize) / 2,
		DirX: 0,
		DirY: 0,
	}

	return &Game{
		Player1: player1,
		Player2: player2,
		Ball:    ball,
	}
}

func (g *Game) Update() error {
	if err := g.Player1.Update(); err != nil {
		return err
	}

	if err := g.Player2.Update(); err != nil {
		return err
	}

	g.Ball.Opt.GeoM.Reset()
	g.Ball.Opt.GeoM.Translate(g.Ball.PosX, g.Ball.PosY)

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.IsRunning = true
		g.Ball.DirX = 0.5
		g.Ball.DirY = 0.5
	}

	if !g.IsRunning {
		return nil
	}

	g.Ball.PosX += g.Ball.DirX * BallSpeed
	g.Ball.PosY += g.Ball.DirY * BallSpeed

	// Bottom/top walls collision
	if g.Ball.PosY < 0 || g.Ball.PosY+BallSize > WinH {
		g.Ball.PosY = clamp(g.Ball.PosY, 0, float64(WinH-BallSize))
		g.Ball.DirY *= -1
	}

	// Left/right walls collision
	if g.Ball.PosX < 0 || g.Ball.PosX+BallSize > WinW {
		return fmt.Errorf("Game over!")
	}

	oldPosX := g.Ball.PosX
	oldPosY := g.Ball.PosY

	// Check player collision
	if g.Ball.Collides(g.Player1) || g.Ball.Collides(g.Player2) {
		g.Ball.PosX = oldPosX
		g.Ball.PosY = oldPosY
		g.Ball.DirX *= -1
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.Player1.Img, g.Player1.Opt)
	screen.DrawImage(g.Player2.Img, g.Player2.Opt)
	screen.DrawImage(g.Ball.Img, g.Ball.Opt)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WinW, WinH
}
