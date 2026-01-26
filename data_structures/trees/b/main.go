package main

import "fmt"

type Node struct {
	Value int
	Left *Node
	Right *Node
}

type Tree struct {
	Root *Node
}

func (n *Node) PrintTree() {
	if n == nil {
		return
	}

	n.Left.PrintTree()
	fmt.Printf("%d -> ", n.Value)
	n.Right.PrintTree()
}

func (t *Tree) Insert(val int) {
	if t.Root == nil {
		t.Root = &Node{Value: val}
	} else {
		t.Root.IterativeInsert(val)
	}
}

func (n *Node) RecursiveInsert(val int) {
	newNode := &Node{
		Value: val,
	}

	if n == nil {
		n = newNode
		return
	}

	if val < n.Value {
		if n.Left == nil {
			n.Left = newNode
		} else {
			n.Left.RecursiveInsert(val)
		}
	} else if val > n.Value {
		if n.Right == nil {
			n.Right = newNode
		} else {
			n.Right.RecursiveInsert(val)
		}
	}
}

func (n *Node) IterativeInsert(val int) {
	curr := n

	for {
		if val < curr.Value {
			if curr.Left == nil {
				curr.Left = &Node{Value: val}
				break
			}
			curr = curr.Left
		} else if val > curr.Value {
			if curr.Right == nil {
				curr.Right = &Node{Value: val}
				break
			}
			curr = curr.Right
		} else {
			break
		}
	}
}

func (n *Node) findMin() *Node {
	current := n

	for current.Left != nil {
		current = current.Left
	}

	return current
}

func Delete(root *Node, val int) *Node{
	if root == nil {
		return nil
	}

	if val < root.Value {
		root.Left = Delete(root.Left, val)
	} else if val > root.Value {
		root.Right = Delete(root.Right, val)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		minRight := root.Right.findMin()
		
		root.Value = minRight.Value

		root.Right = Delete(root.Right, minRight.Value)
	}

	return root
}


func NewTree() *Tree {
	return &Tree{}
}

func main() {
	newTree := NewTree()

	newTree.Insert(20)
	newTree.Insert(10)
	newTree.Insert(25)
	
	newTree.Insert(5)
	newTree.Insert(0)

	newTree.Root.PrintTree()
	Delete(newTree.Root, 10)
	fmt.Println("")

	newTree.Root.PrintTree()

}