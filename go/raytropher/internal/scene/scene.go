package scene

import (
	"gopkg.in/yaml.v3"
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

func NewScene(sceneFile string) (scene *Scene, err error) {
	file, err := os.ReadFile("scene.yaml")
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(file, &scene)
	if err != nil {
		return nil, err
	}

	return scene, nil
}
