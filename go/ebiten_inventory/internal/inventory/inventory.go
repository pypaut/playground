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
	HoverImage  *ebiten.Image
	Items       []*item.Item
	DrawOptions *ebiten.DrawImageOptions
	InPadding   int
	OutPadding  int
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

	// Hover image
	hoverImg := ebiten.NewImage(itemSize+2*inPadding, itemSize+2*inPadding)
	hoverImg.Fill(color.RGBA{R: 200, A: 255})

	var items []*item.Item

	for i := 0; i < widthItems; i++ {
		for j := 0; j < heightItems; j++ {
			// Build random image
			itemImg := ebiten.NewImage(itemSize, itemSize)
			itemImg.Fill(color.Gray{Y: uint8(rand.IntN(200))})

			// Setup draw options for position
			drawOptions := &ebiten.DrawImageOptions{}
			drawOptions.GeoM.Reset()
			itemPosX := invPosX + float64(i*itemSize) + float64(outPadding) + float64(inPadding*i)
			itemPosY := invPosY + float64(j*itemSize) + float64(outPadding) + float64(inPadding*j)
			drawOptions.GeoM.Translate(itemPosX, itemPosY)

			// Create the item
			items = append(items, &item.Item{
				Name:        fmt.Sprintf("item%d", i),
				Image:       itemImg,
				DrawOptions: drawOptions,
				PosX:        itemPosX,
				PosY:        itemPosY,
				Size:        float64(itemSize),
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
		HoverImage:  hoverImg,
		InPadding:   inPadding,
		OutPadding:  outPadding,
	}
}

func (i *Inventory) Draw(screen *ebiten.Image) {
	// Draw background image
	screen.DrawImage(i.Image, i.DrawOptions)

	// Draw each items
	oneIsHovered := false // avoids useless calls to "IsHovered" once one item is already hovered
	for _, it := range i.Items {
		if !oneIsHovered && it.IsHovered() {
			oneIsHovered = true
			hoverDrawOptions := &ebiten.DrawImageOptions{}
			hoverDrawOptions.GeoM.Translate(it.PosX-float64(i.InPadding), it.PosY-float64(i.InPadding))
			screen.DrawImage(i.HoverImage, hoverDrawOptions)
		}

		if it != nil {
			it.Draw(screen)
		}
	}

	return
}
