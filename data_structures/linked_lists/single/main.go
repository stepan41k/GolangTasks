package main

import "fmt"

type Node[T comparable] struct {
	Val      T
	nextNode *Node[T]
}

type LinkedList[T comparable] struct {
	head *Node[T]
}

func (l *LinkedList[T]) Insert(val T) {
	newNode := &Node[T]{Val: val, nextNode: nil}

	if l.head == nil {
		l.head = newNode
		return
	}
	
	
	currNode := l.head

	for currNode.nextNode != nil {
		currNode = currNode.nextNode
	}

	currNode.nextNode = newNode
}

//Delete all occurrences of a number
func (l *LinkedList[T]) Delete(val T) {
	for l.head != nil && l.head.Val == val {
        l.head = l.head.nextNode
    }

	if l.head == nil {
		return
	}

	currNode := l.head

	for currNode.nextNode != nil {
		if currNode.nextNode.Val == val {
			currNode.nextNode = currNode.nextNode.nextNode
		} else {
			currNode = currNode.nextNode
		}	
	}
}

func (l *LinkedList[T]) Print() {
	if l.head == nil {
		return
	}

	currNode := l.head

	for currNode != nil {
		fmt.Println(currNode.Val)
		currNode = currNode.nextNode
	}
}

func main() {
	newList := &LinkedList[int]{}

	newList.Insert(10)
	newList.Insert(20)
	newList.Insert(30)

	newList.Delete(10)

	newList.Print()
}
