package main

import (
	"fmt"
)

// Условие: Найдите число Фибоначчи на позиции n (нумерация с 0). Первые два числа Фибоначчи: 0 и 1. Пример: n = 7 → 13
// 0, 1, 1, 2, 3, 5, 8, 13, 21

func recursiveFibonachi(n int) int {
	if n <= 1 {
		return n
	}

	return recursiveFibonachi(n-1) + recursiveFibonachi(n-2)
}

func iterativeFibonachi(n int) int {
	if n < 0 {
		return -1
	}

	if n <= 1 {
		return n
	}

	first, second := 0, 1

	for i := 2; i <= n; i++ {
		next := first + second
		first = second
		second = next
	}

	return second
}


func main() {

	fmt.Println(recursiveFibonachi(4))
	fmt.Println(recursiveFibonachi(0))
	fmt.Println(recursiveFibonachi(7))
	fmt.Println(recursiveFibonachi(1))

	fmt.Println("")

	fmt.Println(iterativeFibonachi(4))
	fmt.Println(iterativeFibonachi(0))
	fmt.Println(iterativeFibonachi(7))
	fmt.Println(iterativeFibonachi(1))
}	