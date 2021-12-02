package main

import (
	"github.com/go-vgo/robotgo"
)

var sx, sy = robotgo.GetScreenSize()
var screenCenterX = int(sx / 2)
var screenCenterY = int(sy / 2)

func changeMap(direction string) {
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
	changeMap("UP")
	return
}
