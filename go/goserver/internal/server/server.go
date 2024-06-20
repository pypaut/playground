package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

type Server struct {
	conn net.Conn

	posX float64
	posY float64
}

func NewServer() *Server {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening on port 8000")

	conn, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}

	return &Server{
		conn: conn,
		posX: 0,
		posY: 0,
	}
}

func (s *Server) Serve() {
	for {
		/* Receive player's direction from client */
		msgFromClient, err := bufio.NewReader(s.conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print("Message from client: " + msgFromClient)
		msgFromClient = strings.TrimSuffix(msgFromClient, "\n")

		/* Update player's position */
		dirX, dirY, err := parseStrDir(msgFromClient)
		if err != nil {
			log.Fatal(err)
		}

		s.posX += dirX
		s.posY += dirY

		s.posX = clamp(s.posX, 0, 800)
		s.posY = clamp(s.posY, 0, 1000)

		/* Send player's positions to client */
		posStr := fmt.Sprintf("x:%f,y:%f", s.posX, s.posY)
		_, err = s.conn.Write([]byte(posStr + "\n"))
		if err != nil {
			return
		}
	}
}

func parseStrDir(msgFromClient string) (dirX, dirY float64, err error) {
	// Format: x:0.000000,y:0.000000
	dirs := strings.Split(msgFromClient, ",")
	dirXStr := strings.Split(dirs[0], ":")[1]
	dirYStr := strings.Split(dirs[1], ":")[1]

	dirX, err = strconv.ParseFloat(dirXStr, 8)
	if err != nil {
		return 0, 0, err
	}

	dirY, err = strconv.ParseFloat(dirYStr, 8)
	if err != nil {
		return 0, 0, err
	}

	return dirX, dirY, nil
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
