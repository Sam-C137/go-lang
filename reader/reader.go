package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("KING APE STRONK 🦍")
	b := make([]byte, 4)

	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (n int, err error) {
	n, err = r.r.Read(b)

	for i, str := range b {
		if 'A' <= str && str <= 'Z' {
			b[i] = 'A' + (str-'A'+13)%26
		} else if 'a' <= str && str <= 'z' {
			b[i] = 'a' + (str-'a'+13)%26
		}
	}

	return
}

func rotRead() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
