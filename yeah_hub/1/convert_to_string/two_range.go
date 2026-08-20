package main

import (
	"fmt"
)

func TwoRanges() {
	x := 156

	var runes []rune

	for x > 0 {
		digit := x % 10
		runes = append(runes, '0' + rune(digit))
		x /= 10
	}

	for i, j := 0, len(runes) - 1; i < j; i, j = i + 1, j - 1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
		
	fmt.Println(string(runes))
}