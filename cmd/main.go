package main

import (
	"fmt"
)

type Struct1 struct {
	Field1 string
}

type Struct2 struct {
	Struct1
	Field2 string
}

func (s1 Struct1) Method1() {
	fmt.Println("Hello from method1")
}

func (s2 Struct2) Method1() {
	fmt.Println("Hello from method2")
}	

func main() {
	struct2 := Struct2{
	}

	struct2.Struct1.Method1()
}