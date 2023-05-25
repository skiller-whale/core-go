package main

import "fmt"

type Priced interface {
	Price() int
}

type Handbag struct {
	pence  int
	colour string
}

type SuedeProtectorSpray struct {
	Price int
}

type CowboyBoots struct {
	cents int
}

func (h Handbag) price() int {
	return h.pence
}
func (h CowboyBoots) price() float64 {
	return float64(h.cents) * 0.8 // convert from US Dollars
}

func main() {
	items := []Priced{
		Handbag{pence: 3854, colour: "red"},
		SuedeProtectorSpray{price: 495},
		CowboyBoots{cents: 9995},
	}

	sum := 0
	for _, item := range items {
		sum += item.Price()
		fmt.Println("Adding", item)
	}
	fmt.Println("Total price =", sum)
}
