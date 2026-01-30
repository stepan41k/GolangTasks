package main

import (
	"errors"
	"fmt"
)

var (
	ErrEmptyQueue = errors.New("queue is empty")
)

type Queue[T comparable] struct {
	items []T
}

func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() (T, error) {
	var zero T

	if q.Size() == 0 {
		return zero, ErrEmptyQueue
	}

	item := q.items[0]

	q.items[0] = zero

	q.items = q.items[1:]
	return item, nil
}

func (q *Queue[T]) Size() int {
	return len(q.items)
}

func (q *Queue[T]) Print() {
	for i, v := range q.items {
		fmt.Printf("%d: %v\n", i, v)
	}
}

func NewQueue[T comparable]() *Queue[T] {
	return &Queue[T]{
		items: make([]T, 0),
	}
}


func main() {
	q := NewQueue[int]()

	q.Enqueue(5)
	q.Enqueue(10)
	q.Enqueue(15)
	q.Enqueue(20)

	q.Dequeue()

	q.Print()
}