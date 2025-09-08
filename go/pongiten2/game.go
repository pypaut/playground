package main

import (
	"fmt"
	"image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	Player1 *Player
	Player2 *Player
	Ball    *Ball

	IsRunning bool
	IsPaused  bool

	Menu *PauseMenu
}

func NewGame() *Game {
	playerImg := loadSprite("assets/player.png")
	s := playerImg.Bounds().Size()

	playerWidth := float64(s.X * PlayerScale)
	playerHeight := float64(s.Y * PlayerScale)

	player1 := &Player{
		Img:     playerImg,
		Opt:     &ebiten.DrawImageOptions{},
		PosX:    float64(WinW) / 10,
		PosY:    float64(WinH-playerHeight) / 2,
		Width:   playerWidth,
		Height:  playerHeight,
		DownKey: ebiten.KeyD,
		UpKey:   ebiten.KeyE,
	}

	// Player 2
	player2 := &Player{
		Img:     playerImg,
		Opt:     &ebiten.DrawImageOptions{},
		PosX:    float64(WinW*9/10 - s.X),
		PosY:    float64(WinH-playerHeight) / 2,
		Width:   playerWidth,
		Height:  playerHeight,
		UpKey:   ebiten.KeyI,
		DownKey: ebiten.KeyK,
	}

	// Ball
	ballImg := loadSprite("assets/ball.png")
	s = ballImg.Bounds().Size()
	ballSize := float64(s.X) * BallScale

	ball := &Ball{
		Img: ballImg,
		Opt: &ebiten.DrawImageOptions{},

		PosX: (WinW - ballSize) / 2,
		PosY: (WinH - ballSize) / 2,
		Size: ballSize,

		DirX: 0,
		DirY: 0,
	}

	return &Game{
		Player1: player1,
		Player2: player2,
		Ball:    ball,
		Menu:    NewPauseMenu(),
	}
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.IsPaused = !g.IsPaused
	}

	if g.IsPaused {
		g.Menu.Update()
		return nil
	}

	if err := g.Player1.Update(); err != nil {
		return err
	}

	if err := g.Player2.Update(); err != nil {
		return err
	}

	g.Ball.Opt.GeoM.Reset()
	g.Ball.Opt.GeoM.Scale(BallScale, BallScale)
	g.Ball.Opt.GeoM.Translate(g.Ball.PosX, g.Ball.PosY)

	if ebiten.IsKeyPressed(ebiten.KeySpace) && !g.IsRunning {
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
	if g.Ball.PosY < 0 || g.Ball.PosY+g.Ball.Size > WinH {
		g.Ball.PosY = clamp(g.Ball.PosY, 0, WinH-g.Ball.Size)
		g.Ball.DirY *= -1
	}

	// Left/right walls collision
	if g.Ball.PosX < 0 || g.Ball.PosX+g.Ball.Size > WinW {
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

	if g.IsPaused {
		g.Menu.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WinW, WinH
}

func loadSprite(path string) *ebiten.Image {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	return ebiten.NewImageFromImage(img)
}
