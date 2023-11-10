package main

import (
	"image"
	"image/color"
	"math/rand"

	"golang.org/x/tour/pic"
)

type Image struct {
	w      int
	h      int
	pixels [][]color.RGBA
}

func (im *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (im *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, im.w, im.h)
}

func (im *Image) At(x, y int) color.Color {
	return im.pixels[x][y]
}

func randUint8() uint8 {
	return uint8(rand.Intn(256))
}

func generateImage(w, h int) Image {
	pixels := make([][]color.RGBA, h)

	for i := 0; i < h; i++ {
		row := make([]color.RGBA, w)
		for j := 0; j < w; j++ {
			row[j] = color.RGBA{randUint8(), randUint8(), randUint8(), randUint8()}
		}
		pixels[i] = row
	}

	return Image{
		w:      w,
		h:      h,
		pixels: pixels,
	}

}

func main() {
	m := generateImage(100, 100)
	pic.ShowImage(&m)
}
