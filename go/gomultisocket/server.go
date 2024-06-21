package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const (
	connHost = "localhost"
	connPort = "8080"
	connType = "tcp"
)

func main() {
	fmt.Println("Starting " + connType + " server on " + connHost + ":" + connPort)
	l, err := net.Listen(connType, connHost+":"+connPort)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer l.Close()

	clientId := 0

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error connecting:", err.Error())
			return
		}

		logMessage(clientId, conn, "Connected")

		go handleConnection(conn, clientId)
		clientId++
	}
}

func handleConnection(conn net.Conn, clientId int) {
	buffer, err := bufio.NewReader(conn).ReadBytes('\n')
	if err != nil {
		logMessage(clientId, conn, "Disconnected")
		conn.Close()
		return
	}

	clientMsg := string(buffer[:len(buffer)-1])
	logMessage(clientId, conn, fmt.Sprintf("Read \"%s\"", clientMsg))

	_, err = conn.Write(buffer)
	if err != nil {
		log.Fatal(err)
	}

	handleConnection(conn, clientId)
}

func logMessage(clientId int, conn net.Conn, message string) {
	log.Printf("[ID: %d, ADDR: %s]: %s\n", clientId, conn.RemoteAddr().String(), message)
}
