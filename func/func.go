package main

import (
	"fmt"
	"math"
)

func main() {
	fibTest()
}

func funcVal() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt((x * x) + (y * y))
	}

	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func closure() {
	pos, neg := adder(), adder()

	for i := range 10 {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func adder() func(int) int {
	sum := 0

	return func(x int) int {
		sum += x
		return sum
	}
}

func fibonacci() func() int {
	prev := 0
	curr := 1
	count := 0

	return func() int {
		if count == 0 || count == 1 {
			count++
			return count - 1
		}
		prev, curr = curr, prev+curr
		return curr
	}
}

func fibTest() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
