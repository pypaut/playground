package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type ClickDir string

const (
	ClickDirLeft  ClickDir = "L"
	ClickDirRight ClickDir = "R"
)

func main() {
	landedOnZero, passedByZero := Compute("input")
	fmt.Println(landedOnZero)
	fmt.Println(passedByZero)
	return
}

func Compute(path string) (landedOnZero int, passedByZero int) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	currentPos := 50
	lastPassedByZero := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		l := scanner.Text()

		clickDir := l[:1]
		nbClicks, err := strconv.Atoi(l[1:])
		if err != nil {
			panic(err)
		}

		currentPos, lastPassedByZero = ComputeNextPos(currentPos, ClickDir(clickDir), nbClicks)

		if currentPos == 0 {
			landedOnZero += 1
		}

		passedByZero += lastPassedByZero
	}

	return
}

func ComputeNextPos(currentPos int, clickDir ClickDir, nbClicks int) (nextPos int, passagesByZero int) {
	nextPos = currentPos

	switch clickDir {
	case ClickDirLeft:
		nextPos = currentPos - nbClicks
	case ClickDirRight:
		nextPos = currentPos + nbClicks
	default:
		panic(errors.New("wrong left right"))
	}

	// I don't get why but it works
	tmpCurrent := nextPos
	supp := 0
	if tmpCurrent < 0 {
		tmpCurrent *= -1
		if currentPos != 0 {
			supp = 1
		}
	} else if tmpCurrent == 0 && currentPos > 0 {
		supp = 1
	}

	passagesByZero = tmpCurrent/100 + supp

	nextPos %= 100
	if nextPos < 0 {
		nextPos += 100
	}

	return nextPos, passagesByZero
}
