package sphere

import m "raytropher/internal/math"

type Sphere struct {
	Position m.Vec3  `yaml:"position"`
	Radius   float64 `yaml:"radius"`
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
	rayPosition, rayDirection m.Vec3,
) (intersects bool, point m.Vec3) {
	return false, m.Vec3{0, 0, 0}
}
