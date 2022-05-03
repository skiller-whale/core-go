package main

import (
	"image"
	"image/jpeg"
	"os"
)

func (p photograph) everyPixel(f func(x int, y int, p pixel)) {
	for x := 0; x < 6000; x++ {
		for y := 0; y < 4000; y++ {
			f(x, y, p[x][y])
		}
	}
}

func LoadPhotographFromJPEG(n string) (out photograph) {
	f, err := os.Open(n)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	out.everyPixel(func(x, y int, p pixel) {
		r, g, b, _ := img.At(x, y).RGBA()
		out[x][y] = pixel{r: uint8(r >> 8), g: uint8(g >> 8), b: uint8(b >> 8)}
	})

	return
}

func SavePhotographToJPEG(p photograph, n string) {
	f, err := os.Create(n)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rgba := image.NewRGBA(image.Rectangle{Min: image.Point{X: 0, Y: 0}, Max: image.Point{X: 6000, Y: 4000}})
	p.everyPixel(func(x, y int, p pixel) {
		offset := (y-rgba.Rect.Min.Y)*rgba.Stride + (x-rgba.Rect.Min.X)*4
		rgba.Pix[offset] = p.r
		rgba.Pix[offset+1] = p.g
		rgba.Pix[offset+2] = p.b
		rgba.Pix[offset+3] = 255
	})

	o := jpeg.Options{
		Quality: 90,
	}
	err = jpeg.Encode(f, rgba, &o)
	if err != nil {
		panic(err)
	}
}
