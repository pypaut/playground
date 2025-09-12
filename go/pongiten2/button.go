package main

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Button interface {
	Update() (wasClicked bool)
	Draw(menu *ebiten.Image)
}

type button struct {
	DefaultImage *ebiten.Image
	HoveredImage *ebiten.Image
	ClickedImage *ebiten.Image
	Opt          *ebiten.DrawImageOptions

	IsHovered bool
	IsClicked bool

	PosX, PosY                 float64
	PosXAbsolute, PosYAbsolute float64
	SizeX, SizeY               float64
}

func NewButton(
	buttonPos image.Point,
	buttonSize image.Point,
	menuSize image.Point,
	menuPos image.Point,
	defaultImg *ebiten.Image,
	hoverImg *ebiten.Image,
	clickImg *ebiten.Image,
) Button {
	buttonPosXAbsolute := float64(menuPos.X + buttonPos.X)
	buttonPosYAbsolute := float64(menuPos.Y + buttonPos.Y)

	buttonOpt := &ebiten.DrawImageOptions{}
	buttonOpt.GeoM.Translate(float64(buttonPos.X), float64(buttonPos.Y))

	return &button{
		DefaultImage: defaultImg,
		Opt:          buttonOpt,
		HoveredImage: hoverImg,
		ClickedImage: clickImg,

		PosX: float64(buttonPos.X),
		PosY: float64(buttonPos.Y),

		PosXAbsolute: buttonPosXAbsolute,
		PosYAbsolute: buttonPosYAbsolute,

		SizeX: float64(buttonSize.X),
		SizeY: float64(buttonSize.Y),
	}
}

func (b *button) Update() (wasClicked bool) {
	// Check if is hovered
	mouseX, mouseY := ebiten.CursorPosition()

	posX := int(b.PosXAbsolute)
	posY := int(b.PosYAbsolute)

	sizeX := int(b.SizeX)
	sizeY := int(b.SizeY)

	if posX <= mouseX && mouseX <= posX+sizeX &&
		posY <= mouseY && mouseY <= posY+sizeY {
		b.IsHovered = true
	} else {
		b.IsHovered = false
	}

	// Check if clicked
	if b.IsHovered && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		b.IsClicked = true
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		if b.IsClicked && b.IsHovered {
			b.IsClicked = false
			return true
		}

		if b.IsClicked {
			b.IsClicked = false
		}
	}

	return false
}

func (b *button) Draw(menu *ebiten.Image) {
	toDraw := b.DefaultImage

	if b.IsHovered {
		toDraw = b.HoveredImage
	}

	if b.IsClicked {
		toDraw = b.ClickedImage
	}

	menu.DrawImage(toDraw, b.Opt)
}
