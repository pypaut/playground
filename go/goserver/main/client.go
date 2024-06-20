package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"goserver/internal/client"
	"log"
)

func main() {
	c := client.NewClient()

	ebiten.SetWindowSize(c.WinW, c.WinH)
	ebiten.SetWindowTitle("Client")

	if err := ebiten.RunGame(c); err != nil {
		log.Fatal(err)
	}

}
