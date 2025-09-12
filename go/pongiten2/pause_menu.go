package main

import (
	"image/color"
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

	resumeButton := createResumeButton(bgSize, image.Point{int(bgPosX), int(bgPosY)})

	return &PauseMenu{
		BackgroundImage: backgroundImg,
		BackgroundOpt:   bgOpt,
		BackgroundPosX:  bgPosX,
		BackgroundPosY:  bgPosY,
		resumeButton: resumeButton,
	}
}

func (pm *PauseMenu) Update() {
	if pm.resumeButton.Update() {
		pm.isEnabled = false
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

func createResumeButton(bgSize image.Point, bgPos image.Point) Button {
	buttonSizeX := bgSize.X / 3
	buttonSizeY := bgSize.Y / 5

	defaultImg := ebiten.NewImage(buttonSizeX, buttonSizeY)
	defaultImg.Fill(image.Black)

	hoveredImg := ebiten.NewImage(buttonSizeX, buttonSizeY)
	hoveredImg.Fill(color.RGBA{R: 50, G: 50, B: 50, A: 255})

	clickedImg := ebiten.NewImage(buttonSizeX, buttonSizeY)
	clickedImg.Fill(color.RGBA{R: 50, G: 100, B: 100, A: 255})

	resumeButtonPos := image.Point{
		X: (bgSize.X-buttonSizeX)/2,
		Y: (bgSize.Y-buttonSizeY)/2,
	}

	return NewButton(
		resumeButtonPos,
		defaultImg.Bounds().Size(),
		bgSize,
		bgPos,
		defaultImg,
		hoveredImg,
		clickedImg,
	)
}
