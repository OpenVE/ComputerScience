package vectors

import (
	"math"
)

func Hypotenuse(n ... float64) float64 {
	var pows float64
	for i := 0; i < len(n); i++ {
		pows += math.Pow(n[i], 2)
	}
	return math.Sqrt(pows)	
}
