package vectors

import (
	"strconv"
)

// type Vector []float64 ???
type Vector struct {
	X float64
	Y float64
	Z float64
}

func (v *Vector) Length() float64 {
	// return Hypotenuse(v...) ???
	return Hypotenuse(v.X, v.Y, v.Z)
}

func (v *Vector) String() string {
	s := "[" + strconv.FormatFloat(v.X, 'f', 2, 64)
	s += ", " + strconv.FormatFloat(v.Y, 'f', 2, 64)
	if v.Z != 0 {
		s += ", " + strconv.FormatFloat(v.Z, 'f', 4, 64)
	}
	return s + "]"
}

func (v *Vector) Add(vs ... Vector) {
	for _, vi := range vs {
		v.X += vi.X
		v.Y += vi.Y
		v.Z += vi.Z
	}
}

func (v *Vector) Mult(n float64) {
	v.X *= n
	v.Y *= n
	v.Z *= n
}
