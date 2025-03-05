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
	GamepadEnabled    bool
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

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
		g.GamepadEnabled = false
		ebiten.SetCursorMode(ebiten.CursorModeVisible)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyW) ||
		inpututil.IsKeyJustPressed(ebiten.KeyA) ||
		inpututil.IsKeyJustPressed(ebiten.KeyS) ||
		inpututil.IsKeyJustPressed(ebiten.KeyD) {
		g.GamepadEnabled = true
		ebiten.SetCursorMode(ebiten.CursorModeHidden)
	}

	if err := g.Inventory.Update(g.GamepadEnabled); err != nil {
		return err
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
