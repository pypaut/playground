package main

import (
	"github.com/go-vgo/robotgo"
)

var sx, sy = robotgo.GetScreenSize()
var screenCenterX = int(sx / 2)
var screenCenterY = int(sy / 2)

var potionX, potionY = screenCenterX, 5

func simpleClick() {
	robotgo.Toggle("left")
	robotgo.MilliSleep(100)
	robotgo.Toggle("left", "up")
	robotgo.MilliSleep(100)
}

func doubleClick() {
	simpleClick()
	simpleClick()
}

func dragTowards(xMove int, yMove int) {
	robotgo.Toggle("left")
	robotgo.MoveRelative(xMove, yMove)
	robotgo.Toggle("left", "up")
}

func changeMap(direction string) {
	dragLength := 200
	robotgo.Move(screenCenterX, screenCenterY)

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

	dragTowards(xMove, yMove)
}

func potionRappel() {
	robotgo.MoveSmooth(potionX, potionY)
	doubleClick()
}

func main() {
	potionRappel()
	// changeMap("DOWN")
	return
}
