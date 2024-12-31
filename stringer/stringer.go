package main

import (
	"fmt"
	"slices"
	"strings"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}

	fmt.Println(a, z)

	s := "APE STRONK 🔥"
	fmt.Println(iterateStringWrong(s))
	fmt.Println(iterateStringRight(s))
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func iterateStringWrong(s string) string {
	arr := make([]string, len(s))

	for i := range s {
		arr[i] = string(s[i])
	}

	slices.Reverse(arr)

	return strings.Join(arr, "")
}

func iterateStringRight(s string) string {
	arr := make([]string, len(s))

	for i, r := range s {
		arr[i] = string(r)
	}

	slices.Reverse(arr)

	return strings.Join(arr, "")
}
