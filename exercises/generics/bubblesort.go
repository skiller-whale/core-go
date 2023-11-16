package main

import "fmt"

// Sort a slice in place
//   compare should return:
//   * <0 if a < b
//   *  0 if a == b
//   * >0 if a > b
//
func sortInPlace(l []int, compare func(int, int) int) {
	for {
		swapped := false
		for i := range l[0 : len(l)-1] {
			if compare(l[i], l[i+1]) > 0 {
				l[i], l[i+1] = l[i+1], l[i]
				swapped = true
			}
		}
		if !swapped {
			return
		}
	}
}

func compare(a int, b int) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

func main() {
	a1 := []int{42, 4, 23, 15, 16, 8}
	//a2 := []string{"The", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}
	//a3 := []float64{10.3, 4.5, -3.4, 23.2, 2.33, 0.5, 6.66}

	sortInPlace(a1, compare)
	fmt.Println(a1)
	//bubbleSortInPlace(a2, compare)
	//fmt.Println(a2)
	//bubbleSortInPlace(a3, compare)
	//fmt.Println(a3)
}
