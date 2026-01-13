package main

import (
	"fmt"
)

func isAnnagramm(str1, str2 string) bool {
	m := map[rune]int{}

	if len(str1) != len(str2) {
		return false
	}

	for _, v := range str1 {
		if _, ok := m[v]; !ok {
			m[v]++
		}
	}

	for _, v := range str2 {
		if val, ok := m[v]; val <= 0 || !ok {
			return false
		} else {
			m[v]--
		}
	}

	return true
}

func main() {
	fmt.Println(isAnnagramm("listen", ""))
}