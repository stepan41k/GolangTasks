package main

import (
	"fmt"
	"strconv"
)

type Maybe[T any] struct {
	value *T
}

func Some[T any](val T) Maybe[T] {
	return Maybe[T]{value: &val}
}

func None[T any]() Maybe[T] {
	return Maybe[T]{value: nil}
}

func (m Maybe[T]) FlatMap[S any](f func(T) Maybe[S]) Maybe[S] {
	if m.value == nil {
		return Maybe[S]{}
	}
	return f(*m.value)
}

func (m Maybe[T]) GetOrElse(defaultValue T) T {
	if m.value == nil {
		return defaultValue
	}
	return *m.value
}

func main() {
	parseInt := func(s string) Maybe[int] {
		n, err := strconv.Atoi(s)
		if err != nil {
			return None[int]()
		}
		return Some(n)
	}

	checkAge := func(age int) Maybe[string] {
		if age >= 18 {
			return Some(fmt.Sprintf("Доступ разрешен (возраст %d)", age))
		}
		return None[string]()
	}
	
	input1 := Some("25")
	res1 := input1.FlatMap(parseInt).FlatMap(checkAge)
	fmt.Println("Тест 1:", res1.GetOrElse("Ошибка / Доступ запрещен"))

	input2 := Some("abc")
	res2 := input2.FlatMap(parseInt).FlatMap(checkAge)
	fmt.Println("Тест 2:", res2.GetOrElse("Ошибка / Доступ запрещен"))

	input3 := Some("15")
	res3 := input3.FlatMap(parseInt).FlatMap(checkAge)
	fmt.Println("Тест 3:", res3.GetOrElse("Ошибка / Доступ запрещен"))

	input4 := None[string]()
	res4 := input4.FlatMap(parseInt).FlatMap(checkAge)
	fmt.Println("Тест 4:", res4.GetOrElse("Ошибка / Доступ запрещен"))
}