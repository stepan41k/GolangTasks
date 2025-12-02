package main

import "fmt"

type MyArray struct {
	arr  []int
	size int
}

func NewArray(size int) *MyArray {
	return &MyArray{
		arr:  make([]int, size),
		size: size,
	}
}

func (ma *MyArray) Set(index int, value int) {
	if index >= ma.size {
		fmt.Println("index out of range")
	} else {
		ma.arr[index] = value
	}
}

func (ma *MyArray) Get(index int) int {
	if index >= ma.size {
		fmt.Println("index out of range") 
	} else {
		return(ma.arr[index])
	}

	return -1
}

func main() {
	x := NewArray(10)

	x.arr[1] = 10

	for _, v := range x.arr {
		fmt.Println(v)
	}
}