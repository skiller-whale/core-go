package main

import (
	"math/rand"
	"time"
)

type CellProcessor func(Grid, int, int) bool

// Run the CellProcessor function across every cell, setting the results in-place
func (g Grid) Tick(p CellProcessor) Grid {
	g2 := NewGrid(len(g[0]), len(g))
	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g[0]); x++ {
			g2.SetLive(x, y, p(g, x, y))
		}
	}
	return g2
}

// Given a grid, and some co-ordinates on the grid, returns whether the cell
// should be alive or dead in the next generation based on how many of its 8
// neighbours are alive.
//
func ConwaysLife(g Grid, x, y int) bool {
	liveNeighbours := g.LiveNeighbours(x, y)
	live := g.Live(x, y)
	if live {
		// A live cell dies if it has less than 2 neighbours (starvation) or more than 3 (overcrowding)
		if liveNeighbours < 2 || liveNeighbours > 3 {
			return false
		}
	} else {
		// An empty cell comes alive if it has more than 3 neighbours (reproduction)
		if liveNeighbours > 3 {
			return true
		}
	}
	// Otherwise the cell stays in the same state
	return live
}

func main() {
	rand.Seed(time.Now().UnixNano())
	g := NewGrid(40, 20)
	g.Randomize(50) // fill 50% of the grid randomly
	for {
		g.Print()
		g = g.Tick(ConwaysLife)
		time.Sleep(time.Second / 2)
	}
}
