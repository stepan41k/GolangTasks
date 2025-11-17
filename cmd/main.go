package main

import (
	"fmt"
)


func main() {
	// done := make(chan struct{})

	// for i := 0; i < 5; i++ {
    // 	go func(i int) {
    //     	fmt.Println("work", i)
    //     	done <- struct{}{}
    // 	}(i)

    // <-done // ждём завершения горутины i
	// }

	n := 5
	signals := make([]chan struct{}, n)

	for i := 0; i < n; i++ {
		signals[i] = make(chan struct{})
	}

	for i := 0; i < n; i++ {
		go func(i int) {
			if i > 0 {
				<-signals[i-1]
			}
			fmt.Println("work", i)
			close(signals[i])
		}(i)
	}

	<-signals[n-1]
}