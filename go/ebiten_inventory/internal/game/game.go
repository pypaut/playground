package game

import (
	"errors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"inventory/internal/inventory"
)

type Game struct {
	IsInventoryOpened bool
	Inventory         *inventory.Inventory
}

func NewGame() *Game {
	return &Game{
		Inventory: inventory.NewInventory(),
	}
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		return errors.New("quitting the game")
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyI) {
		if !g.IsInventoryOpened {
			g.IsInventoryOpened = true
		} else {
			g.IsInventoryOpened = false
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.IsInventoryOpened {
		g.Inventory.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1920, 1080
}
