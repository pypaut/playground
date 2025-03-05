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

	InPadding  int
	OutPadding int

	HoverImage       *ebiten.Image
	HoverThickness   int
	HoverDrawOptions *ebiten.DrawImageOptions
	HoverIndex       int
	IsHovered        bool

	SelectedItem            *item.Item
	SelectedItemDrawOptions *ebiten.DrawImageOptions
	Items                   []*item.Item
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

		InPadding:  inPadding,
		OutPadding: outPadding,

		HoverImage:       hoverImg,
		HoverThickness:   hoverThickness,
		HoverDrawOptions: &ebiten.DrawImageOptions{},
		HoverIndex:       -1,

		SelectedItemDrawOptions: &ebiten.DrawImageOptions{},

		Items: items,
	}
}

func (i *Inventory) Update(gamepadEnabled bool) error {
	if !gamepadEnabled {
		i.IsHovered = false

		for index, it := range i.Items {
			// Check hovered item
			if it.IsHovered() {
				i.HoverDrawOptions.GeoM.Reset()
				i.HoverDrawOptions.GeoM.Translate(
					float64(it.PosX-i.HoverThickness),
					float64(it.PosY-i.HoverThickness),
				)

				i.IsHovered = true
				i.HoverIndex = index
				break
			}
		}

		// Check if an item is selected
		if inpututil.IsMouseButtonJustReleased(ebiten.MouseButton0) {
			i.SelectedItem = nil
		} else if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) && i.IsHovered {
			i.SelectedItem = i.Items[i.HoverIndex]
		}
	} else {
		i.IsHovered = true

		if inpututil.IsKeyJustPressed(ebiten.KeyD) {
			if i.HoverIndex < len(i.Items)-1 {
				i.HoverIndex++
			}
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyA) {
			if i.HoverIndex > 0 {
				i.HoverIndex--
			}
		}

		i.HoverDrawOptions.GeoM.Reset()
		i.HoverDrawOptions.GeoM.Translate(
			float64(i.Items[i.HoverIndex].PosX-i.HoverThickness),
			float64(i.Items[i.HoverIndex].PosY-i.HoverThickness),
		)
	}

	if i.SelectedItem != nil {
		if !gamepadEnabled {
			mouseX, mouseY := ebiten.CursorPosition()
			i.SelectedItemDrawOptions.GeoM.Reset()
			i.SelectedItemDrawOptions.GeoM.Translate(float64(mouseX), float64(mouseY))
		}
	}

	return nil
}

func (i *Inventory) Draw(screen *ebiten.Image) {
	// Draw background image
	screen.DrawImage(i.Image, i.DrawOptions)

	// Draw hover image
	if i.IsHovered {
		screen.DrawImage(i.HoverImage, i.HoverDrawOptions)
	}

	// Draw each items
	for _, it := range i.Items {
		if it != nil {
			it.Draw(screen)
		}
	}

	// Draw selected item (while dragging for example)
	if i.SelectedItem != nil {
		screen.DrawImage(i.SelectedItem.Image, i.SelectedItemDrawOptions)
	}

	return
}
