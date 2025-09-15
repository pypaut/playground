package main

import (
	_ "embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const WinW = 1920
const WinH = 1080

const PlayerSpeed = 10
const BallSpeed = 10

const PlayerScale = 4.0
const BallScale = 0.8

const MenuWidth = WinW / 2
const MenuHeight = WinH / 2

const MenuPosX = (WinW - MenuWidth) / 2
const MenuPosY = (WinH - MenuHeight) / 2

//go:embed assets/kongtext.ttf
var kongTTF []byte

func main() {
	ebiten.SetWindowSize(WinW, WinH)
	ebiten.SetWindowTitle("Pong")

	game := NewGame()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
