package main

import (
	"bufio"
	"net"
	"testing"
)

func TestServer(t *testing.T) {
	go runServer()

	// Can connect 2 clients
	conn1, err := net.Dial("tcp", "localhost:"+PORT)
	if err != nil {
		t.Fatalf("Could not connect to the server: %v", err)
	}

	conn2, err := net.Dial("tcp", "localhost:"+PORT)
	if err != nil {
		t.Fatalf("Could not connect to the server: %v", err)
	}

	// Receive ehlo from server
	reader1 := bufio.NewReader(conn1)
	ehlo1, err := reader1.ReadString('\n')
	if err != nil {
		t.Fatalf("Error when receiving ehlo1 from the server: %v", err)
	}
	if ehlo1 != "ehlo 1\n" {
		t.Fatalf("ehlo1 from the server is not ehlo: %s", ehlo1)
	}

	reader2 := bufio.NewReader(conn2)
	ehlo2, err := reader2.ReadString('\n')
	if err != nil {
		t.Fatalf("Error when receiving ehlo2 from the server: %v", err)
	}
	if ehlo2 != "ehlo 2\n" {
		t.Fatalf("ehlo2 from the server is not ehlo: %s", ehlo1)
	}

	// Can send players directions

	// Receives game objects
}
