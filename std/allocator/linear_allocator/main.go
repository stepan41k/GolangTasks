package main

import (
	"fmt"
	"unsafe"
)

func store[T any](pointer unsafe.Pointer, value T) {
	*(*T)(pointer) = value
}

func load[T any](pointer unsafe.Pointer) T {
	return *(*T)(pointer)
}

func main() {
	const MB = 1 << 20
	allocator, err := NewLinearAllocator(MB)
	if err != nil {
		// handling
	}

	defer allocator.Free()

	pointer1, _ := allocator.Allocate(2)
	pointer2, _ := allocator.Allocate(4)

	store[int16](pointer1, 100)
	store[int16](pointer2, 200)

	value1 := load[int16](pointer1)
	value2 := load[int16](pointer2)

	fmt.Println("value1:", value1)
	fmt.Println("value2:", value2)

	fmt.Println("address1:", pointer1)
	fmt.Println("address2:", pointer2)
}