package main

import "fmt"

type Set[T comparable] map[T]bool

func (s Set[T]) Add(i T) {
	s[i] = true
}

func (s Set[T]) Remove(i T) {
	delete(s, i)
}

func (s Set[T]) Contains(i T) bool {
	return s[i]
}

func (s Set[T]) Members() (o []T) {
	for k, _ := range s {
		o = append(o, k)
	}
	return o
}

func (s Set[T]) String() string {
	return fmt.Sprintf("%v", s.Members())
}

func main() {
	whales := make(Set[string])
	whales.Add("Right")
	whales.Add("Minke")
	whales.Add("Blue")
	whales.Add("Gray")

	fmt.Println("whales contains 'Blue'?   → ", whales.Contains("Blue"))
	fmt.Println("whales contains 'Beluga'? → ", whales.Contains("Beluga"))

	//colours := NewSet("Red", "Blue", "Yellow", "Gray")
	//fmt.Println("colours ∪ whales          → ", whales.Union(colours))
}
