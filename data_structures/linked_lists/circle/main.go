package main

import (
	"errors"
	"fmt"
)

var (
	ErrFull  = errors.New("buffer is full")
	ErrEmpty = errors.New("buffer is empty")
)

type CircleBuffer[T comparable] struct {
	data     []T
	head     int
	tail     int
	size     int
	capacity int
}

func NewCircleBuffer[T comparable](capacity int) *CircleBuffer[T] {
	return &CircleBuffer[T]{
		data:     make([]T, capacity),
		capacity: capacity,
	}
}

func (b *CircleBuffer[T]) Enque(item T) error {
	if b.size == b.capacity {
		return ErrFull
	}

	b.data[b.tail] = item
	b.tail = (b.tail + 1) % b.capacity 
	b.size++

	return nil
}

func (b *CircleBuffer[T]) Deque() (T, error) {
	var zero T

	if b.size == 0 {
		return zero, ErrEmpty
	}

	item := b.data[b.head]
	b.data[b.head] = zero

	b.head = (b.head + 1) % b.capacity
	b.size--

	return item, nil
}

func (b *CircleBuffer[T]) IsFull() bool {
	return b.size == b.capacity
}

func (b *CircleBuffer[T]) IsEmpty() bool {
	return b.size == 0
}

func (b *CircleBuffer[T]) Print() {
	if flag := b.IsEmpty(); flag {
		fmt.Println("buffer is empty")
		return
	}

	for i, v := range b.data {
		fmt.Printf("ind %d: val %v\n", i, v)
	}
}

func main() {
	newBuf := NewCircleBuffer[int](10)
	newBuf.Enque(1)
	newBuf.Enque(2)
	newBuf.Enque(3)
	newBuf.Enque(4)
	newBuf.Enque(5)

	newBuf.Print()
}