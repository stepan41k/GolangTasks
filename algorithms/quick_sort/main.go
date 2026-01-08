package main

import (
	"fmt"
)

func quickSort(arr []int, low int, high int) {
	if low < high {
		p := partition(arr, low, high)
		quickSort(arr, low, p-1)
		quickSort(arr, p+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j]= arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return i
}

func main(){

	arr := []int{10, 1, 12, 3, 4, 5, 8, 1}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}