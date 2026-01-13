package main

import (
	"fmt"
)

func maxLength(str1 string) int {
	maxLength := 1
	curLength := 1
	rs := []rune(str1)

	if len(rs) == 0 {
		return 0
	}

	for i := 0; i < len(rs)-1; i++ {
		if rs[i] == rs[i+1] {
			curLength++
		} else {
			if curLength > maxLength {
				maxLength = curLength
			}
			curLength = 1
		}
	}

	return maxLength
}

func main() {
	fmt.Println(maxLength("aaabbccccdd"))
	fmt.Println(maxLength("a"))
	fmt.Println(maxLength(""))
}