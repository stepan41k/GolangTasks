package main

import (
	"fmt"
	"unsafe"
)

func SliceArythmetic[T any](s []T) {
	p := unsafe.Pointer(&s[0])
	size := unsafe.Sizeof(s[0])
	
	for i := 0; i < len(s); i++ {
		fmt.Println(*(*int)(p))
		p = unsafe.Add(p, size)
	}

	fmt.Println()
}
