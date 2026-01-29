package main

import "fmt"

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
	Prev  *Node[T]
}

type DoublyLinkedList[T comparable] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func (l *DoublyLinkedList[T]) PushBack(val T) {
	newNode := &Node[T]{Value: val}

	if l.tail == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.Prev = l.tail
		l.tail.Next = newNode
		l.tail = newNode
	}

	l.size++
}

func (l *DoublyLinkedList[T]) PushFront(val T) {
	newNode := &Node[T]{Value: val}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.Next = l.head
		l.head.Prev = newNode
		l.head = newNode
	}

	l.size++
}

func (l *DoublyLinkedList[T]) Remove(node *Node[T]) {
	if node == nil {
		return
	}

	if node.Prev != nil {
		node.Prev.Next = node.Next
	} else {
		l.head = node.Next
	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	} else {
		l.tail = node.Prev
	}

	l.size--
}

func (l *DoublyLinkedList[T]) RemoveByValue(val T) bool {
	curr := l.head

	for curr != nil {
		if curr.Value == val {
			l.Remove(curr)
			return true
		}
	}

	return false
}

func (l *DoublyLinkedList[T]) DisplayForward() {
	curr := l.head

	for curr != nil {
		fmt.Printf("%v <-> ", curr.Value)
		curr = curr.Next
	}
	fmt.Println("nil")
}

func (l *DoublyLinkedList[T]) DisplayBackward() {
	curr := l.tail

	for curr != nil {
		fmt.Printf("%v <-> ", curr.Value)
		curr = curr.Prev
	}

	fmt.Println("nil")
}

func NewList[T comparable]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}

func main() {
	myList := NewList[int]()

	myList.PushFront(1)
	myList.PushFront(2)
	myList.PushFront(3)

	myList.DisplayForward()
}
