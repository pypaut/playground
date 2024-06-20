package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type Server struct {
	conn net.Conn
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

		/* Update player's position */
		posX, posY := 0.0, 0.0

		/* Send player's positions to client */
		posStr := fmt.Sprintf("x:%f,y:%f", posX, posY)
		_, err = s.conn.Write([]byte(posStr + "\n"))
		if err != nil {
			return
		}
	}
}
