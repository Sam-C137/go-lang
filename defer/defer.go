package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	_, err := copy("in.txt", "out.txt")
	if err != nil {
		fmt.Println("Failed to copy in.txt to out.txt")
	}

	fmt.Println("Copy successful")
}

func copy(sourcePath, dstPath string) (res int64, err error) {
	src, err := os.Open(sourcePath)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstPath)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
