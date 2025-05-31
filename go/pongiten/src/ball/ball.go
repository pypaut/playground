package ball

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"math"
	"math/rand"
	"pongiten/src/player"

	"image/color"
)

type Ball struct {
	PosX  float64
	PosY  float64
	DirX  float64
	DirY  float64
	Size  float64
	Speed float64
	Color color.Color
}

func NewBall(screenWidth, screenHeight float64) *Ball {
	dirX := rand.Float64() - 0.5
	dirY := rand.Float64() - 0.5

	return &Ball{
		PosX:  screenWidth/2 - 5,
		PosY:  screenHeight/2 - 5,
		DirX:  dirX,
		DirY:  dirY,
		Size:  10,
		Speed: 10,
		Color: color.RGBA{R: 150, B: 150, A: 255},
	}
}

func (b *Ball) Update(
	screenWidth, screenHeight float64,
	isRunning, hasStarted bool,
	player1, player2 *player.Player,
) (isStillRunning bool) {
	b.DirX, b.DirY = normalize(b.DirX, b.DirY)

	if hasStarted && isRunning {
		b.moveVector(b.DirX*b.Speed/2, b.DirY*b.Speed/2)
	}

	// Check if the ball is out
	if b.PosX < 0-b.Size || screenWidth < b.PosX {
		return false
	}

	// Wall collision
	if b.PosY < 0 || screenHeight < b.PosY+b.Size {
		b.DirY = -b.DirY
	}

	// Player 1 collision
	if b.PosX < player1.PosX+player1.Width &&
		b.PosY < player1.PosY+player1.Height+b.Size &&
		player1.PosY-b.Size < b.PosY {
		b.DirX = -b.DirX
	}

	// Player 2 collision
	if player2.PosX < b.PosX+b.Size &&
		b.PosY < player2.PosY+player2.Height+b.Size &&
		player2.PosY-b.Size < b.PosY {
		b.DirX = -b.DirX
	}

	return true
}

func normalize(x, y float64) (float64, float64) {
	norm := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
	return x / norm, y / norm
}

func (b *Ball) moveVector(dirX, dirY float64) {
	b.PosX += dirX
	b.PosY += dirY
}

func (b *Ball) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, b.PosX, b.PosY, b.Size, b.Size, b.Color)
}
