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

func NewSet[T comparable](s ...T) (o Set[T]) {
	o = make(Set[T])
	for _, v := range s {
		o.Add(v)
	}
	return o
}

// func SetProduct[A comparable, B comparable, ???](a Set[A], b Set[B], ???) Set[ ??? ] {
// 	p := make( Set[ ??? ] )
// 	for am, _ := range a {
// 		for bm, _ := range b {
// 			p.Add( ??? )
// 		}
// 	}
// 	return p
// }

type CardSuit rune
type CardRank string
type Card struct {
	rank CardRank
	suit CardSuit
}

func NewCard(n CardRank, s CardSuit) Card {
	return Card{rank: n, suit: s}
}

func (c Card) String() string {
	return fmt.Sprintf("%s%s", c.rank, string(c.suit))
}

type ChessBoardSquare struct {
	col rune
	row int
}

func NewSquare(col rune, row int) ChessBoardSquare {
	return ChessBoardSquare{col: col, row: row}
}

func (c ChessBoardSquare) String() string {
	return fmt.Sprintf("%s%d", string(c.col), c.row)
}

func main() {
	suits := NewSet([]CardSuit{'♥', '♣', '♦', '♠'}...)
	ranks := NewSet([]CardRank{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}...)
	//deck := SetProduct(ranks, suits) ???
	//fmt.Println("Deck of cards → ", deck)

	cols := NewSet([]rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H'}...)
	rows := NewSet([]int{1, 2, 3, 4, 5, 6, 7, 8}...)
	//chessBoard := SetProduct(cols, rows) ???
	//fmt.Println("Chess board   → ", chessBoard)
}
