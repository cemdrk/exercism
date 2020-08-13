// Package triangle provides utility functions for a triangle.
package triangle

import "math"

// Kind triangle type.
type Kind int

const (
	// NaT not a triangle
	NaT = iota
	// Equ equilateral
	Equ
	// Iso isosceles
	Iso
	// Sca scalene
	Sca
)

func isSideValid(a, b, c float64) bool {
	return math.Abs(b-c) <= a && a <= (b+c)
}

func isTriangle(a, b, c float64) bool {
	return isSideValid(a, b, c) && isSideValid(b, a, c) && isSideValid(c, b, a)
}

func isAnyNegativeOrZero(nums ...float64) bool {
	for _, num := range nums {
		if num <= 0 {
			return true
		}
	}
	return false
}

// KindFromSides determines type of triangle.
func KindFromSides(a, b, c float64) Kind {
	switch {
	case isAnyNegativeOrZero(a, b, c):
		fallthrough
	case !isTriangle(a, b, c):
		return NaT
	}

	if a == b && a == c {
		return Equ
	}

	if a == b || b == c || c == a {
		return Iso
	}

	return Sca
}
