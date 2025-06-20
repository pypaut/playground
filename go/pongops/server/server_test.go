package main

import (
	"fmt"
	"io"
	"net"
	"testing"
)

// func TestReceivePlayerDir(t *testing.T) {
// 	// Create a server
// 	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintln(w, "Hello, client")
// 	}))
// 	defer ts.Close()
//
// 	// Create a client
// 	client := ts.Client()
// 	res, err := client.Get(ts.URL)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	greeting, err := io.ReadAll(res.Body)
// 	res.Body.Close()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	// fmt.Printf("%s", greeting)
// 	if string(greeting) != "Hello, client\n" {
// 		t.Fatalf("Expected '%s', got '%s'", "Hello, client", string(greeting))
// 	}
// }

func TestServer(t *testing.T) {
	go runServer()

	conn, err := net.Dial("tcp", "localhost:"+PORT)
	if err != nil {
		t.Fatalf("Could not connect to the server: %v", err)
	}

	n, err := fmt.Fprintf(conn, "hello world")
	if n == 0 || err != nil {
		t.Fatalf("Could not send a message to the server: %v", err)
	}

	message, err := io.ReadAll(conn)
	if err != nil {
		t.Fatalf("Could not read a message from the server: %v", err)
	}

	if string(message) != "hello world" {
		t.Fatalf("Unexpected message: %s", message)
	}
}
