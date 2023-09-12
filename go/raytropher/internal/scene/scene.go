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

func (u *Vec3) Add(v Vec3) {
	u.X += v.X
	u.Y += v.Y
	u.Z += v.Z
}

func (v *Vec3) Times(x float64) (vResult Vec3) {
	vResult.X = v.X * x
	vResult.Y = v.Y * x
	vResult.Z = v.Z * x
	return
}

func (v *Vec3) Minus(u Vec3) (vResult Vec3) {
	vResult.X = v.X - u.X
	vResult.Y = v.Y - u.Y
	vResult.Z = v.Z - u.Z
	return
}

func (u *Vec3) Cross(v Vec3) (p Vec3) {
	p.X = u.Y*v.Z - u.Z*v.Y
	p.Y = u.Z*v.X - u.X*v.Z
	p.Z = u.X*v.Y - u.Y*v.X
	return
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

	out.Close()
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
			var screenPixelPosition Vec3
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

func (s *Scene) CheckRayCollisionWithSphere(x, y int) (c color.Color, err error) {
	// Compute color to display (formulae with light, texture, blabla)
	c = color.NRGBA{
		R: uint8((x + y) & 255),
		G: uint8((x + y) << 1 & 255),
		B: uint8((x + y) << 2 & 255),
		A: 255,
	}
	return c, nil
}

/*
** Proof **
Line :
x(t) = vect_x * t + pt_x
y(t) = vect_y * t + pt_y
z(t) = vect_z * t + pt_z

Sphere :
(x - c_x)**2 + (y - c_y)**2 + (z - c_z)**2 = r**2

We get a 2nd degree polynom with substitution method,
and the answer we seek depends on the discriminant's sign.

Return the intersection point, if any.
If there are two, we return the closest to pos.

a, b and c are the same as in ax**2 + b*x + c = 0.
*/
func (s *Sphere) IntersectsRay(
	rayPosition, rayDirection Vec3,
) (intersects bool, point Vec3) {
	return false, Vec3{0, 0, 0}
}
