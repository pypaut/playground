package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type GameOverOutput int

const (
	GameOverOutputNothing GameOverOutput = iota
	GameOverOutputTryAgain
	GameOverOutputQuit
)

type GameOverMenu struct {
	image *ebiten.Image
	opt   *ebiten.DrawImageOptions

	tryAgainButton *Button
	quitButton     *Button

	isEnabled bool
}

func NewGameOverMenu() *GameOverMenu {
	backgroundImg := ebiten.NewImage(WinW/2, WinH/2)
	backgroundImg.Fill(image.White)

	bgOpt := &ebiten.DrawImageOptions{}
	bgOpt.GeoM.Translate(MenuPosX, MenuPosY)

	tryAgainButton := CreateTryAgainButton()
	quitButton := CreateQuitButton()

	return &GameOverMenu{
		image:          backgroundImg,
		opt:            bgOpt,
		tryAgainButton: tryAgainButton,
		quitButton:     quitButton,
	}
}

func (m *GameOverMenu) Update() GameOverOutput {
	if m.tryAgainButton.Update() {
		m.isEnabled = false
		return GameOverOutputTryAgain
	}

	if m.quitButton.Update() {
		return GameOverOutputQuit
	}

	return GameOverOutputNothing
}

func (m *GameOverMenu) IsEnabled() bool {
	return m.isEnabled
}

func (m *GameOverMenu) Toggle() {
	m.isEnabled = !m.isEnabled
}

func (m *GameOverMenu) Draw(screen *ebiten.Image) {
	m.tryAgainButton.Draw(m.image)
	m.quitButton.Draw(m.image)
	screen.DrawImage(m.image, m.opt)
}
