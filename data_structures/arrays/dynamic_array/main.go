package main

import "fmt"

type DynamicArray struct {
	arr  []int
	size int
	cap  int
}

func NewDynamicArray(size int) *DynamicArray {
	return &DynamicArray{
		arr:  make([]int, 0, size),
		size: 0,
		cap: size,
	}
}

func (da *DynamicArray) Expand(cap int) *DynamicArray {
	newCap := da.cap * 2

	if da.cap >= 1024 {
		newCap = da.cap * 5 / 4
	}

	newArr := &DynamicArray{
		arr:  make([]int, da.size, newCap),
		size: da.size,
		cap:  newCap,
	}

	copy(newArr.arr, da.arr)
	return newArr
}

func (da *DynamicArray) Append(elems ...int) *DynamicArray {
	if da.size + len(elems) > da.cap {
		newArr := da.Expand(da.cap)
		da = newArr
	}

	da.arr = da.arr[:da.size+len(elems)]
	copy(da.arr[da.size:], elems)
	da.size += len(elems)

	return da
}

func main() {
	x := NewDynamicArray(2)

	fmt.Println(x.arr, x.size, x.cap)
	fmt.Printf("before cycle\n\n")

	for i := 0; i < 10; i++ {
		x = x.Append(i)
		fmt.Println(x.arr, x.size, x.cap)
	}
}