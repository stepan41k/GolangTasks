package main

import (
	"fmt"
	"strconv"
	"strings"
)

func fizzBuzz() string {
	builder := strings.Builder{}

	for i := 1; i <= 100; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			builder.WriteString("FizzBuzz")
		} else if i % 3 == 0 {
			builder.WriteString("Fizz")
		} else if i % 5 == 0 {
			builder.WriteString("Buzz")
		} else {
			builder.WriteString(strconv.Itoa(i))
		}

		if i != 100 {
			builder.WriteRune(' ')
		}
	}

	return builder.String()
}
func main() {
	fmt.Println(fizzBuzz())
}
