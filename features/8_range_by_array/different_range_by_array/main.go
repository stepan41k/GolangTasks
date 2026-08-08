package main

import "fmt"

func main() {
	data := [...]int{1, 2, 3}
	for value := range data {
		fmt.Println(value)
	}

	for value := range &data {
		fmt.Println(value)
	}

	for value := range data[:] {
		fmt.Println(value)
	}
}
