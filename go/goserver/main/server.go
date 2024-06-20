package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
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
		msgFromClient, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print("Message from client: " + msgFromClient)

		posX, posY := 0.0, 0.0
		posStr := fmt.Sprintf("x:%f,y:%f", posX, posY)
		_, err = conn.Write([]byte(posStr + "\n"))
		if err != nil {
			return
		}
	}
}
