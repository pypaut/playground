package item

import "github.com/hajimehoshi/ebiten/v2"

type Item struct {
	Name        string
	Image       *ebiten.Image
	DrawOptions *ebiten.DrawImageOptions
	PosX, PosY  float64
	Size        float64
}

func (i *Item) Draw(screen *ebiten.Image) {
	screen.DrawImage(i.Image, i.DrawOptions)
}

func (i *Item) IsHovered() bool {
	mouseX, mouseY := ebiten.CursorPosition()
	return (int(i.PosX) < mouseX && mouseX < int(i.PosX+i.Size)) &&
		(int(i.PosY) < mouseY && mouseY < int(i.PosY+i.Size))
}
