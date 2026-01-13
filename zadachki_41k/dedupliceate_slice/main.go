package main

import (
	"fmt"
)

func deduplicateWithMap(arr []int) []int {
	m := map[int]struct{}{}
	newArr := []int{}

	for _, v := range arr {
		if _, f := m[v]; !f {
			m[v] = struct{}{}
			newArr = append(newArr, v)
		}
	}

	return newArr
}



func main() {
	fmt.Println(deduplicateWithMap([]int{1, 2, 2, 3, 4, 3, 5}))
	
	// fmt.Println(deduplicateInPlace([]int{1, 2, 2, 3, 4, 3, 5}))
}