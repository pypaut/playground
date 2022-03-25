package game

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type Game struct {
	Window   *sdl.Window
	Renderer *sdl.Renderer
}

func NewRenderer(window *sdl.Window) **sdl.Renderer {
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("Renderer initialization:", err)
		return nil
	}
	defer func(renderer *sdl.Renderer) {
		err := renderer.Destroy()
		if err != nil {
			return
		}
	}(renderer)

	return &renderer
}

func NewGame(screenWidth, screenHeight int32, playerSpeed, ballSpeed float32) (*Game, error) {

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("SDL initialization:", err)
		return nil, err
	}

	window, err := sdl.CreateWindow(
		"GoPong",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		screenWidth,
		screenHeight,
		sdl.WINDOW_OPENGL,
	)
	if err != nil {
		fmt.Println("Window initialization:", err)
		return nil, err
	}
	defer func(window *sdl.Window) {
		err := window.Destroy()
		if err != nil {
			return
		}
	}(window)

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("Renderer initialization:", err)
		return nil, err
	}
	defer func(renderer *sdl.Renderer) {
		err := renderer.Destroy()
		if err != nil {
			return
		}
	}(renderer)


	return &Game{
		Window: window,
		Renderer: renderer,
	}, nil

}


func (g *Game) Draw(p1Rect, p2Rect, ballRect sdl.Rect) error {

	err := g.Renderer.SetDrawColor(0, 0, 0, 255)
	if err != nil {
		fmt.Println("Background Renderer.SetDrawColor:", err)
		return err
	}
	err = g.Renderer.Clear()
	if err != nil {
		fmt.Println("Background Renderer.Clear:", err)
		return err
	}

	// Player 1
	err = g.Renderer.SetDrawColor(150, 0, 150, 255)
	if err != nil {
		fmt.Println("Player 1 Renderer.SetDrawColor:", err)
		return err
	}
	err = g.Renderer.FillRect(&p1Rect)
	if err != nil {
		fmt.Println("Player 1 Renderer.FillRect:", err)
		return err
	}

	// Player 2
	err = g.Renderer.SetDrawColor(150, 0, 150, 255)
	if err != nil {
		fmt.Println("Player 2 Renderer.SetDrawColor:", err)
		return err
	}
	err = g.Renderer.FillRect(&p2Rect)
	if err != nil {
		fmt.Println("Player 2 Renderer.FillRect:", err)
		return err
	}

	// Ball
	err = g.Renderer.SetDrawColor(150, 0, 150, 255)
	if err != nil {
		fmt.Println("Ball Renderer.SetDrawColor:", err)
		return err
	}
	err = g.Renderer.FillRect(&ballRect)
	if err != nil {
		fmt.Println("Ball Renderer.FillRect:", err)
		return err
	}

	g.Renderer.Present()

	return nil
}