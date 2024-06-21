package parser

import (
	"strconv"
	"strings"
)

func ParseXandY(msg string) (x, y float64, err error) {
	// Format: x:0.000000,y:0.000000
	dirs := strings.Split(msg, ",")
	xStr := strings.Split(dirs[0], ":")[1]
	yStr := strings.Split(dirs[1], ":")[1]

	x, err = strconv.ParseFloat(xStr, 8)
	if err != nil {
		return 0, 0, err
	}

	y, err = strconv.ParseFloat(yStr, 8)
	if err != nil {
		return 0, 0, err
	}

	return x, y, nil
}
