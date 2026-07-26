package main

import (
	"fmt"
	"strings"
)

func CapitalizeWords(s string) string {
	if len(s) == 0 {
		return ""
	}
	
	words := strings.Split(s, " ")
	builder := strings.Builder{}

	for i, w := range words {
		if 97 <= w[0] && w[0] <= 122 {
			builder.WriteByte(w[0] - 32)
			builder.WriteString(w[1:])
		} else {
			builder.WriteString(w)
		}

		if i != len(words)-1 {
			builder.WriteRune(' ')
		}
	}
	return builder.String()
}

func main() {
	fmt.Println(CapitalizeWords("already Capitalize"))
}