// string
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	x := "Hello dear friemd"

	ptr := unsafe.Pointer(unsafe.StringData(x))

	for i := 0; i < len(x); i++ {
		ptrVal := *(*byte)(ptr)
		fmt.Println(string(ptrVal))
		ptr = unsafe.Add(ptr, 1)
	}
}