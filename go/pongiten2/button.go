package main

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Button struct {
	defaultImage *ebiten.Image
	hoverImage   *ebiten.Image
	clickImage   *ebiten.Image
	opt          *ebiten.DrawImageOptions

	label    string
	labelOpt *text.DrawOptions

	isHovered bool
	isClicked bool

	posX, posY                 float64
	posXAbsolute, posYAbsolute float64
	sizeX, sizeY               float64
}

func NewButton(
	buttonPos image.Point,
	buttonSize image.Point,
	defaultImg *ebiten.Image,
	hoverImg *ebiten.Image,
	clickImg *ebiten.Image,
	label string,
) *Button {
	buttonPosXAbsolute := float64(MenuPosX + buttonPos.X)
	buttonPosYAbsolute := float64(MenuPosY + buttonPos.Y)

	buttonOpt := &ebiten.DrawImageOptions{}
	buttonOpt.GeoM.Translate(float64(buttonPos.X), float64(buttonPos.Y))

	labelOpt := &text.DrawOptions{
		DrawImageOptions: ebiten.DrawImageOptions{},
		LayoutOptions: text.LayoutOptions{
			PrimaryAlign:   text.AlignCenter,
			SecondaryAlign: text.AlignCenter,
		},
	}
	labelOpt.GeoM.Translate(float64(buttonSize.X/2), float64(buttonSize.Y/2))

	return &Button{
		defaultImage: defaultImg,
		opt:          buttonOpt,
		hoverImage:   hoverImg,
		clickImage:   clickImg,

		posX: float64(buttonPos.X),
		posY: float64(buttonPos.Y),

		posXAbsolute: buttonPosXAbsolute,
		posYAbsolute: buttonPosYAbsolute,

		sizeX: float64(buttonSize.X),
		sizeY: float64(buttonSize.Y),

		label:    label,
		labelOpt: labelOpt,
	}
}

func (b *Button) Update() (wasClicked bool) {
	// Check if is hovered
	mouseX, mouseY := ebiten.CursorPosition()

	posX := int(b.posXAbsolute)
	posY := int(b.posYAbsolute)

	sizeX := int(b.sizeX)
	sizeY := int(b.sizeY)

	if posX <= mouseX && mouseX <= posX+sizeX &&
		posY <= mouseY && mouseY <= posY+sizeY {
		b.isHovered = true
	} else {
		b.isHovered = false
	}

	// Check if clicked
	if b.isHovered && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		b.isClicked = true
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		if b.isClicked && b.isHovered {
			b.isClicked = false
			return true
		}

		if b.isClicked {
			b.isClicked = false
		}
	}

	return false
}

func (b *Button) Draw(menu *ebiten.Image) {
	toDraw := b.defaultImage

	if b.isHovered {
		toDraw = b.hoverImage
	}

	if b.isClicked {
		toDraw = b.clickImage
	}

	text.Draw(toDraw, b.label, TextFace, b.labelOpt)
	menu.DrawImage(toDraw, b.opt)
}

func CreateResumeButton() *Button {
	buttonSizeX := MenuWidth / 3
	buttonSizeY := MenuHeight / 5

	defaultImg := ebiten.NewImage(buttonSizeX, buttonSizeY)
	defaultImg.Fill(image.Black)

	hoveredImg := ebiten.NewImage(buttonSizeX, buttonSizeY)
	hoveredImg.Fill(color.RGBA{R: 50, G: 50, B: 50, A: 255})

	clickedImg := ebiten.NewImage(buttonSizeX, buttonSizeY)
	clickedImg.Fill(color.RGBA{R: 50, G: 100, B: 100, A: 255})

	buttonPos := image.Point{
		X: (MenuWidth - buttonSizeX) / 2,
		Y: (MenuHeight - buttonSizeY) / 5,
	}

	return NewButton(
		buttonPos,
		defaultImg.Bounds().Size(),
		defaultImg,
		hoveredImg,
		clickedImg,
		"Resume",
	)
}

func CreateTryAgainButton() *Button {
	buttonSizeX := MenuWidth / 3
	buttonSizeY := MenuHeight / 5

	defaultImg := ebiten.NewImage(buttonSizeX, buttonSizeY)
	defaultImg.Fill(image.Black)

	hoveredImg := ebiten.NewImage(buttonSizeX, buttonSizeY)
	hoveredImg.Fill(color.RGBA{R: 50, G: 50, B: 50, A: 255})

	clickedImg := ebiten.NewImage(buttonSizeX, buttonSizeY)
	clickedImg.Fill(color.RGBA{R: 50, G: 100, B: 100, A: 255})

	buttonPos := image.Point{
		X: (MenuWidth - buttonSizeX) / 2,
		Y: (MenuHeight - buttonSizeY) / 5,
	}

	return NewButton(
		buttonPos,
		defaultImg.Bounds().Size(),
		defaultImg,
		hoveredImg,
		clickedImg,
		"Try again",
	)
}

func CreateQuitButton() *Button {
	buttonSizeX := MenuWidth / 3
	buttonSizeY := MenuHeight / 5

	defaultImg := ebiten.NewImage(buttonSizeX, buttonSizeY)
	defaultImg.Fill(image.Black)

	hoveredImg := ebiten.NewImage(buttonSizeX, buttonSizeY)
	hoveredImg.Fill(color.RGBA{R: 50, G: 50, B: 50, A: 255})

	clickedImg := ebiten.NewImage(buttonSizeX, buttonSizeY)
	clickedImg.Fill(color.RGBA{R: 50, G: 100, B: 100, A: 255})

	buttonPos := image.Point{
		X: (MenuWidth - buttonSizeX) / 2,
		Y: (MenuHeight - buttonSizeY) * 4 / 5,
	}

	return NewButton(
		buttonPos,
		defaultImg.Bounds().Size(),
		defaultImg,
		hoveredImg,
		clickedImg,
		"Quit",
	)
}
