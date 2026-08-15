package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet(msg string) string {
	return fmt.Sprintf("%s, says %s", p.Name, msg)
}

func main() {
	// 1
	x := 42
	v := reflect.ValueOf(x)
	fmt.Println(v.Int())
	fmt.Println(v.Kind())
	fmt.Println(v.Type())
	
	// 2
	original := v.Interface().(int)
	fmt.Println("Original from interface:", original)
	
	// 3
	y := 10.5
	vp := reflect.ValueOf(&y)
	ve := vp.Elem()
	
	if ve.CanSet() {
		ve.SetFloat(20.5)
		fmt.Println("New value:", y)
	}
	
	// 4
	p := Person{"Alice", 25}
	pv := reflect.ValueOf(p)
	
	nameField := pv.Field(0)
	fmt.Println("Field(0) Name", nameField.String())
	
	ageField := pv.FieldByName("Age")
	if ageField.CanSet() {
		ageField.SetInt(30)
	}
	fmt.Println("Updated struct:", p)
	
	// 5
	slice := []int{10, 20, 30}
	sv := reflect.ValueOf(slice)
	fmt.Println(sv.Len())
	fmt.Println(sv.Index(1).Int())
	
	// 6
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	mv := reflect.ValueOf(m)
	
	keyValue := reflect.ValueOf("one")
	resultValue := mv.MapIndex(keyValue)
	fmt.Println(resultValue.Int())
	
	mv.SetMapIndex(reflect.ValueOf("three"), reflect.ValueOf(3))
	
	keys := mv.MapKeys()
	fmt.Println("Map Keys:", keys)
	
	// 7
	method := pv.MethodByName("Greet")
	args := []reflect.Value{
		reflect.ValueOf("Hello Reflection!"),
	}
	results := method.Call(args)
	fmt.Println("Method Call Result:", results[0].String())
	
	// 8
	var strPtr *string
	vNil := reflect.ValueOf(strPtr)
	fmt.Println("Is Nil:", vNil.IsNil())
	
	vInvalid := reflect.ValueOf(nil)
	fmt.Println("Is Valid:", vInvalid.IsValid())
	
	// 9
	typeType := reflect.TypeOf(int(0))
	newIntPtr := reflect.New(typeType)
	newIntPtr.Elem().SetInt(999)
	fmt.Println("New Int Value:", newIntPtr.Elem().Int())
}