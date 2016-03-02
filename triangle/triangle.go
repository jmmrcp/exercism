package triangle

import (
	"fmt"
	"math"
	"sort"
)

type Kind string

const (
	Equ = "Equ"
	Iso = "Iso"
	Sca = "Sca"
	NaT = "NaT"
)

const testVersion = 2

// Code this function.
func KindFromSides(a, b, c float64) Kind {
	t := []float64{a, b, c}
	sort.Float64s(t)
	fmt.Println(t)
	switch {
	case (math.IsNaN(t[0]) || math.IsInf(t[2], 0) || t[0] <= 0 || t[0]+t[1] < t[2]):
		return NaT
	case t[0] == t[2]:
		return Equ
	case t[0] == t[1] || t[1] == t[2]:
		return Iso
	default:
		return Sca
	}
}

// Notice it returns this type.  Pick something suitable.

// Pick values for the following identifiers used by the test program.
// NaT // not a triangle
// Equ // equilateral
// Iso // isosceles
// Sca // scalene

// Organize your code for readability.
