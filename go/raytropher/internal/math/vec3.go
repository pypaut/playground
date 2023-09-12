package math

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
