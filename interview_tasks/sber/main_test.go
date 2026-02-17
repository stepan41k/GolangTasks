package main

import (
	"testing"
)

func Test_NewObject(t *testing.T) {
	tests := []struct {
		title string
		str string
		result Sayer
	}{
		{
			title: "Success base",
			str: "Base",
			result: Base{name: "Parent"},
		
		},
		{
			title: "Success child",
			str: "Child",
			result: Child{lastName: "Inherited", Base: Base{name: "Child"}},
		},
		{
			title: "Invalid argument",
			str: "abraabrabra",
			result: nil,
		},
	}

	for _, v := range tests {
		t.Run(v.title, func(t *testing.T){
			res := NewObject(v.str)
			if res != v.result {
				t.Fatalf("error not equal: expected %s, found: %s", v.result, res)
			}
			
		})
	}
}