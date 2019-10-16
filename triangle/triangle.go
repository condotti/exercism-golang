package triangle

import "math"

type Kind int

const (
	NaT = iota	// not a triangle
	Equ		// equilateral
	Iso		// isosceles
	Sca		// scalene
)

func KindFromSides(a, b, c float64) Kind {
	switch {
	case math.IsNaN(a + b + c) || math.IsInf(a + b + c, 0) || a * b * c == 0 ||
		a > b + c || b > c + a || c > a + b:
		return NaT
	case a == b && b == c:
		return Equ
	case a == b || b == c || c == a:
		return Iso
	default:
		return Sca
	}
}
