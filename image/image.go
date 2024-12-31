package main

import (
	"fmt"
	"image"
	"image/color"
)

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}

type Image struct{}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 256, 256)
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{
		R: uint8(x),
		G: uint8(y),
		B: uint8(x ^ y),
		A: 255,
	}
}
