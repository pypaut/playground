package main

import (
	"bytes"
	"errors"
	"fmt"
	"image/color"
	"image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font"
	"golang.org/x/text/language"
)

type Game struct {
	Player1 *Player
	Player2 *Player
	Ball    *Ball

	HasStarted bool

	PauseMenu    *PauseMenu
	GameOverMenu *GameOverMenu

	FontFace   font.Face
	FontColor  color.Color
	FaceSource *text.GoTextFaceSource
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

	var kongFaceSource *text.GoTextFaceSource
	textFaceSource, err := text.NewGoTextFaceSource(bytes.NewReader(kongTTF))
	if err != nil {
		log.Fatal(err)
	}
	kongFaceSource = textFaceSource

	return &Game{
		Player1:      player1,
		Player2:      player2,
		Ball:         ball,
		PauseMenu:    NewPauseMenu(),
		GameOverMenu: NewGameOverMenu(),
		FaceSource:   kongFaceSource,
	}
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.PauseMenu.Toggle()
	}

	if g.PauseMenu.IsEnabled() {
		switch g.PauseMenu.Update() {
		case PauseOutputResume:
			g.PauseMenu.Toggle()
		case PauseOutputQuit:
			return errors.New("quit from pause menu")
		case PauseOutputNothing:
			return nil
		default:
			fmt.Printf("unhandled default case for PauseMenu.Update()\n")
			return nil
		}

		return nil
	}

	if g.GameOverMenu.IsEnabled() {
		switch g.GameOverMenu.Update() {
		case GameOverOutputTryAgain:
			g.Reset()
		case GameOverOutputQuit:
			return errors.New("quit from game over menu")
		case GameOverOutputNothing:
			return nil
		default:
			fmt.Printf("unhandled default case for GameOverMenu.Update()\n")
			return nil
		}

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

	if ebiten.IsKeyPressed(ebiten.KeySpace) && !g.HasStarted {
		g.HasStarted = true
		g.Ball.DirX = 0.5
		g.Ball.DirY = 0.5
	}

	if !g.HasStarted {
		return nil
	}

	switch g.Ball.Update(g.Player1, g.Player2) {
	case BallUpdateOutputGameOver:
		g.GameOverMenu.Toggle()
	case BallUpdateOutputNone:
		return nil
	default:
		return nil
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.Player1.Img, g.Player1.Opt)
	screen.DrawImage(g.Player2.Img, g.Player2.Opt)
	screen.DrawImage(g.Ball.Img, g.Ball.Opt)

	textOpt := &text.DrawOptions{
		DrawImageOptions: ebiten.DrawImageOptions{},
		LayoutOptions:    text.LayoutOptions{},
	}

	textOpt.GeoM.Reset()
	textOpt.GeoM.Translate(WinW-WinW/5, WinH-WinH/15)
	f := &text.GoTextFace{
		Source:    g.FaceSource,
		Direction: text.DirectionRightToLeft,
		Size:      24,
		Language:  language.English,
	}
	text.Draw(screen, "hehehe that's my text", f, textOpt)

	if g.PauseMenu.IsEnabled() {
		g.PauseMenu.Draw(screen)
	}

	if g.GameOverMenu.IsEnabled() {
		g.GameOverMenu.Draw(screen)
	}
}

func (g *Game) Reset() {
	g.Ball.Reset()
	g.Player1.Reset()
	g.Player2.Reset()
	g.HasStarted = false
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
