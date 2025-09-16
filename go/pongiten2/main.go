package main

import (
	"bytes"
	_ "embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/text/language"
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

var TextFace *text.GoTextFace

func main() {
	ebiten.SetWindowSize(WinW, WinH)
	ebiten.SetWindowTitle("Pong")

	kongFaceSource, err := text.NewGoTextFaceSource(bytes.NewReader(kongTTF))
	if err != nil {
		log.Fatal(err)
	}

	TextFace = &text.GoTextFace{
		Source:    kongFaceSource,
		Direction: text.DirectionLeftToRight,
		Size:      24,
		Language:  language.English,
	}

	game := NewGame()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
