package main

import (
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	TileSize   = 40.0
	screenW    = 480
	screenH    = 320
	tileWidth  = 0.0
	tileHeight = 0.0
)

const (
	CameraSpeed = 6
)

func main() {
	tileWidth = math.Sqrt(TileSize * TileSize * 2)
	tileHeight = tileWidth / 2.0

	ebiten.SetWindowSize(screenW*2, screenH*2)
	ebiten.SetWindowTitle("Dofus")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
