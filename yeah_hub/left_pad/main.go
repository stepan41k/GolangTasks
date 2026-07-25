package main

import (
	"fmt"
	"strings"
)

func leftPad(symbolCount int, str string) string {
	if symbolCount <= len(str) {
		return str
	}

	dif := symbolCount - len(str)
	res := strings.Builder{}

	for i := 0; i < dif; i++ {
		res.WriteRune(' ')
	}

	res.WriteString(str)
	
    return res.String()
}

func main() {
	fmt.Println(leftPad(10, "hello"))
}