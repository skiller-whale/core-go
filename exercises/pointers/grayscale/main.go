package main

import (
	"fmt"
	"time"
)

type pixel struct{ r, g, b uint8 }
type photograph [6 * 1000][4 * 1000]pixel

func GrayPhotograph(in photograph) (out photograph) {
	for x := range in {
		for y, p := range in[x] {
			// this looks much nicer than an average!
			gray := uint8(0.2126*float32(p.r) + 0.7152*float32(p.g) + 0.0722*float32(p.b))
			out[x][y] = pixel{r: gray, g: gray, b: gray}
		}
	}
	return
}

func main() {
	gray := LoadPhotographFromJPEG("./wavy-shapes.jpg")
	start := time.Now()
	GrayPhotograph(gray)
	fmt.Println("Conversion took %ss, saving file...", time.Since(start))
	SavePhotographToJPEG(gray, "./wavy-shapes-gray.jpg")
}
