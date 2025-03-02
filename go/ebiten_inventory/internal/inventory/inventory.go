package inventory

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/color"
	"inventory/internal/item"
	"math/rand/v2"
)

type Inventory struct {
	Image       *ebiten.Image
	DrawOptions *ebiten.DrawImageOptions

	HoverImage     *ebiten.Image
	HoverThickness int
	Items          []*item.Item

	InPadding  int
	OutPadding int

	DraggedItem *item.Item
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
		width+int(outPadding)*2+inPadding*(widthItems-1),
		height+int(outPadding)*2+inPadding*(heightItems-1),
	)
	img.Fill(color.White)

	// Hover image
	hoverThickness := inPadding - 2
	hoverImg := ebiten.NewImage(itemSize+2*hoverThickness, itemSize+2*hoverThickness)
	hoverImg.Fill(color.RGBA{R: 200, A: 255})

	var items []*item.Item

	for i := 0; i < widthItems; i++ {
		for j := 0; j < heightItems; j++ {
			// Build random image
			itemImg := ebiten.NewImage(itemSize, itemSize)
			itemImg.Fill(color.Gray{Y: uint8(rand.IntN(200))})

			// Setup draw options for position
			drawOptions := &ebiten.DrawImageOptions{}
			itemPosX := invPosX + float64(i*itemSize+outPadding+inPadding*i)
			itemPosY := invPosY + float64(j*itemSize+outPadding+inPadding*j)
			drawOptions.GeoM.Translate(itemPosX, itemPosY)

			// Create the item
			items = append(items, &item.Item{
				Name:        fmt.Sprintf("item%d", i),
				Image:       itemImg,
				DrawOptions: drawOptions,
				PosX:        int(itemPosX),
				PosY:        int(itemPosY),
				Size:        itemSize,
			})
		}
	}

	drawOptions := ebiten.DrawImageOptions{}
	drawOptions.GeoM.Translate(invPosX, invPosY)

	return &Inventory{
		Image:       img,
		DrawOptions: &drawOptions,

		HoverImage:     hoverImg,
		HoverThickness: hoverThickness,
		Items:          items,

		InPadding:  inPadding,
		OutPadding: outPadding,
	}
}

func (i *Inventory) Draw(screen *ebiten.Image) {
	// Draw background image
	screen.DrawImage(i.Image, i.DrawOptions)

	// Draw each items
	for _, it := range i.Items {
		if inpututil.IsMouseButtonJustReleased(ebiten.MouseButton0) {
			i.DraggedItem = nil
		} else if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
			i.DraggedItem = it
		}

		if it.IsHovered() {
			hoverDrawOptions := &ebiten.DrawImageOptions{}
			hoverDrawOptions.GeoM.Translate(
				float64(it.PosX-i.HoverThickness),
				float64(it.PosY-i.HoverThickness),
			)

			screen.DrawImage(i.HoverImage, hoverDrawOptions)
		}

		if it != nil {
			it.Draw(screen)
		}

		if i.DraggedItem != nil {
			drawOptions := &ebiten.DrawImageOptions{}
			mouseX, mouseY := ebiten.CursorPosition()
			drawOptions.GeoM.Translate(float64(mouseX), float64(mouseY))
			screen.DrawImage(i.DraggedItem.Image, drawOptions)
		}
	}

	return
}
