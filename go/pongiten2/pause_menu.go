package main

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type PauseMenu struct {
	BackgroundImage                *ebiten.Image
	BackgroundOpt                  *ebiten.DrawImageOptions
	BackgroundPosX, BackgroundPosY float64

	QuitButtonImage                *ebiten.Image
	QuitButtonOpt                  *ebiten.DrawImageOptions
	QuitButtonPosX, QuitButtonPosY float64

	QuitButtonIsHovered bool
}

func NewPauseMenu() *PauseMenu {
	backgroundImg := ebiten.NewImage(WinW/2, WinH/2)
	backgroundImg.Fill(image.White)

	bgSize := backgroundImg.Bounds().Size()
	bgPosX := float64(WinW-bgSize.X) / 2
	bgPosY := float64(WinH-bgSize.Y) / 2

	bgOpt := &ebiten.DrawImageOptions{}
	bgOpt.GeoM.Translate(bgPosX, bgPosY)

	quitButtonImg := ebiten.NewImage(bgSize.X/3, bgSize.Y/5)
	quitButtonImg.Fill(image.Black)

	buttonSize := quitButtonImg.Bounds().Size()
	buttonPosX := float64(bgSize.X-buttonSize.X) / 2
	buttonPosY := float64(bgSize.Y-buttonSize.Y) / 2

	quitButtonOpt := &ebiten.DrawImageOptions{}
	quitButtonOpt.GeoM.Translate(buttonPosX, buttonPosY)

	backgroundImg.DrawImage(quitButtonImg, quitButtonOpt)

	return &PauseMenu{
		BackgroundImage: backgroundImg,
		BackgroundOpt:   bgOpt,
		BackgroundPosX:  bgPosX,
		BackgroundPosY:  bgPosY,

		QuitButtonImage: quitButtonImg,
		QuitButtonOpt:   quitButtonOpt,
		QuitButtonPosX:  buttonPosX,
		QuitButtonPosY:  buttonPosY,
	}
}

func (pm *PauseMenu) Update() {
	// Check if is hovered
	mouseX, mouseY := ebiten.CursorPosition()
	buttonSize := pm.QuitButtonImage.Bounds().Size()
	posX := int(pm.BackgroundPosX + pm.QuitButtonPosX)
	posY := int(pm.BackgroundPosY + pm.QuitButtonPosY)

	if posX < mouseX && mouseX < posX+buttonSize.X &&
		posY < mouseY && mouseY < posY+buttonSize.Y {
		pm.QuitButtonIsHovered = true
	} else {
		pm.QuitButtonIsHovered = false
	}
}

func (pm *PauseMenu) Draw(screen *ebiten.Image) {
	if pm.QuitButtonIsHovered {
		pm.QuitButtonImage.Fill(color.RGBA{R: 50, G: 50, B: 50, A: 255})
	} else {
		pm.QuitButtonImage.Fill(image.Black)
	}

	pm.BackgroundImage.DrawImage(pm.QuitButtonImage, pm.QuitButtonOpt)
	screen.DrawImage(pm.BackgroundImage, pm.BackgroundOpt)
}
