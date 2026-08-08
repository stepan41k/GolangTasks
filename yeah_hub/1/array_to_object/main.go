package main

import (
	"fmt"
)

func ArrayToObject(arr []map[string]interface{}) map[string]interface{} {
	res := make(map[string]interface{}, len(arr))
	
	for i := 0; i < len(arr); i++ {
		elem := arr[i]
		key := ""
		var val interface{}
		
		for k, v := range elem {
			if k == "name" {
				key = v.(string)
			} else {
				val = v
			}
		}

		res[key] = val
	}
	
	return res
}

func main() {
	fmt.Println((ArrayToObject([]map[string]interface{}{})))
}