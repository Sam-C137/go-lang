package main

import (
	"fmt"
	"time"
)

func main() {
	go say("world")
	say("hello")
}

func say(msg string) {
	for i := range 5 {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(msg, i)
	}
}
