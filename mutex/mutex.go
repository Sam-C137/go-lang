package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	mut sync.Mutex
	val map[string]int
}

func (c *SafeCounter) Inc(key string) {
	c.mut.Lock()
	defer c.mut.Unlock()
	c.val[key]++
}

func (c *SafeCounter) Value(key string) int {
	c.mut.Lock()
	defer c.mut.Unlock()
	return c.val[key]
}

func main() {
	wg_counter()
}

func counter() {
	c := SafeCounter{val: make(map[string]int)}

	for range 1000 {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}

func wg_counter() {
	c := SafeCounter{val: make(map[string]int)}
	var wg sync.WaitGroup

	wg.Add(1000)

	for range 1000 {
		go func() {
			c.Inc("somekey")
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(c.Value("somekey"))
}
