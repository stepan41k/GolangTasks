package main

import "fmt"

func closureFunc() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}

func main() {
	next := closureFunc()

	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())

}
