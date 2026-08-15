package main

import (
	"fmt"
	"reflect"
)

type Printer interface {
	PrintName()
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u User) PrintName() {
	fmt.Println(u.Name)
}

func main() {
	u := User{ID: 1, Name: "41k", Email: "stepson@example.com"}

	t := reflect.TypeOf(u)

	// 1
	tfk := t.Kind()
	fmt.Println(tfk)

	// 2
	tfn := t.Name()
	fmt.Println(tfn)

	// 3
	tfpp := t.PkgPath()
	fmt.Println(tfpp)

	// 4
	tfmf := t.NumField()
	fmt.Println(tfmf)

	// 5
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("Field [%d]: %s, Type: %v, Тег: %v\n", i, field.Name, field.Type, field.Tag)
	}

	// 6
	if nameField, found := t.FieldByName("Email"); found {
		fmt.Printf("found fields 'Email': тип %v\n", nameField.Type)
	}
	
	printerInterface := reflect.TypeOf((*Printer)(nil)).Elem()
	
	// 7
	fmt.Printf("User реализует интейрфейс Printer? %v\n", t.Implements(printerInterface))
	
	// 8
	fmt.Printf("User можно присвоить переменной типа Printer? %v\n", t.AssignableTo(printerInterface))
	
	// 9
	intType := reflect.TypeOf(10)
	fmt.Printf("User можно присвоить переменной типа int? %v\n", t.AssignableTo(intType))
}
