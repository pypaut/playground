package client

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Client struct {
	game *Game
}

func NewClient() *Client {
	g := NewGame()
	return &Client{game: g}
}

func (c *Client) Run() {
	ebiten.SetWindowSize(c.game.Width, c.game.Height)
	ebiten.SetWindowTitle("Client")

	if err := ebiten.RunGame(c.game); err != nil {
		log.Fatal(err)
	}

	return
}
