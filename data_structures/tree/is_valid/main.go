package main

import (
	"fmt"
)

type Node struct { 
	Val int
	Left *Node
	Right *Node
	flag bool
}

func IsValid(node *Node) {
	if node == nil {
		return 
	}
	
	if node.Left == nil || node.Right == nil {
		return
	}

	if node.Left.Val > node.Val {
		node.flag = false
		return
	}

	if node.Right.Val < node.Val {
		node.flag = false
		return
	}

	IsValid(node.Left)
	IsValid(node.Right)
}
 
func main() {
	newTree1 := &Node{2, &Node{Val: 1}, &Node{Val: 3}, true}
	newTree2 := &Node{2, &Node{Val:3}, &Node{Val: 1}, true}
	IsValid(newTree1)
	fmt.Println(newTree1.flag)
	IsValid(newTree2)
	fmt.Println(newTree2.flag)
	IsValid(newTree1)
	fmt.Println(newTree1.flag)
}