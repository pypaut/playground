package scene

import (
	"gopkg.in/yaml.v3"
	"image"
	"image/color"
	"image/draw"
	"image/png"
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
	Position       Vec3    `yaml:"position"`
	Forward        Vec3    `yaml:"forward"`
	Up             Vec3    `yaml:"up"`
	ScreenWidth    int     `yaml:"screen_width"`
	ScreenHeight   int     `yaml:"screen_height"`
	ScreenDistance float64 `yaml:"screen_distance"`
}

type Scene struct {
	Sphere Sphere `yaml:"sphere"`
	Camera Camera `yaml:"camera"`
}

func NewScene(sceneFile string) (scene *Scene, err error) {
	file, err := os.ReadFile(sceneFile)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(file, &scene)
	if err != nil {
		return nil, err
	}

	return scene, nil
}

func (s *Scene) RenderToImageFile(imageName string) (err error) {
	// Init image
	myImg := image.NewRGBA(image.Rect(0, 0, s.Camera.ScreenWidth, s.Camera.ScreenHeight))
	black := color.RGBA{R: 0, G: 0, B: 0, A: 255}
	draw.Draw(myImg, myImg.Bounds(), &image.Uniform{C: black}, image.Point{}, draw.Src)

	// TODO: Render image using raytracing

	for y := 0; y < s.Camera.ScreenHeight; y++ {
		for x := 0; x < s.Camera.ScreenWidth; x++ {
			myImg.Set(x, y, color.NRGBA{
				R: uint8((x + y) & 255),
				G: uint8((x + y) << 1 & 255),
				B: uint8((x + y) << 2 & 255),
				A: 255,
			})
		}
	}

	// Export image file
	out, err := os.Create(imageName)
	if err != nil {
		return err
	}

	err = png.Encode(out, myImg)
	if err != nil {
		return err
	}

	out.Close()
	return nil
}
