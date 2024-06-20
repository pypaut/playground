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
	dir_x := 0.0
	dir_y := 0.0

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		dir_y--
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		dir_x--
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		dir_y++
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		dir_x++
	}

	dir_x, dir_y = normalized(dir_x, dir_y)
	fmt.Printf("x:%f,y:%f\n", dir_x, dir_y)

	/* Server communication */
	text := "message from client"
	fmt.Fprintf(g.conn, text+"\n")
	message, _ := bufio.NewReader(g.conn).ReadString('\n')
	fmt.Print("Message from server: " + message)

	return nil
}

func normalized(dir_x, dir_y float64) (float64, float64) {
	if dir_x == 0 && dir_y == 0 {
		return 0, 0
	}

	norm := math.Sqrt(math.Pow(dir_x, 2) + math.Pow(dir_y, 2))
	return dir_x / norm, dir_y / norm
}

func (g *Game) Draw(screen *ebiten.Image) {
	return
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Width, g.Height
}
