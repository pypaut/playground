package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	CurrentMap *Map
	CameraPos  Point
}

func (g *Game) Update() error {
	direction := Point{}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		direction.X -= CameraSpeed
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		direction.X += CameraSpeed
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		direction.Y -= CameraSpeed
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		direction.Y += CameraSpeed
	}

	direction.Normalize()

	g.CameraPos.X += direction.X
	g.CameraPos.Y += direction.Y

	return nil
}

func (g *Game) isoToScreen(col, row int) (float64, float64) {
	// Standard isometric projection math:
	// X moves right for columns, left for rows.
	// Y moves down for both columns and rows.
	x := float64(col-row)*tileWidth/2 - g.CameraPos.X
	y := float64(col+row)*tileHeight/2 - g.CameraPos.Y

	return x, y
}

func (g *Game) drawGrid(screen *ebiten.Image) {
	var gridColor = color.RGBA{100, 100, 100, 255}
	var strokeWidth float32 = 1.0

	// Columns
	for i := range int(g.CurrentMap.Size.X) {
		x1, y1 := g.isoToScreen(i, 0)
		x2, y2 := g.isoToScreen(i, int(g.CurrentMap.Size.X)-1)
		vector.StrokeLine(
			screen,
			float32(x1), float32(y1),
			float32(x2), float32(y2),
			strokeWidth,
			gridColor,
			true,
		)
	}

	// Lines
	for j := range int(g.CurrentMap.Size.Y) {
		x1, y1 := g.isoToScreen(0, j)
		x2, y2 := g.isoToScreen(int(g.CurrentMap.Size.Y)-1, j)
		vector.StrokeLine(
			screen,
			float32(x1), float32(y1),
			float32(x2), float32(y2),
			strokeWidth,
			gridColor,
			true,
		)
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	// Draw map tiles
	opts := &ebiten.DrawImageOptions{}
	tileImg := ebiten.NewImage(int(TileSize), int(TileSize))
	tileImg.Fill(color.RGBA{R: 100, A: 255})

	for _, p := range myMap.WalkableTiles {
		opts.GeoM.Reset()
		opts.GeoM.Rotate(0.79) // Approx PI/4
		opts.GeoM.Scale(1, 0.5)

		x, y := g.isoToScreen(int(p.X), int(p.Y))

		opts.GeoM.Translate(x, y)
		screen.DrawImage(tileImg, opts)
	}

	g.drawGrid(screen)
}

func (g *Game) Layout(outW, outH int) (int, int) {
	return screenW, screenH
}
