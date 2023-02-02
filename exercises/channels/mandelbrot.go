package main

import (
	"fmt"
	"math/cmplx"
)

func mandelbrot(a complex128, it int) (z complex128) {
	for i := 0; i < it; i++ {
		z = z*z + a
	}
	return
}

func shade(a complex128) rune {
	s := int(cmplx.Abs(a)/0.5) + 2
	if s > 4 {
		s = 4
	}
	if s < 0 {
		s = 0
	}
	return []rune{'.', '░', '▒', '▓', '█'}[s]
}

func main() {
	iterations := 50
	x0, x1, xr := -2.0, 0.5, 100
	y0, y1, yr := -1.0, 1.0, 40

	for y := y1; y >= y0; y -= ((y1 - y0) / float64(yr)) {
		line := make([]rune, 0, xr)
		for x := x0; x <= x1; x += ((x1 - x0) / float64(xr)) {
			line = append(line, shade(mandelbrot(complex(x, y), iterations)))
		}
		fmt.Println(string(line))
	}
}
