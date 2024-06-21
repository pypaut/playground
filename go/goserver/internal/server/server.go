package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/spf13/viper"

	"goserver/internal/parser"
	"goserver/internal/player"
)

type Server struct {
	players []*player.Player

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
		players:     []*player.Player{},
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
		s.players = append(s.players, &player.Player{X: 0, Y: 0})

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

		s.players[clientId].X += dirX * s.playerSpeed
		s.players[clientId].Y += dirY * s.playerSpeed

		s.players[clientId].X = clamp(
			s.players[clientId].X, 0, float64(s.winH-s.playerSize),
		)
		s.players[clientId].Y = clamp(
			s.players[clientId].Y, 0, float64(s.winW-s.playerSize),
		)

		/* Send players' players to client */

		posStr := ""
		for _, pos := range s.players {
			posStr += fmt.Sprintf(";x:%f,y:%f", pos.X, pos.Y)
		}

		// posStr := fmt.Sprintf(
		// 	"x:%f,y:%f", s.players[clientId].X, s.players[clientId].Y,
		// )

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
