package main

import "fmt"

type MyStack struct {
	stack   []int
	maxStack []int
}

func Constructor() MyStack {
	return MyStack{
		stack: []int{},
		maxStack: []int{},
	}
}

func (s *MyStack) Push(val int) {
    s.stack = append(s.stack, val)
    
    if len(s.maxStack) == 0 {
        s.maxStack = append(s.maxStack, val)
    } else {
        currentMax := s.maxStack[len(s.maxStack)-1]
        if val > currentMax {
            s.maxStack = append(s.maxStack, val)
        } else {
            s.maxStack = append(s.maxStack, currentMax)
        }
    }
}

func (s *MyStack) Pop() interface{} {
    if len(s.stack) == 0 {
        return nil
    }
    
    topIdx := len(s.stack) - 1
    val := s.stack[topIdx]
    
    s.stack = s.stack[:topIdx]
    s.maxStack = s.maxStack[:topIdx]
    
    return val
}

func (s *MyStack) Max() interface{} {
	if len(s.maxStack) == 0 {
		return nil
	}

	return s.maxStack[len(s.maxStack) - 1]
}

func main() {
	s := Constructor()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	fmt.Println("max", s.Max())
	fmt.Println(s.Pop())
	fmt.Println("max", s.Max())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}
