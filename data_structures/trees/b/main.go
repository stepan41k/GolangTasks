package main

import (
	"cmp"
	"fmt"
)

type Node[T cmp.Ordered] struct {
	leaf     bool
	keys     []T
	children []*Node[T]
}

type BTree[T cmp.Ordered] struct {
	root *Node[T]
	t    int
}

func NewBTree[T cmp.Ordered](t int) *BTree[T] {
	return &BTree[T]{
		t: t,
	}
}

func (t *BTree[T]) Search(key T) (*Node[T], int) {
	if t.root == nil {
		return nil, -1
	}
	return t.root.search(key)
}

func (n *Node[T]) search(key T) (*Node[T], int) {
	i := 0
	for i < len(n.keys) && key > n.keys[i] {
		i++
	}

	if i < len(n.keys) && key == n.keys[i] {
		return n, i
	}

	if n.leaf {
		return nil, -1
	}

	return n.children[i].search(key)
}

func (t *BTree[T]) Insert(key T) {
	if t.root == nil {
		t.root = &Node[T]{
			leaf: true,
			keys: []T{key},
		}
		return
	}

	if len(t.root.keys) == 2*t.t-1 {
		newRoot := &Node[T]{
			leaf: false,
		}

		newRoot.children = append(newRoot.children, t.root)

		t.splitChild(newRoot, 0)
		t.root = newRoot
	}

	t.insertNonFull(t.root, key)
}

func (t *BTree[T]) splitChild(parent *Node[T], i int) {
	child := parent.children[i]
	newNode := &Node[T]{leaf: child.leaf}

	midIndex := t.t - 1
	midKey := child.keys[midIndex]

	newNode.keys = append(newNode.keys, child.keys[midIndex+1:]...)
	child.keys = child.keys[:midIndex]

	if !child.leaf {
		newNode.children = append(newNode.children, child.children[midIndex+1:]...)
		child.children = child.children[:midIndex+1]
	}

	parent.children = append(parent.children, nil)
	copy(parent.children[i+2:], parent.children[i+1:])
	parent.children[i+1] = newNode

	parent.keys = append(parent.keys, midKey)
	copy(parent.keys[i+1:], parent.keys[i:])
	parent.keys[i] = midKey
}

func (t *BTree[T]) insertNonFull(n *Node[T], key T) {
	i := len(n.keys) - 1

	if n.leaf {
		n.keys = append(n.keys, key)
		for i >= 0 && key < n.keys[i] {
			n.keys[i+1] = n.keys[i]
			i--
		}
		n.keys[i+1] = key
	} else {
		for i >= 0 && key < n.keys[i] {
			i--
		}
		i++

		if len(n.children[i].keys) == 2 * t.t-1 {
			t.splitChild(n, i)
			if key > n.keys[i] {
				i++
			}
		}
		t.insertNonFull(n.children[i], key)
	}
}

func main() {
	btree := NewBTree[int](3)

	values := []int{10, 20, 5, 6, 12, 30, 7, 17}
	for _, v := range values {
		btree.Insert(v)
	}

	node, idx := btree.Search(6)
	if node != nil {
		fmt.Printf("Found: %d in node with keys %v (index %d)\n", node.keys[idx], node.keys, idx)
	} else {
		fmt.Println("Not found")
	}
}