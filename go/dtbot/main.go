package main

import (
	"github.com/go-vgo/robotgo"
)

func changeMap(direction string, screenCenterX int, screenCenterY int) {
	robotgo.Move(screenCenterX, screenCenterY)
	dragLength := 200
	xMove := 0
	yMove := 0
	switch direction {
	case "UP":
		xMove, yMove = 0, dragLength
	case "DOWN":
		xMove, yMove = 0, -dragLength
	case "LEFT":
		xMove, yMove = -dragLength, 0
	case "RIGHT":
		xMove, yMove = dragLength, 0
	default:
		xMove, yMove = 0, 0
	}

	robotgo.Toggle("left")
	robotgo.MoveRelative(xMove, yMove)
	robotgo.Toggle("left", "up")
}

func main() {
	sx, sy := robotgo.GetScreenSize()
	centerX := int(sx / 2)
	centerY := int(sy / 2)
	changeMap("UP", centerX, centerY)
	return
}
