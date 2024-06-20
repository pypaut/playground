package client

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Client struct {
	Game *Game
}

func NewClient() *Client {
	g := NewGame()
	return &Client{Game: g}
}

func (c *Client) Run() {
	ebiten.SetWindowSize(c.Game.Width, c.Game.Height)
	ebiten.SetWindowTitle("Client")

	if err := ebiten.RunGame(c.Game); err != nil {
		log.Fatal(err)
	}

	return
}
