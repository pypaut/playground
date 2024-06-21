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

type PlayerPos struct {
	X float64
	Y float64
}

type Server struct {
	positions []*PlayerPos

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
		positions:   []*PlayerPos{},
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

	clientId := 0

	for {
		/* Accept connection from new client */
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		logMessage(clientId, conn, "Connected")
		s.positions = append(s.positions, &PlayerPos{X: 0, Y: 0})

		go s.handleConnection(conn, clientId)
		clientId++
	}
}

func (s *Server) handleConnection(conn net.Conn, clientId int) {
	for {
		/* Receive player's direction from client */
		playerDirStr, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			logMessage(clientId, conn, "Disconnected")
			conn.Close()
			return
		}
		playerDirStr = strings.TrimSuffix(playerDirStr, "\n")

		/* Update player's position */
		dirX, dirY, err := parser.ParseXandY(playerDirStr)
		if err != nil {
			log.Fatal(err)
		}

		s.positions[clientId].X += dirX * s.playerSpeed
		s.positions[clientId].Y += dirY * s.playerSpeed

		s.positions[clientId].X = clamp(s.positions[clientId].X, 0, float64(s.winH-s.playerSize))
		s.positions[clientId].Y = clamp(s.positions[clientId].Y, 0, float64(s.winW-s.playerSize))

		/* Send player's positions to client */
		posStr := fmt.Sprintf("x:%f,y:%f", s.positions[clientId].X, s.positions[clientId].Y)
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

func logMessage(clientId int, conn net.Conn, message string) {
	log.Printf("[ID: %d, ADDR: %s]: %s\n", clientId, conn.RemoteAddr().String(), message)
}
