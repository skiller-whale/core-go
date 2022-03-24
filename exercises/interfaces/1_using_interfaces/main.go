package main

import "fmt"

type priced interface {
	price() int
}

type Handbag struct {
	pence  int
	colour string
}
type Footspray struct {
	price int
}
type ColaBottle struct{}

func (h Handbag) price() int {
	return h.pence
}
func (h ColaBottle) price() float64 {
	return 1.0 // Always costs 1 penny
}

func main() {
	items := []priced{
		Handbag{pence: 1995},
		Footspray{price: 495},
		ColaBottle{},
	}

	sum := 0
	for _, item := range items {
		sum += item.price()
		fmt.Println("Adding", item)
	}
	fmt.Println("Total price =", sum)
}
