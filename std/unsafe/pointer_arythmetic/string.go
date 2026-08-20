package main

import (
	"fmt"
	"unsafe"
)

func StringArythmetic(s string) {
	ptr := unsafe.Pointer(unsafe.StringData(s))

	for i := 0; i < len(s); i++ {
		ptrVal := *(*byte)(ptr)
		fmt.Println(string(ptrVal))
		ptr = unsafe.Add(ptr, 1)
	}
}