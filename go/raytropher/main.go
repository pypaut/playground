package main

import (
	"fmt"
	"log"
	"raytropher/internal/scene"
)

func main() {
	myScene, err := scene.NewScene("scene.yaml")
	Check(err)

	fmt.Printf("%+v\n", myScene)
	return
}

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
