package main

import "fmt"

func functional() {
	exampleFunc := func(i int) {
		fmt.Println(i)
	}

	for v := range 10 {
		exampleFunc(v)
	}
}

func main() {
	functional()
}
