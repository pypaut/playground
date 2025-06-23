package main

import (
	"fmt"
	"log"
	"net"
)

const (
	PORT = "8086"
	WinW = 1920
	WinH = 1080
)

type Direction struct {
	X, Y float64
}

type GameObject struct {
	PosX, PosY    float64
	DirX, DirY    float64
	Width, Height float64
}

type GameData struct {
	Player1 *GameObject
	Player2 *GameObject
	Ball    *GameObject
}

func main() {
	runServer()
}

func initGameData() *GameData {

	player1 := &GameObject{
		PosX:   200,
		PosY:   0,
		Width:  50,
		Height: 300,
	}

	player2 := &GameObject{
		PosX:   WinW - 200,
		PosY:   0,
		Width:  50,
		Height: 300,
	}

	ball := &GameObject{
		PosX:   WinW / 2,
		PosY:   WinH / 2,
		Width:  20,
		Height: 20,
	}

	return &GameData{Player1: player1, Player2: player2, Ball: ball}
}

func runServer() {
	// gd := initGameData()

	ln, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening on port " + PORT)

	conn1, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}
	// defer CloseConn(conn1) // WARNING this closes the connection too early
	go handleConnection(conn1, 1)

	conn2, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}
	// defer CloseConn(conn2) // WARNING this closes the connection too early
	go handleConnection(conn2, 2)
}

func handleConnection(conn net.Conn, id int) {
	// Send ehlo
	n, err := fmt.Fprintf(conn, "ehlo %d\n", id)
	if n == 0 {
		fmt.Println("Sent empty message")
	}
	if err != nil {
		log.Fatal(err)
	}

}

func CloseConn(conn net.Conn) {
	err := conn.Close()
	if err != nil {
		log.Fatal(err)
	}
}
