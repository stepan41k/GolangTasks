package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var x = [16]uint8{0, 1, 2, 4, 8, 16, 32, 64, 128, 255}
	ptr := unsafe.Pointer(&x)

	for i := 0; i < len(x); i++ {
		fmt.Println(*(*uint8)(ptr))
		ptr = unsafe.Add(ptr, unsafe.Sizeof(x[0]))
	}

	fmt.Println(unsafe.Sizeof(x))
}