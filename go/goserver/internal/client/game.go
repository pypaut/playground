package client

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"net"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Width  int
	Height int

	conn net.Conn
}

func NewGame() *Game {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	return &Game{
		Width:  1000,
		Height: 800,
		conn:   conn,
	}
}

func (g *Game) Update() error {
	/* Compute direction */
	dirX, dirY := computePlayerDir()
	dirStr := fmt.Sprintf("x:%f,y:%f\n", dirX, dirY)

	/* Send message */
	_, err := fmt.Fprintf(g.conn, dirStr+"\n")
	if err != nil {
		return err
	}

	/* Read message */
	message, _ := bufio.NewReader(g.conn).ReadString('\n')
	fmt.Print("Message from server: " + message)

	return nil
}

func computePlayerDir() (dirX, dirY float64) {

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		dirY--
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		dirX--
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		dirY++
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		dirX++
	}

	return normalized(dirX, dirY)
}

func normalized(dirX, dirY float64) (float64, float64) {
	if dirX == 0 && dirY == 0 {
		return 0, 0
	}

	norm := math.Sqrt(math.Pow(dirX, 2) + math.Pow(dirY, 2))
	return dirX / norm, dirY / norm
}

func (g *Game) Draw(screen *ebiten.Image) {
	return
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Width, g.Height
}
