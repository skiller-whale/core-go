package main

import (
	"fmt"
	"strings"
)

type pricedInKrona interface {
	price() int
}

type coloured interface {
	colour() string
}

type titled interface {
	title() string
}

type BaseProduct struct {
	titleEnglish string
	looksLike    string
	priceInPence int
}

func (b BaseProduct) price() int {
	// exchange rate
	return (b.priceInPence * 171) / 100
}

func (b BaseProduct) colour() string {
	// Always list colour first
	return strings.Split(b.looksLike, ",")[0]
}

func (b BaseProduct) title() string {
	// eigum við að þýða þetta sjálfkrafa?
	return b.titleEnglish
}

func printIcelandicProducts(items []IcelandicProduct) {
	fmt.Printf("|Item name                     |Price (kr)|Colour         |\n")
	fmt.Printf("-----------------------------------------------------------\n")
	for _, item := range items {
		fmt.Printf("|%-30v|% 10v|%-15v|\n", item.title(), item.price(), item.colour())
	}
}

func main() {
	items := []BaseProduct{
		{titleEnglish: "Whale soft toy", looksLike: "Red", priceInPence: 995},
		{titleEnglish: "Poster of the Icelandic plains", looksLike: "White, entirely", priceInPence: 1495},
		{titleEnglish: "Puffin keychain", looksLike: "Orange", priceInPence: 300},
	}

	printIcelandicProducts(items)
}
