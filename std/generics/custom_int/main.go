package main

import (
	"fmt"
	"strconv"
)

type CustomConstraint interface {
	~int
	String() string
}

type customInt int

func (i customInt) String() string {
	return strconv.Itoa(int(i))
}

func GenericsTest [Z CustomConstraint] (x1 Z, x2 Z) bool {
	t := x1 == x2
	fmt.Println(t)
	return t
}

func main() {
	var ci1 customInt = 1
	var ci2 customInt = 1
	GenericsTest(ci1, ci2)
}