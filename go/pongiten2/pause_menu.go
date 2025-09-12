package main

import (
	"fmt"
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
	quitButton Button

	isEnabled bool
}

func NewPauseMenu() *PauseMenu {
	backgroundImg := ebiten.NewImage(WinW/2, WinH/2)
	backgroundImg.Fill(image.White)

	bgSize := backgroundImg.Bounds().Size()
	bgPosX := float64(WinW-bgSize.X) / 2
	bgPosY := float64(WinH-bgSize.Y) / 2
	bgPos := image.Point{int(bgPosX), int(bgPosY)}

	bgOpt := &ebiten.DrawImageOptions{}
	bgOpt.GeoM.Translate(bgPosX, bgPosY)

	resumeButton := createResumeButton(bgSize, bgPos)
	quitButton := createQuitButton(bgSize, bgPos)

	return &PauseMenu{
		BackgroundImage: backgroundImg,
		BackgroundOpt:   bgOpt,
		BackgroundPosX:  bgPosX,
		BackgroundPosY:  bgPosY,
		resumeButton: resumeButton,
		quitButton: quitButton,
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
	if pm.isEnabled {
		pm.isEnabled = false
	} else {
		pm.isEnabled = true
	}
}

func (pm *PauseMenu) Draw(screen *ebiten.Image) {
	pm.resumeButton.Draw(pm.BackgroundImage)
	pm.quitButton.Draw(pm.BackgroundImage)
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

	buttonPos := image.Point{
		X: (bgSize.X-buttonSizeX)/2,
		Y: (bgSize.Y-buttonSizeY)/5,
	}

	return NewButton(
		buttonPos,
		defaultImg.Bounds().Size(),
		bgSize,
		bgPos,
		defaultImg,
		hoveredImg,
		clickedImg,
	)
}

func createQuitButton(bgSize image.Point, bgPos image.Point) Button {
	buttonSizeX := bgSize.X / 3
	buttonSizeY := bgSize.Y / 5

	defaultImg := ebiten.NewImage(buttonSizeX, buttonSizeY)
	defaultImg.Fill(image.Black)

	hoveredImg := ebiten.NewImage(buttonSizeX, buttonSizeY)
	hoveredImg.Fill(color.RGBA{R: 50, G: 50, B: 50, A: 255})

	clickedImg := ebiten.NewImage(buttonSizeX, buttonSizeY)
	clickedImg.Fill(color.RGBA{R: 50, G: 100, B: 100, A: 255})

	buttonPos := image.Point{
		X: (bgSize.X-buttonSizeX)/2,
		Y: (bgSize.Y-buttonSizeY)*4/5,
	}

	return NewButton(
		buttonPos,
		defaultImg.Bounds().Size(),
		bgSize,
		bgPos,
		defaultImg,
		hoveredImg,
		clickedImg,
	)
}
