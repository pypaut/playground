package game

import (
	"errors"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"

	"pongiten/src/ball"
	"pongiten/src/player"
)

// Game implements ebiten.Game interface.
type Game struct {
	Width      int
	Height     int
	Player1    *player.Player
	Player2    *player.Player
	Ball       *ball.Ball
	HasStarted bool
	IsRunning  bool
}

func NewGame() *Game {
	width := 1000
	height := 800

	player1 := player.NewPlayer(
		100.0,
		float64(height)/2-50,
		10.0,
		100.0,
		10,
		color.RGBA{R: 150, B: 150, A: 255},
		1,
	)

	player2 := player.NewPlayer(
		float64(width)-100,
		float64(height)/2-50,
		10.0,
		100.0,
		10,
		color.RGBA{R: 150, B: 150, A: 255},
		2,
	)

	b := ball.NewBall(float64(width), float64(height))

	return &Game{
		Width:      width,
		Height:     height,
		Player1:    player1,
		Player2:    player2,
		Ball:       b,
		HasStarted: false,
		IsRunning:  true,
	}
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	if !g.HasStarted && ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.HasStarted = true
	}

	g.Player1.Update(float64(g.Width), float64(g.Height))
	g.Player2.Update(float64(g.Width), float64(g.Height))
	g.IsRunning = g.Ball.Update(float64(g.Width), float64(g.Height), g.IsRunning, g.HasStarted)
	if !g.IsRunning {
		return errors.New("game over")
	}

	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	g.Player1.Draw(screen)
	g.Player2.Draw(screen)
	g.Ball.Draw(screen)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Width, g.Height
}
