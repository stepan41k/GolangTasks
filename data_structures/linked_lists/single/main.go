package main

import "fmt"

type Node struct {
	Val      int
	nextNode *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Insert(val int) {
	newNode := &Node{Val: val, nextNode: nil}

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
func (l *LinkedList) Delete(val int) {
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

func (l *LinkedList) Print() {
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
	newList := &LinkedList{}

	newList.Insert(10)
	newList.Insert(20)
	newList.Insert(30)

	newList.Delete(10)

	newList.Print()
}
