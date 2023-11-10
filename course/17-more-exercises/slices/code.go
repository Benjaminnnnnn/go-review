package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		x := make([]uint8, dx)
		for j := 0; j < dx; j++ {
			x[j] = uint8(i * j / 2)
		}
		pic[i] = x
	}
	return pic
}

func main() {
	pic.Show(Pic)
}
