package main

type Parrot struct{}

func (p Parrot) Fly() string     { return "Parrot flying" }
func (p Parrot) Glide() string   { return "Parrot gliding" }
func (p Parrot) Breathe() string { return "Parrot breathing" }
func (p Parrot) Talk() string    { return "Parrot talking" }
func (p Parrot) Walk() string    { return "Parrot walking" }

type FruitBat struct{}

func (f FruitBat) Fly() string     { return "Fruit bat flying" }
func (f FruitBat) Breathe() string { return "Fruit bat breathing" }
func (f FruitBat) Glide() string   { return "Fruit bat gliding" }

type Dog struct{}

func (d Dog) Walk() string    { return "Dog walking" }
func (d Dog) Breathe() string { return "Dog breathing" }

type Human struct{}

func (h Human) Walk() string    { return "Human walking" }
func (h Human) Breathe() string { return "Human breathing" }
func (h Human) Talk() string    { return "Human talking" }

type SugarGlider struct{}

func (s SugarGlider) Glide() string   { return "Sugar glider gliding" }
func (s SugarGlider) Breathe() string { return "Sugar glider breathing" }
func (s SugarGlider) Walk() string    { return "Sugar glider walking" }

func PrintTypes(b ...any) {
	for _, v := range b {
		println(v)
	}
}

// Add the interfaces and methods here.

func main() {
	p, f, d, h, s := Parrot{}, FruitBat{}, Dog{}, Human{}, SugarGlider{}

	// Implement these methods, whose arguments types should be interfaces.
	PrintTypes(p, f, d, h, s)
	BreatheTogether(p, f, d, h, s)
	TalkTogether(p, h)
	WalkTogether(p, d, h, s)
	FlyThenGlideTogether(p, f)
}
