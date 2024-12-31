package main

import (
	"fmt"
	"strings"
)

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["foo"] = Vertex{4.002, -2.0675}
	fmt.Println(m["foo"])
	// WordCount("foo bar baz")
	fmt.Print(WordCount("foo, bar, baz"))
}

func mapLiteral() {
	var m = map[string]Vertex{
		"foo": {
			4.002, -2.0675,
		},
		"bar": {
			37.42202, -122.08408,
		},
	}
	fmt.Println(m)
}

func mutMap() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value: ", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value: ", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value: ", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value: ", v, "Present? ", ok)
}

func WordCount(s string) map[string]int {
	dict := make(map[string]int)

	for _, w := range strings.Fields(s) {
		_, ok := dict[w]
		if !ok {
			dict[w] = 1
			continue
		}
		dict[w]++
	}

	return dict
}
