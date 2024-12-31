package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Vertex struct {
	X int
	Y int
}

func main() {
	odogwu := Person{name: "KING APE"}
	o := &odogwu
	fmt.Print(o.name)
}
