package main

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type PauseMenu struct {
	image *ebiten.Image
	opt   *ebiten.DrawImageOptions

	resumeButton *Button
	quitButton   *Button

	isEnabled bool
}

func NewPauseMenu() *PauseMenu {
	backgroundImg := ebiten.NewImage(WinW/2, WinH/2)
	backgroundImg.Fill(image.White)

	bgOpt := &ebiten.DrawImageOptions{}
	bgOpt.GeoM.Translate(MenuPosX, MenuPosY)

	resumeButton := CreateResumeButton()
	quitButton := CreateQuitButton()

	return &PauseMenu{
		image:        backgroundImg,
		opt:          bgOpt,
		resumeButton: resumeButton,
		quitButton:   quitButton,
	}
}

func (pm *PauseMenu) Update() error {
	if pm.resumeButton.Update() {
		pm.isEnabled = false
	}

	if pm.quitButton.Update() {
		return fmt.Errorf("quit button was pressed")
	}

	return nil
}

func (pm *PauseMenu) IsEnabled() bool {
	return pm.isEnabled
}

func (pm *PauseMenu) Toggle() {
	pm.isEnabled = !pm.isEnabled
}

func (pm *PauseMenu) Draw(screen *ebiten.Image) {
	pm.resumeButton.Draw(pm.image)
	pm.quitButton.Draw(pm.image)
	screen.DrawImage(pm.image, pm.opt)
}
