package main

import (
	"fmt"
	"slices"
)

func differenceInAges(ages []int) []int {
	max := slices.Max(ages)
	min := slices.Min(ages)

	return []int{min, max, max - min}
}

func main() {
	fmt.Println(differenceInAges([]int{82, 15, 6, 38, 35})) // -> [6, 82, 76]
	fmt.Println(differenceInAges([]int{57, 99, 14, 32}))    // -> [14, 99, 85]
	fmt.Println(differenceInAges([]int{25}))                // -> [25, 25, 0]
	fmt.Println(differenceInAges([]int{10, 10, 10}))        // -> [10, 10, 0]
}
