package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Sqrt(6))
}

func Sqrt(x float64) float64 {
	z := x / 2

	for v := z; 1e-7 < math.Abs(z-v); {
		v = z
		z -= (z*z - x) / (2 * z)
	}

	return z
}
