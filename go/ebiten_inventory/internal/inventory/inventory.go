package inventory

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"inventory/internal/item"
	"math/rand/v2"
)

type Inventory struct {
	Image       *ebiten.Image
	Items       []*item.Item
	DrawOptions *ebiten.DrawImageOptions
}

func NewInventory() *Inventory {
	itemSize := 100

	// Grid size
	widthItems := 5
	heightItems := 8

	// Overall inventory pos
	var invPosX float64 = 1200
	var invPosY float64 = 200

	// Inventory img
	outPadding := 5
	inPadding := 5
	width := itemSize * widthItems
	height := itemSize * heightItems
	img := ebiten.NewImage(
		width+outPadding*2+inPadding*(widthItems-1),
		height+outPadding*2+inPadding*(heightItems-1),
	)
	img.Fill(color.White)

	var items []*item.Item

	for i := 0; i < widthItems; i++ {
		for j := 0; j < heightItems; j++ {
			// Build random image
			itemImg := ebiten.NewImage(itemSize, itemSize)
			itemImg.Fill(color.Gray{Y: uint8(rand.IntN(200))})

			// Setup draw options for position
			drawOptions := &ebiten.DrawImageOptions{}
			drawOptions.GeoM.Reset()
			drawOptions.GeoM.Translate(
				invPosX+float64(i*itemSize)+float64(outPadding)+float64(inPadding*i),
				invPosY+float64(j*itemSize)+float64(outPadding)+float64(inPadding*j),
			)

			// Create the item
			items = append(items, &item.Item{
				Name:        fmt.Sprintf("item%d", i),
				Image:       itemImg,
				DrawOptions: drawOptions,
			})
		}
	}

	drawOptions := ebiten.DrawImageOptions{}
	drawOptions.GeoM.Reset()
	drawOptions.GeoM.Translate(invPosX, invPosY)

	return &Inventory{
		Image:       img,
		Items:       items,
		DrawOptions: &drawOptions,
	}
}

func (i *Inventory) Draw(screen *ebiten.Image) {
	screen.DrawImage(i.Image, i.DrawOptions)

	for _, i := range i.Items {
		if i != nil {
			i.Draw(screen)
		}
	}

	return
}
