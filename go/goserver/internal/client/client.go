package client

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"net"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/spf13/viper"
)

type Client struct {
	WinW int
	WinH int

	conn net.Conn
}

func NewClient() *Client {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	viper.SetConfigFile("../goserver.yml")
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	return &Client{
		WinW: viper.GetInt("window.width"),
		WinH: viper.GetInt("window.height"),
		conn: conn,
	}
}

func (c *Client) Update() error {
	/* Compute direction */
	dirX, dirY := computePlayerDir()
	dirStr := fmt.Sprintf("x:%f,y:%f\n", dirX, dirY)

	/* Send player's direction to server */
	_, err := fmt.Fprintf(c.conn, dirStr+"\n")
	if err != nil {
		return err
	}

	/* Get player's position from server */
	message, _ := bufio.NewReader(c.conn).ReadString('\n')
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

func (c *Client) Draw(screen *ebiten.Image) {
	return
}

func (c *Client) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return c.WinW, c.WinH
}
