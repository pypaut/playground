package item

import "github.com/hajimehoshi/ebiten/v2"

type Item struct {
	Name        string
	Image       *ebiten.Image
	DrawOptions *ebiten.DrawImageOptions
}

func (i *Item) Draw(screen *ebiten.Image) {
	screen.DrawImage(i.Image, i.DrawOptions)
}
