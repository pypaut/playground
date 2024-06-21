package client

import (
	"bufio"
	"fmt"
	"image/color"
	"log"
	"math"
	"net"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/spf13/viper"

	"goserver/internal/parser"
)

type Client struct {
	WinW int
	WinH int

	playerPosX  float64
	playerPosY  float64
	playerSize  float64
	playerColor color.Color

	conn net.Conn
}

func NewClient() *Client {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	viper.SetConfigFile("goserver.yml")
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	return &Client{
		WinW:        viper.GetInt("window.width"),
		WinH:        viper.GetInt("window.height"),
		playerSize:  viper.GetFloat64("player.size"),
		playerColor: color.RGBA{150, 0, 150, 255},
		conn:        conn,
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
	playerPosStr, _ := bufio.NewReader(c.conn).ReadString('\n')
	playerPosStr = strings.TrimSuffix(playerPosStr, "\n")
	c.playerPosX, c.playerPosY, err = parser.ParseXandY(playerPosStr)
	if err != nil {
		return err
	}

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
	ebitenutil.DrawRect(
		screen, c.playerPosX, c.playerPosY, c.playerSize, c.playerSize, c.playerColor,
	)
	return
}

func (c *Client) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return c.WinW, c.WinH
}
