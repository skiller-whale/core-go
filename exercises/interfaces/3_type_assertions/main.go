package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Crewmate struct {
	colorCache string
}

func (c Crewmate) color() string {
	return c.colorCache
}

func (c Crewmate) reportBody() {
}

type Imposter struct {
	Crewmate
}

func (i Imposter) vent() {
}
func (i Imposter) sabotage() {
}

// It could be anybody
type Player interface {
	color() string
}

func main() {
	var players []Player = []Player{
		Imposter{Crewmate: Crewmate{colorCache: "red"}},
		Crewmate{colorCache: "green"},
		Crewmate{colorCache: "blue"},
		Crewmate{colorCache: "yellow"},
	}

	// They could be anywhere
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(4, func(a int, b int) { players[a], players[b] = players[b], players[a] })

	for _, player := range players {
		fmt.Print(player.color(), " player is ")
		if true {
			fmt.Println("the imposter!")
		} else {
			fmt.Println("a crewmate")
		}
	}
}
