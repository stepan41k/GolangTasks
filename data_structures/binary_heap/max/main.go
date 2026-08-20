package main

import (
	"cmp"
	"fmt"
)

type MaxHeap[T cmp.Ordered] struct {
	data []T
}

func NewMaxHeap[T cmp.Ordered]() *MaxHeap[T] {
	return &MaxHeap[T]{data: []T{}}
}

func (h *MaxHeap[T]) Push(val T) {
	h.data = append(h.data, val)
	h.siftUp(len(h.data) - 1)
}

func (h *MaxHeap[T]) Pop() (T, bool) {
	var zero T
	if len(h.data) == 0 {
		return zero, false
	}

	maxVal := h.data[0]
	lastIdx := len(h.data) - 1

	h.data[0] = h.data[lastIdx]
	h.data = h.data[:lastIdx]

	if len(h.data) > 0 {
		h.siftDown(0)
	}

	return maxVal, true
}

func (h *MaxHeap[T]) Peek() (T, bool) {
	var zero T
	if len(h.data) == 0 {
		return zero, false
	}
	return h.data[0], true
}

func (h *MaxHeap[T]) siftUp(idx int) {
	for idx > 0 {
		parentIdx := (idx - 1) / 2
		if h.data[idx] > h.data[parentIdx] {
			h.data[idx], h.data[parentIdx] = h.data[parentIdx], h.data[idx]
			idx = parentIdx
		} else {
			break
		}
	}
}

func (h *MaxHeap[T]) siftDown(idx int) {
	n := len(h.data)

	for {
		left := 2*idx + 1
		right := 2*idx + 2
		largest := idx

		if left < n && h.data[left] > h.data[largest] {
			largest = left
		}
		if right < n && h.data[right] > h.data[largest] {
			largest = right
		}

		if largest != idx {
			h.data[idx], h.data[largest] = h.data[largest], h.data[idx]
			idx = largest
		} else {
			break
		}
	}
}

func main() {
	heap := NewMaxHeap[int]()

	numbers := []int{15, 30, 8, 10, 50, 20}
	for _, num := range numbers {
		heap.Push(num)
	}

	fmt.Println("Max elem (Peek):", func() int { v, _ := heap.Peek(); return v }())

	fmt.Println("\nExtract by order (desc order):")
	for {
		val, ok := heap.Pop()
		if !ok {
			break
		}
		fmt.Printf("%d ", val)
	}
	// Output: 50 30 20 15 10 8
}
