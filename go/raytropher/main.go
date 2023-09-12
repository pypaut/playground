package main

import (
	"log"
	"raytropher/internal/scene"
)

func main() {
	// Load scene from conf
	myScene, err := scene.NewScene("scene.yaml")
	Check(err)

	// fmt.Printf("%+v\n", myScene)

	// Render to image file
	err = myScene.RenderToFile("rendered.png")
	Check(err)

	return
}

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
