package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/spf13/viper"

	"goserver/internal/parser"
)

type Server struct {
	posX float64
	posY float64

	winW int
	winH int

	playerSize  int
	playerSpeed float64
}

func NewServer() *Server {
	viper.SetConfigFile("goserver.yml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	return &Server{
		posX:        0,
		posY:        0,
		winW:        viper.GetInt("window.width"),
		winH:        viper.GetInt("window.height"),
		playerSize:  viper.GetInt("player.size"),
		playerSpeed: viper.GetFloat64("player.speed"),
	}
}

func (s *Server) Serve() {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening on port 8000")

	conn, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}

	for {
		/* Receive player's direction from client */
		playerDirStr, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print("Message from client: " + playerDirStr)
		playerDirStr = strings.TrimSuffix(playerDirStr, "\n")

		/* Update player's position */
		dirX, dirY, err := parser.ParseXandY(playerDirStr)
		if err != nil {
			log.Fatal(err)
		}

		s.posX += dirX * s.playerSpeed
		s.posY += dirY * s.playerSpeed

		s.posX = clamp(s.posX, 0, float64(s.winH-s.playerSize))
		s.posY = clamp(s.posY, 0, float64(s.winW-s.playerSize))

		/* Send player's positions to client */
		posStr := fmt.Sprintf("x:%f,y:%f", s.posX, s.posY)
		_, err = conn.Write([]byte(posStr + "\n"))
		if err != nil {
			return
		}
	}
}

func clamp(value, min, max float64) float64 {
	if value < min {
		return min
	}

	if value > max {
		return max
	}

	return value
}
