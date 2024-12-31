package main

import (
	"fmt"
	"strings"
)

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var s []int = primes[1:4]
	fmt.Println(s)
	ranger()
}

func slice() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

func sliceLiterals() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

func loop(arr []int) {
	for i := range len(arr) {
		fmt.Print(i, " ")
	}
}

func capacity() {
	s := []int{2, 3, 5, 7, 11, 13}
	printCapSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printCapSlice(s)

	// Extend its length.
	s = s[:4]
	printCapSlice(s)

	// Drop its first two values.
	s = s[2:]
	printCapSlice(s)
}

func printCapSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func makeSlice() {
	a := make([]int, 5)
	printMakeSlice("a", a)

	b := make([]int, 0, 5)
	printMakeSlice("b", b)

	c := b[:2]
	printMakeSlice("c", c)

	d := c[2:5]
	printMakeSlice("d", d)
}

func printMakeSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func tictactoe() {
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := range len(board) {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func appender() {
	var s []int
	printAppendSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printAppendSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printAppendSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printAppendSlice(s)
}

func printAppendSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func ranger() {
	var pow = make([]int, 10)

	for i := range pow {
		pow[i] = 1 << uint(i)
	}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
