package scene

import (
	"gopkg.in/yaml.v3"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	m "raytropher/internal/math"
	"raytropher/internal/sphere"
)

type Camera struct {
	Position       m.Vec3  `yaml:"position"`
	Forward        m.Vec3  `yaml:"forward"`
	Up             m.Vec3  `yaml:"up"`
	ScreenWidth    int     `yaml:"screen_width"`
	ScreenHeight   int     `yaml:"screen_height"`
	ScreenDistance float64 `yaml:"screen_distance"`
}

type Scene struct {
	Sphere sphere.Sphere `yaml:"sphere"`
	Camera Camera        `yaml:"camera"`
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

func (s *Scene) RenderToFile(imageName string) (err error) {
	img, err := s.RenderToImage()
	if err != nil {
		return err
	}

	// Export image file
	out, err := os.Create(imageName)
	if err != nil {
		return err
	}

	err = png.Encode(out, img)
	if err != nil {
		return err
	}

	err = out.Close()
	if err != nil {
		return err
	}

	return nil
}

func (s *Scene) RenderToImage() (img draw.Image, err error) {
	// Init image
	img = image.NewRGBA(image.Rect(0, 0, s.Camera.ScreenWidth, s.Camera.ScreenHeight))
	black := color.RGBA{R: 0, G: 0, B: 0, A: 255}
	draw.Draw(img, img.Bounds(), &image.Uniform{C: black}, image.Point{}, draw.Src)

	pixelSize := 0.01

	// For each pixel of the screen
	for y := 0; y < s.Camera.ScreenHeight; y++ {
		for x := 0; x < s.Camera.ScreenWidth; x++ {
			// Compute ray direction
			var screenPixelPosition m.Vec3
			screenPixelPosition.Add(s.Camera.Position)
			screenPixelPosition.Add(s.Camera.Forward.Times(s.Camera.ScreenDistance))
			screenPixelPosition.Add(s.Camera.Up.Times(float64(s.Camera.ScreenHeight/2-y) * pixelSize))
			left := s.Camera.Up.Cross(s.Camera.Forward)
			screenPixelPosition.Add(left.Times(float64(s.Camera.ScreenWidth/2+x) * pixelSize))
			rayDirection := screenPixelPosition.Minus(s.Camera.Position)

			intersects, _ := s.Sphere.IntersectsRay(s.Camera.Position, rayDirection)
			if intersects {
				c := color.NRGBA{
					R: uint8((x + y) & 255),
					G: uint8((x + y) << 1 & 255),
					B: uint8((x + y) << 2 & 255),
					A: 255,
				}
				img.Set(x, y, c)

			}
		}
	}

	return img, err
}
