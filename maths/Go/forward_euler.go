// A forward Euler example
// for a free fall for 50 steps

package main

import "fmt"

const steps = 50

func main() {
	h := 0.1  // s    // step of time
	g := 9.81 // m/s2 // Grativy Acceleration
	hg := h * g
	var t, x, v [steps + 1]float64
	for s := 0; s < steps; s++ {
		t[s + 1] = h * (float64(s) + 1) // Time
		x[s + 1] = x[s] + h * v[s]      // Position
		v[s + 1] = v[s] - hg            // Velocity
	}
	println(makeUrlPlot(t, x), "\n")
	println(makeUrlPlot(t, v))
}

func makeUrlPlot(x, y [steps + 1]float64) string {
	sx := fmt.Sprintf("%#v", x)
	sx = sx[12:len(sx)-1]
	sy := fmt.Sprintf("%#v", y)
	sy = sy[12:len(sy)-1]
	return "http://www.mathcracker.com/scatterplotimage.php?datax=" + sx + "&datay=" + sy
}
