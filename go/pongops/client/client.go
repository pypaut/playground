package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"pongops/structs"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	playerDir := structs.Direction{
		X: 0,
		Y: 0,
	}

	var readBuf bytes.Buffer
	var writeBuf bytes.Buffer
	enc := gob.NewEncoder(&writeBuf)
	dec := gob.NewDecoder(&readBuf)

	var objects map[string]structs.Object

	for {
		// Wait one frame
		time.Sleep(1 / 60 * time.Second)

		// Encode direction
		if err = enc.Encode(playerDir); err != nil {
			fmt.Println("Error encoding player dir:", err)
			return
		}

		// Send direction
		n, err := fmt.Fprintf(conn, writeBuf.String()+"\n")
		//n, err := conn.Write(writeBuf.Bytes())
		if n == 0 {
			fmt.Println("No data written")
			continue
		}
		if err != nil {
			log.Fatal(err)
		}

		// Receive objects
		n, err = conn.Read(readBuf.Bytes())
		if n == 0 {
			fmt.Println("No data read")
			continue
		}
		if err != nil {
			log.Fatal(err)
		}

		// Decode objects
		if err = dec.Decode(&objects); err != nil {
			fmt.Println("Error decoding objects:", err)
			return
		}

		// Print objects
		fmt.Println(objects)
	}
}
