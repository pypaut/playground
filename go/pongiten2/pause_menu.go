package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type PauseMenu struct {
	BackgroundImage                *ebiten.Image
	BackgroundOpt                  *ebiten.DrawImageOptions
	BackgroundPosX, BackgroundPosY float64

	QuitButtonImage                *ebiten.Image
	QuitButtonOpt                  *ebiten.DrawImageOptions
	QuitButtonPosX, QuitButtonPosY float64

	QuitButtonIsHovered  bool
	QuitButtonWasClicked bool

	resumeButton Button

	isEnabled bool
}

func NewPauseMenu() *PauseMenu {
	backgroundImg := ebiten.NewImage(WinW/2, WinH/2)
	backgroundImg.Fill(image.White)

	bgSize := backgroundImg.Bounds().Size()
	bgPosX := float64(WinW-bgSize.X) / 2
	bgPosY := float64(WinH-bgSize.Y) / 2

	bgOpt := &ebiten.DrawImageOptions{}
	bgOpt.GeoM.Translate(bgPosX, bgPosY)

	bgPos := image.Point{X: int(bgPosX), Y: int(bgPosY)}
	resumeButton := NewResumeButton(bgSize, bgPos)

	return &PauseMenu{
		BackgroundImage: backgroundImg,
		BackgroundOpt:   bgOpt,
		BackgroundPosX:  bgPosX,
		BackgroundPosY:  bgPosY,

		resumeButton: resumeButton,
	}
}

func (pm *PauseMenu) Update() {
	pm.isEnabled = pm.resumeButton.Update()
	if !pm.isEnabled {
		println("false")
	}
}

func (pm *PauseMenu) IsEnabled() bool {
	return pm.isEnabled
}

func (pm *PauseMenu) Toggle() {
	if pm.isEnabled {
		pm.isEnabled = false
	} else {
		pm.isEnabled = true
	}
}

func (pm *PauseMenu) Draw(screen *ebiten.Image) {
	pm.resumeButton.Draw(pm.BackgroundImage)
	screen.DrawImage(pm.BackgroundImage, pm.BackgroundOpt)
}
