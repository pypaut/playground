package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"

	"pongiten/src/player"
)

// Game implements ebiten.Game interface.
type Game struct {
	Width   int
	Height  int
	Player1 player.Player
}

func NewGame() *Game {
	width := 1000
	height := 800

	player1 := player.Player{
		PosX:   100.0,
		PosY:   float64(height)/2 - 50,
		Width:  10.0,
		Height: 100.0,
		Speed:  10,
		Color:  color.RGBA{150, 0, 150, 255},
	}

	return &Game{Width: width, Height: height, Player1: player1}
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	g.Player1.Update(float64(g.Width), float64(g.Height))
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	g.Player1.Draw(screen)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Width, g.Height
}
