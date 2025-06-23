package main

import (
	"bufio"
	"fmt"
	"net"
	"testing"
)

func TestServer(t *testing.T) {
	go runServer()

	messageToSend := "my message hehe\n"

	conn, err := net.Dial("tcp", "localhost:"+PORT)
	if err != nil {
		t.Fatalf("Could not connect to the server: %v", err)
	}

	n, err := fmt.Fprint(conn, messageToSend)
	if n == 0 || err != nil {
		t.Fatalf("Could not send a message to the server: %v", err)
	}

	receivedMessage, err := bufio.NewReader(conn).ReadBytes('\n')
	if err != nil {
		t.Fatalf("Could not read a message from the server: %v", err)
	}

	if string(receivedMessage) != messageToSend {
		t.Fatalf("Unexpected message: %s", receivedMessage)
	}
}
