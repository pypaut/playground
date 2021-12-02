package main

import (
	"github.com/go-vgo/robotgo"
)

var sx, sy = robotgo.GetScreenSize()
var screenCenterX = int(sx / 2)
var screenCenterY = int(sy / 2)

var potionX, potionY = screenCenterX, 0

func simpleClick() {
	robotgo.Toggle("left")
	robotgo.MilliSleep(100)
	robotgo.Toggle("left", "up")
	robotgo.MilliSleep(100)
}

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

func potionRappel() {
	robotgo.MoveSmooth(potionX, potionY+5)
	simpleClick()
	simpleClick()
}

func main() {
	potionRappel()
	return
}
