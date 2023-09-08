package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Vec3 struct {
	X float64 `yaml:"x"`
	Y float64 `yaml:"y"`
	Z float64 `yaml:"z"`
}

type Sphere struct {
	Position Vec3    `yaml:"position"`
	Radius   float64 `yaml:"radius"`
}

type Camera struct {
	Position Vec3 `yaml:"position"`
	Forward  Vec3 `yaml:"forward"`
	Up       Vec3 `yaml:"up"`
}

type Scene struct {
	Sphere Sphere `yaml:"sphere"`
	Camera Camera `yaml:"camera"`
}

func main() {
	file, err := os.ReadFile("scene.yaml")
	Check(err)

	var scene Scene
	err = yaml.Unmarshal(file, &scene)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", scene)
	return
}

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
