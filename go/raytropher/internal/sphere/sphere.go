package sphere

import (
	"math"
	m "raytropher/internal/math"
)

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

	/*
	 a = rayDir.dot(rayDir)
	        b = 2 * rayDir.dot(vec)
	        c = vec.dot(vec) - self.ray ** 2

	        delta = b ** 2 - 4 * a * c

	        if delta < 0:
	            return None

	        if delta == 0:  # Single solution
	            t = -b / (2 * a)
	            return rayDir.times(t) + rayPos

	        # Two solutions
	        t1 = (-b - m.sqrt(delta)) / (2 * a)
	        t2 = (-b + m.sqrt(delta)) / (2 * a)

	        pt1 = rayDir.times(t1) + rayPos
	        pt2 = rayDir.times(t2) + rayPos

	        if pt1.dist(rayPos) < pt2.dist(rayPos):
	            return pt1
	        return pt2
	*/

	vec := rayPosition.Minus(s.Position)
	a := rayDirection.Dot(rayDirection)
	b := rayDirection.Dot(vec)
	c := vec.Dot(vec) - math.Pow(s.Radius, 2)

	delta := math.Pow(b, 2) - 4*a*c

	if delta < 0 {
		return false, m.Vec3{}
	} else if delta == 0 {
		t := -b / (2 * a)
		result := rayDirection.Times(t)
		result.Add(rayPosition)
		return true, result
	} else if delta > 0 {
		t1 := (-b - math.Sqrt(delta)) / (2 * a)
		t2 := (-b + math.Sqrt(delta)) / (2 * a)

		pt1 := rayDirection.Times(t1)
		pt1.Add(rayPosition)

		pt2 := rayDirection.Times(t2)
		pt2.Add(rayPosition)

		if pt1.Distance(rayPosition) < pt2.Distance(rayPosition) {
			return true, pt1
		} else {
			return true, pt2
		}
	}
	return false, m.Vec3{}
}
