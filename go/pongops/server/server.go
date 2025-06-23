package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const PORT = "8086"

func main() {
	runServer()
}

func runServer() {
	ln, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening on port " + PORT)

	conn, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)

	for {
		/* Receive player's direction from client */
		message, err := bufio.NewReader(conn).ReadBytes('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Received message '%s'\n", string(message))

		n, err := fmt.Fprintf(conn, "%s", string(message))
		if n == 0 {
			fmt.Println("Sent empty message")
		}
		if err != nil {
			log.Fatal(err)
		}
	}
}
