package main

import (
	V "./vectors"
)

func main() {
	v := V.Vector{ X: 4, Y: 3 }
	println(v.Length())
	v = V.Vector{ X: 1, Y: 4, Z: 3 }
	println(v.Length())
	
	println("Add Vectors")
	b := V.Vector{ X: 2, Y: 1 }
	c := V.Vector{ X: 1, Y: -3 }
	b.Add(c)
	println(b.String())
	b = V.Vector{ X: 2, Y: 1 }
	b.Mult(3.0/2)
	println(b.String())
	b = V.Vector{ X: 2, Y: 1 }
	b.Mult(-1.0/2)
	println(b.String())
}
