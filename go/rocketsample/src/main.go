package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth, screenHeight = 1000, 800
)

var (
	background *ebiten.Image
	spaceship  *ebiten.Image
	playerOne  player
)

type player struct {
	image      *ebiten.Image
	xPos, yPos float64
	speed      float64
}

func loadAsset(path string) *ebiten.Image {
	asset, _, err := ebitenutil.NewImageFromFile(path, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	return asset
}

func init() {
	background = loadAsset("assets/space.png")
	spaceship = loadAsset("assets/spaceship.png")
	playerOne = player{spaceship, screenWidth / 2.0, screenHeight / 2.0, 4}
}

func update(screen *ebiten.Image) error {
	movePlayer()

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(background, op)

	playerOp := &ebiten.DrawImageOptions{}
	playerOp.GeoM.Translate(playerOne.xPos, playerOne.yPos)
	screen.DrawImage(playerOne.image, playerOp)

	return nil
}

func movePlayer() {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		playerOne.yPos -= playerOne.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		playerOne.yPos += playerOne.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		playerOne.xPos -= playerOne.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		playerOne.xPos += playerOne.speed
	}
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "Hello, World!"); err != nil {
		log.Fatal(err)
	}
}
