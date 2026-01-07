package main

import (
	"errors"
	"fmt"
)

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type Stack[T any] struct {
	top  *Node[T]
	size int
}


func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}


func (s *Stack[T]) Push(value T) {
	newNode := &Node[T]{
		Value: value,
		Next:  s.top,
	}

	s.top = newNode
	s.size++
}


func (s *Stack[T]) Pop() (T, error) {
	var zeroVal T

	if s.IsEmpty() {
		return zeroVal, errors.New("stack is empty")
	}

	value := s.top.Value
	s.top = s.top.Next
	s.size--

	return value, nil
}


func (s *Stack[T]) Peek() (T, error) {
	var zeroVal T

	if s.IsEmpty() {
		return zeroVal, errors.New("stack is empty")
	}

	return s.top.Value, nil
}


func (s *Stack[T]) Size() int {
	return s.size
}


func (s *Stack[T]) Print() {
	current := s.top
	
	for current != nil {
		fmt.Println(current)
		current = current.Next
	}
}


func (s *Stack[T]) IsEmpty() bool {
	return s.top == nil
}

func main() {
	newStack := NewStack[int]()

	newStack.Push(10)
	newStack.Push(20)
	newStack.Push(30)

	newStack.Peek()

	newStack.Pop()

	newStack.Print()
}