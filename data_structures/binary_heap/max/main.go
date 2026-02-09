package main

import "fmt"

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.heapifUp(len(h.array) - 1)
}

func (h *MaxHeap) ExtractMap() (int, error) {
	if len(h.array) == 0 {
		return 0, fmt.Errorf("heap is emprty")
	}

	extracted := h.array[0]
	lastIndex := len(h.array) - 1

	h.array[0] = h.array[lastIndex]
	h.array = h.array[:lastIndex]

	h.heapifDown(0)

	return extracted, nil
}

func (h *MaxHeap) heapifUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MaxHeap) heapifDown(index int) {
	listIndex := len(h.array) - 1
	l, r := 
}