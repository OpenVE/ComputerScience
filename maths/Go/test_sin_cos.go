// go run test_sin_cos.go
// Sin after Asin:  +7.071068e-001 +7.071068e-001 true
// Cos after Asin:  +7.071068e-001 +7.071068e-001 false +1.110223e-016
// Sin after Acos:  +7.071068e-001 +7.071068e-001 false +1.110223e-016
// Cos after Acos:  +7.071068e-001 +7.071068e-001 true

package main

import M "math"

func main() {
	// Triangle Rectangle
	Tr := [3][2]float64{{0,0}, {50,0}, {50, 30}}

	// Sides
	a := Tr[2][0] - Tr[0][0]
	b := Tr[2][0] - Tr[0][0]
	c := M.Sqrt(M.Pow(a, 2) + M.Pow(b, 2))
	
	// Reason of b over c, or
	// How many times c is b?
	bc := b/c
	ac := a/c
	
	// Angle in vertex Tr[0]
	ß := M.Asin(bc)
	
	println("Sin after Asin: ", M.Sin(ß), bc, M.Sin(ß) == bc)
	println("Cos after Asin: ", M.Cos(ß), ac, M.Cos(ß) == bc, M.Cos(ß) - bc)

	// Angle in vertex Tr[0]
	ß = M.Acos(ac)
	println("Sin after Acos: ", M.Sin(ß), bc, M.Sin(ß) == bc, M.Sin(ß) - bc)
	println("Cos after Acos: ", M.Cos(ß), ac, M.Cos(ß) == bc)
}
