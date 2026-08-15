package main

import (
	"fmt"
	"iter"
)

type User struct {
	ID int
	Active bool
}

func ActiveUsers(users []*User) iter.Seq[*User] {
	return func(yield func(*User) bool) {
		for _, u := range users {
			if u.Active {
				if !yield(u) {
					return
				}
			}
		}
	}
}

func main() {
	users := []*User{
		{ID: 1, Active: true},
		{ID: 2, Active: false},
		{ID: 3, Active: true},
	}

	for u := range ActiveUsers(users) {
		fmt.Println("Active user ID:", u.ID)
	}
}