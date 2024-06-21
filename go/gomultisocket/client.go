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
	fmt.Println("Connecting to " + connType + " server " + connHost + ":" + connPort)

	conn, err := net.Dial(connType, connHost+":"+connPort)
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}
	reader := bufio.NewReader(os.Stdin)

	for {
		/* Send */
		fmt.Print("Text to send: ")
		input, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, input+"\n")

		/* Read */
		message, _ := bufio.NewReader(conn).ReadString('\n')
		log.Printf("Server relay: \"%s\"\n", message[:len(message)-1])

	}
}
