package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type GameOverOutput int

const (
	Nothing GameOverOutput = iota
	TryAgain
	Quit
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

	tryAgainButton := CreateResumeButton()
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
		return TryAgain
	}

	if m.quitButton.Update() {
		return Quit
	}

	return Nothing
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
