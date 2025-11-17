package main

import (
	"fmt"
)

func AllReplaces(arr []int) [][]int {
    var result [][]int

    var backtrack func(pos int)
    backtrack = func(pos int) {
        if pos == len(arr) {
            perm := append([]int{}, arr...)
            result = append(result, perm)
            return
        }

        for i := pos; i < len(arr); i++ {
            arr[pos], arr[i] = arr[i], arr[pos]

            backtrack(pos + 1)

            arr[pos], arr[i] = arr[i], arr[pos]
        }
    }

    backtrack(0)
    return result
}
 
func main() {
	fmt.Println(AllReplaces([]int{1, 2, 3}))
}