package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}

	i, err := strconv.Atoi("42")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Successfully converted %d\n", i)
}

type ErrorImpl struct {
	When time.Time
	What string
}

func (e *ErrorImpl) Error() string {
	return fmt.Sprintf("At %v, %s", e.When, e.What)
}

func run() error {
	return &ErrorImpl{
		time.Now(),
		"It did not work",
	}
}
