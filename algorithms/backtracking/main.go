package main

import (
	"fmt"
)

// func AllReplaces(arr []int) [][]int {
//     var result [][]int

//     var backtrack func(pos int)
//     backtrack = func(pos int) {
//         if pos == len(arr) {
//             perm := append([]int{}, arr...)
//             result = append(result, perm)
//             return
//         }

//         for i := pos; i < len(arr); i++ {
//             arr[pos], arr[i] = arr[i], arr[pos]

//             backtrack(pos + 1)

//             arr[pos], arr[i] = arr[i], arr[pos]
//         }
//     }

//     backtrack(0)
//     return result
// }

func combine(n int, k int) [][]int {
	result := [][]int{}

	var backtrack func(start int, current []int)
	backtrack = func(start int, current []int) {
		if len(current) == k {
			comb := make([]int, k)
			copy(comb, current)
			result = append(result, comb)
			return
		}

		for i := start; i <= n; i++ {
			current = append(current, i)
			backtrack(i+1, current)
			current = current[:len(current)-1]
		}
	}

	backtrack(1, []int{})
	return result
}

func main() {
	fmt.Println(combine(4, 4))
}
