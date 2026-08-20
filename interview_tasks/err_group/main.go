package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// err_group_2.go
	ids := make([]int, 100)
	for i := range ids {
		ids[i] = i + 1
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	users, err := ProcessUsers(ctx, ids, 10)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Successfully fetched %d users\n", len(users))

	// err_group_1.go
	// urls := []string{
	// 	"https://google.com",
	// 	"https://github.com",
	// 	"https://bad-url-that-fails.com",
	// 	"https://go.dev",
	// 	"https://yandex.ru",
	// }

	// results, err := fetchUrls(urls)
	// if err != nil {
	// 	fmt.Printf("Error occurred: %v\n", err)
	// 	return
	// }

	// fmt.Println(results)
}
