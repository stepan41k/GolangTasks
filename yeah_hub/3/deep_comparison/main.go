package main

import (
	"fmt"
	"reflect"
)

func deepCompare(array1, array2 []interface{}) bool {
	if len(array1) != len(array2) {
		return false
	}

	for i := 0; i < len(array1); i++ {
		v1 := array1[i]
		v2 := array2[i]
		
		if reflect.TypeOf(array1[i]) != reflect.TypeOf(array2[i]) {
			return false
		}

		val1 := reflect.ValueOf(v1)
		val2 := reflect.ValueOf(v2)

		if v1 != nil && (val1.Kind() == reflect.Slice || val1.Kind() == reflect.Array) {
			if val1.Len() != val2.Len() {
				return false
			}
		} else {
			if v1 != v2 {
				return false
			}
		}
	}

	return true
}

func main() {
	array1 := []interface{}{1, 2, 3}
	array2 := []interface{}{1, 2, 3}
	fmt.Println(deepCompare(array1, array2))
	// Выходtrue
	// Пояснениедлины равны, типы совпадают, значения равны

	array1 = []interface{}{1, []any{5, 7}}
	array2 = []interface{}{1, []any{2, 2}}
	fmt.Println(deepCompare(array1, array2))
	// Выходtrue
	// Пояснениедлины равны, оба элемента - массивы, длины вложенных равны

	array1 = []interface{}{1, []any{8, 1}}
	array2 = []interface{}{[]any{20, 2}, 2}
	fmt.Println(deepCompare(array1, array2))
	// Выходfalse
	// Пояснениена позиции 0 тип не совпадает (примитив vs массив)

	array1 = []interface{}{1, []any{1, 10}}
	array2 = []interface{}{1, []any{4}}
	fmt.Println(deepCompare(array1, array2))
}
