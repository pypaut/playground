package main

import (
	"goserver/internal/server"
)

func main() {
	s := server.NewServer()
	s.Serve()
}
