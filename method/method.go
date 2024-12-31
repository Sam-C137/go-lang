package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt((v.X * v.X) + (v.Y * v.Y))
}

func (v *Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}

type SuperFloat float64

func (f SuperFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}
