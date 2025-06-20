package main

import (
	"fmt"
	"io"
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
	defer conn.Close()

	for {
		// n, err := conn.Read(buf.Bytes())
		message, err := io.ReadAll(conn)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Received message '%s'\n", string(message))

		n, err := fmt.Fprintf(conn, "%s", message)
		if n == 0 {
			fmt.Println("Sent empty message")
		}
		if err != nil {
			log.Fatal(err)
		}
	}
}

// Init game, with objects
// Accept connections (2)
//

/*
func main() {
	gameObjects := CreateGameObjects()

	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening on port 8000")

	conn1, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Player 1 is connected")
	go handlePlayer(conn1, gameObjects["player1"], 1, gameObjects)

	conn2, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Player 2 is connected")
	go handlePlayer(conn2, gameObjects["player2"], 2, gameObjects)

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	for {
		// Wait one frame
		time.Sleep(1 / 60 * time.Second)

		// Update each object according to direction, collision
		Update(gameObjects)

		if err = enc.Encode(gameObjects); err != nil {
			fmt.Println("Error encoding game objects:", err)
			return
		}

		_, err = conn1.Write(buf.Bytes())
		if err != nil {
			return
		}

		_, err = conn1.Write(buf.Bytes())
		if err != nil {
			return
		}
	}
}

func handlePlayer(conn net.Conn, player *structs.Object, playerID int, objects map[string]*structs.Object) {
	// var readBuf bytes.Buffer
	var playerDir structs.Direction

	var writeBuf bytes.Buffer
	enc := gob.NewEncoder(&writeBuf)

	for {
		// Wait one frame
		time.Sleep(1 / 60 * time.Second)

		err := receivePlayerDir(conn, &playerDir)
		if err != nil {
			fmt.Printf("Player %d disconnected\n", playerID)
		}

		// Update player object with direction
		player.DirX, player.DirY = playerDir.X, playerDir.Y

		if err = enc.Encode(objects); err != nil {
			fmt.Println("Error encoding encoding objects:", err)
			return
		}

		n, err := fmt.Fprintf(conn, writeBuf.String()+"\n")
		if n == 0 {
			fmt.Println("No data written")
		}
		if err != nil {
			fmt.Println("Error sending data:", err)
		}
	}
}

func receivePlayerDir(conn net.Conn, playerDir *structs.Direction) (err error) {
	/* Receive player's direction from client */
// n, err := conn.Read(readBuf.Bytes())
// if n == 0 {
// 	fmt.Printf("Player %d disconnected\n", playerID)
// 	return
// }
// if err != nil {
// 	log.Fatal(err)
// }

/*
	playerDirByte, err := bufio.NewReader(conn).ReadBytes('\n')
	if err != nil {
		return
	}

	dec := gob.NewDecoder(bytes.NewBuffer(playerDirByte))
	// dec := gob.NewDecoder(&readBuf)
	if err = dec.Decode(&playerDir); err != nil {
		fmt.Println("Error decoding struct:", err)
		return
	}

	return nil
}

func Update(gameObjects map[string]*structs.Object) {
	for _, obj := range gameObjects {
		obj.PosX += obj.DirX
		obj.PosY += obj.DirY
	}

	return
}

func CreateGameObjects() map[string]*structs.Object {
	player1 := structs.Object{
		PosX:   200,
		PosY:   0,
		Width:  50,
		Height: 300,
	}

	player2 := structs.Object{
		PosX:   structs.WinW - 200,
		PosY:   0,
		Width:  50,
		Height: 300,
	}

	ball := structs.Object{
		PosX:   structs.WinW / 2,
		PosY:   structs.WinH / 2,
		Width:  20,
		Height: 20,
	}

	return map[string]*structs.Object{
		"player1": &player1,
		"player2": &player2,
		"ball":    &ball,
	}
}

*/
