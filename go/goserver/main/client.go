package main

import "goserver/internal/client"

func main() {
	c := client.NewClient()
	c.Run()
	return
}
