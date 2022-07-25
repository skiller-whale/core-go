package main

import (
	"fmt"
	"math/rand"
)

type Grid [][]bool

// Modulus function, but only positive numbers (like C)
func mod2(i, n int) int {
	return ((i % n) + n) % n
}

// Create a new Grid of the given dimensions
func NewGrid(w, h int) (o Grid) {
	o = make([][]bool, h)
	for i := 0; i < len(o); i++ {
		o[i] = make([]bool, w)
	}
	return o
}

// Copy the contents of `i` into a new Grid
func CopyGrid(i Grid) (o Grid) {
	o = NewGrid(len(i[0]), len(i))
	for y := 0; y < len(o); y++ {
		for x := 0; x < len(o[0]); x++ {
			o[y][x] = i[y][x]
		}
	}
	return o
}

// Set a random percentage of the grid to be live cells
func (g Grid) Randomize(percent int) {
	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g[0]); x++ {
			g.SetLive(x, y, rand.Intn(100) <= percent)
		}
	}

}

// Set the live state of a given cell (references wrap around both ways)
func (g Grid) SetLive(x, y int, v bool) {
	g[mod2(y, len(g))][mod2(x, len(g[0]))] = v
}

// Read the live state of a given cell
func (g Grid) Live(x, y int) bool {
	return g[mod2(y, len(g))][mod2(x, len(g[0]))]
}

// Return the number of live neighbour cells of the given cell
func (g Grid) LiveNeighbours(x, y int) (n int) {
	if g.Live(x-1, y-1) {
		n += 1
	}
	if g.Live(x, y-1) {
		n += 1
	}
	if g.Live(x+1, y-1) {
		n += 1
	}
	if g.Live(x+1, y) {
		n += 1
	}
	if g.Live(x+1, y+1) {
		n += 1
	}
	if g.Live(x, y+1) {
		n += 1
	}
	if g.Live(x-1, y+1) {
		n += 1
	}
	if g.Live(x-1, y) {
		n += 1
	}
	return n
}

// Print the grid to STDOUT, returning the cursor to 0,0 first
func (g Grid) Print() {
	fmt.Print("\033[H")
	for _, row := range g {
		for _, cell := range row {
			if cell {
				fmt.Print("🤍")
			} else {
				fmt.Print("🖤")
			}
		}
		fmt.Println()
	}
}
