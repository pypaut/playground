package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const WinW = 1920
const WinH = 1080

const PlayerSpeed = 10
const BallSpeed = 10

const BallSize = 15
const PlayerWidth = 15
const PlayerHeight = 100

func main() {
	ebiten.SetWindowSize(WinW, WinH)
	ebiten.SetWindowTitle("Pong")

	game := NewGame()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
