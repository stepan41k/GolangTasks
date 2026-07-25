package main

import "fmt"

type Person struct {
    Name string
}

func list(names []Person) string {
	if len(names) == 0 {
		return ""
	}

	if len(names) == 1 {
		return names[0].Name
	}

	res := ""
	
	for i, v := range names {
		if i == len(names) - 1 {
			res += string('&') + v.Name
			continue
		}

		res += string(',') + v.Name
	}
	
    return res
}

func main() {
	fmt.Println(list([]Person{Person{Name:"Bart"}, Person{Name: "Lisa"}, Person{Name: "Maggie"}}))
}