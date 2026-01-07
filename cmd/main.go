package main

import (
	"fmt"
	"strconv"
)



func main() {
	var x int32 = 10

	y := strconv.FormatInt(int64(x), 10)

	fmt.Println(y)
}