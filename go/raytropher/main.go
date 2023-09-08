package main

import (
	"fmt"
	"log"
	"raytropher/internal/scene"
)

func main() {
	// Load scene from conf
	myScene, err := scene.NewScene("scene.yaml")
	Check(err)

	fmt.Printf("%+v\n", myScene)

	// Render to image file
	err = myScene.RenderToImageFile("rendered.png")
	Check(err)

	return
}

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
