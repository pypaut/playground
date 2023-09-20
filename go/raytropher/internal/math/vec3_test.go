package math

import "testing"

func TestVec3Add(t *testing.T) {
	cases := []struct{ u, v, expected Vec3 }{
		{Vec3{0, 0, 0}, Vec3{1, 1, 1}, Vec3{1, 1, 1}},
		{Vec3{-10.5, 4.321, 0}, Vec3{123.3, -4.321, 1}, Vec3{112.8, 0, 1}},
		{Vec3{0, 0, 0}, Vec3{0, 0, 0}, Vec3{0, 0, 0}},
	}

	for _, c := range cases {

		c.u.Add(c.v)
		if !c.u.Equals(c.expected) {
			t.Logf("error: should be %+v, but got %+v", c.expected, c.u)
			t.Fail()
		}
	}
}

func TestVec3Times(t *testing.T) {
	cases := []struct {
		u        Vec3
		x        float64
		expected Vec3
	}{
		{Vec3{123, -39, 3}, 0, Vec3{0, 0, 0}},
		{Vec3{-1983.0, 3921, 0.00003}, 0, Vec3{0, 0, 0}},
		{Vec3{123, -39, 3}, 1, Vec3{123, -39, 3}},
		{Vec3{123, -39, 3}, -1, Vec3{-123, 39, -3}},
	}

	for _, c := range cases {
		result := c.u.Times(c.x)
		if !result.Equals(c.expected) {
			t.Logf("error: should be %+v, but got %+v", c.expected, c.u)
			t.Fail()
		}
	}
}

func TestVec3Cross(t *testing.T) {
	cases := []struct{ u, v, expected Vec3 }{
		{
			Vec3{},
			Vec3{},
			Vec3{},
		},
		{
			Vec3{-1, -2, 3},
			Vec3{4, 0, -8},
			Vec3{16, 4, 8},
		},
		{
			Vec3{1, 2, 3},
			Vec3{1, 5, 7},
			Vec3{-1, -4, 3},
		},
	}

	for _, c := range cases {
		result := c.u.Cross(c.v)
		if !result.Equals(c.expected) {
			t.Logf("error: should be %+v, but got %+v", c.expected, c.u)
			t.Fail()
		}
	}
}
