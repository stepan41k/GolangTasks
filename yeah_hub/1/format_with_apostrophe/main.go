package main

import (
	"fmt"
	"strconv"
	"strings"
)

func FormatWithApostrophe(price int) string {
	s := strconv.Itoa(price)
	n := len(s)
	
	if n <= 3 {
		return s
	}
	
	res := strings.Builder{}
	
	for i, char := range s {
		if i > 0 && (n-i) % 3 == 0 {
			res.WriteByte('\'')
		}
		res.WriteRune(char)
	}
	
	return res.String()
}

func main() {
	fmt.Println(FormatWithApostrophe(12345678))
	// fmt.Println(12345678 % 1000, 12345678 / 1000)
}