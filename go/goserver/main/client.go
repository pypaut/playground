package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"goserver/internal/client"
)

func main() {
	c := client.NewClient()

	ebiten.SetWindowSize(c.WinW, c.WinH)
	ebiten.SetWindowTitle("Client")

	if err := ebiten.RunGame(c); err != nil {
		log.Fatal(err)
	}
}
