package item

import "github.com/hajimehoshi/ebiten/v2"

type Item struct {
	Name        string
	Image       *ebiten.Image
	DrawOptions *ebiten.DrawImageOptions
	PosX, PosY  int
	Size        int
}

func (i *Item) Draw(screen *ebiten.Image) {
	screen.DrawImage(i.Image, i.DrawOptions)
}

func (i *Item) IsHovered() bool {
	mouseX, mouseY := ebiten.CursorPosition()
	return (i.PosX < mouseX && mouseX < i.PosX+i.Size) &&
		(i.PosY < mouseY && mouseY < i.PosY+i.Size)
}
