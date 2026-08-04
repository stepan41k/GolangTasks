package main

import (
	"fmt"
	"unsafe"
)

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 10}

	p := unsafe.Pointer(&x[0])

	for i := 0; i < len(x); i++ {
		fmt.Println(*(*int)(p))
		p = unsafe.Add(p, unsafe.Sizeof(x[0]))
	}
}