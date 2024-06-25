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
	"goserver/internal/player"
)

type Client struct {
	WinW int
	WinH int

	playerPosX  float64
	playerPosY  float64
	players     []*player.Player
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

	/* Get players' positions from server */
	playersPosStr, _ := bufio.NewReader(c.conn).ReadString('\n')
	playersPosStr = strings.TrimSuffix(playersPosStr, "\n")

	c.players = []*player.Player{}
	playersPosSplit := strings.Split(playersPosStr, ";")
	for _, playerPosStr := range playersPosSplit {
		if playerPosStr == "" {
			continue
		}
		x, y, err := parser.ParseXandY(playerPosStr)
		if err != nil {
			return err
		}

		c.players = append(c.players, &player.Player{X: x, Y: y})
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
	for _, p := range c.players {
		ebitenutil.DrawRect(
			screen, p.X, p.Y, c.playerSize, c.playerSize, c.playerColor,
		)
	}
	// ebitenutil.DrawRect(
	// 	screen, c.playerPosX, c.playerPosY, c.playerSize, c.playerSize, c.playerColor,
	// )
	return
}

func (c *Client) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return c.WinW, c.WinH
}
