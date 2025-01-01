package main

import "fmt"

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
	fibChannel()
}

func sum(s []int, c chan int) {
	out := 0

	for _, v := range s {
		out += v
	}

	c <- out
}

func bufferedChannel() {
	ch := make(chan int, 3) // Buffer size 3
	ch <- 1
	ch <- 2
	ch <- 3
	//ch <- 4 // BLOCKS if uncommented (buffer full)

	fmt.Println(<-ch) // 1 (buffer now has 2, 3)
	ch <- 4           // Now space, so this works

}

func unbufferedChaos() {
	c := make(chan int)
	go func() { c <- 7 }()
	go func() { c <- 4 }()
	go func() { c <- 3 }()
	go func() { c <- 9 }()
	go func() { c <- 0 }()

	// fmt.Println(<-c)  // don't do a single read or buffer the channel and loop to do all reads
	for range 5 {
		fmt.Println(<-c)
	}
}

func fibChannel() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
	/*for {
		v, ok := <-c

		if !ok {
			break
		}

		fmt.Println(v)
	}*/
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1

	for range n {
		c <- x
		x, y = y, x+y
	}

	close(c)
}

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func isTreeSame() {

}

func Walk(t *Tree, ch chan int) {
	var walker func(t *Tree)

	walker = func(t *Tree) {
		if t == nil {
			return
		}

		Walk(t.Left, ch)
		ch <- t.Value
		Walk(t.Right, ch)
	}

	walker(t)
	close(ch)
}

func Same(t1, t2 *Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	x, y := <-ch1, <-ch2

	if x != y {
		return false
	}

	return true
}
