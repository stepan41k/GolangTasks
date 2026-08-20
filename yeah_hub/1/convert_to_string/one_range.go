package main

import (
	"fmt"
)

func OneRange() {
	x := 156

	var runes [11]rune
	pos := len(runes) - 1 

	for x > 0 {
		digit := x % 10
		runes[pos] = '0' + rune(digit)
		pos--
		x /= 10
	}
		
	fmt.Println(string(runes[:]))
}