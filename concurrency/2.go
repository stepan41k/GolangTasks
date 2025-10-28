package main

import (
	"context"
	"fmt"
	"time"
)

const defaultTimeout = 1 * time.Second

func getDiscount() float64 {
	time.Sleep(2 * time.Second)
	return 12.0
}

func getDiscountWithTimeout(ctx context.Context) (float64, error) {
	ch := make(chan float64)

	go func() {
		res := getDiscount()
		ch <- res
		close(ch)
	}()

	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, defaultTimeout)
		defer cancel()
	}

	to, _ := ctx.Deadline()
	fmt.Println(time.Since(to))

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case val := <-ch:
		return val, nil
	}
}

<<<<<<< HEAD
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()
	
	res, err := getDiscountWithTimeout(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
}
=======
// func main() {
// 	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
// 	defer cancel()
	
// 	res, err := getDiscountWithTimeout(ctx)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Println(res)
// }
>>>>>>> d4945cb (change files)
