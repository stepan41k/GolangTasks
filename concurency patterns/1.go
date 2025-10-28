package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)


func main() {
	channels := make([]chan int64, 1000)
	for i := range channels {
		channels[i] = make(chan int64)
	}

	for i := range channels {
		go func(i int) {
			channels[i] <- int64(i)
			close(channels[i])
		}(i)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for v := range merge(ctx, channels...) {
		fmt.Println(v)
	}
}

func merge(ctx context.Context, channels ...chan int64) chan int64 {
	resChan := make(chan int64)
	
	merger := func(ch chan int64) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("context done")
				return
			case val, ok := <-ch:
				if !ok {
					return
				}

				resChan <- val
			}
		}

		
	}

	wg := sync.WaitGroup{}
	wg.Add(len(channels))
	for _, channel := range channels {
		go func(){
			defer wg.Done()
			merger(channel)
		}()
	}

	go func() {
		wg.Wait()
		close(resChan)
	}()
	
	return resChan
}




func merge(ch ...<-chan int) <-chan int {
	out := make(chan int)
	wg := &sync.WaitGroup{}


	wg.Add(len(ch))
	for _, c := range ch {
		go func() {
			defer wg.Done()
			
			for v := range c {
				out <- v
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func source(sourceFunc func(int) int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := 0; i < 10; i++ {
			ch <- sourceFunc(i)
			time.Sleep(time.Duration(rand.Intn(3) * int(time.Second)))
		}
	}()
	return ch
}

func main() {
	rand.NewSource(time.Now().UnixMilli())

	in1 := source(func(_ int) int{
		return rand.Int()
	})

	in2 := source(func(i int) int {
		return i
	})

	out := merge(in1, in2)

	for val := range out {
		fmt.Println(val)
	}
}





import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	urls := []string{
		"https://www.lamoda.ru",
		"https://www.yandex.ru",
		"https://www.mail.ru",
		"https://www.google.com",
	}

	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg.Add(len(urls))
	for _, url := range urls {
		u := url
		go func(ctx context.Context, currentURL string) {
			defer wg.Done()

			select {
			case <-ctx.Done():
				fmt.Printf("Горутина для %s отменена до старта: %v\n", currentURL, ctx.Err())
				return
			default:
			}

			fmt.Printf("Получаем данные с %s...\n", currentURL)

			time.Sleep(50 * time.Millisecond)

			err := fetchUrl(currentURL)
			if err != nil {
				if ctx.Err() == nil {
					cancel()
					fmt.Printf("Ошибка при получении %s: %v. Отменяем другие операции.\n", currentURL, err)
				} else {
					fmt.Printf("Ошибка при получении %s (контекст уже отменен): %v\n", currentURL, err)
				}
				return
			}

			select {
			case <-ctx.Done():
				fmt.Printf("Горутина для %s завершила получение, но контекст был отменен: %v\n", currentURL, ctx.Err())
				return
			default:
				fmt.Printf("Успешно получено с %s\n", currentURL)
			}
		}(ctx, u)
	}

	wg.Wait()

	fmt.Println("Все горутины завершили работу (или были отменены).")
	fmt.Println("Программа завершена.")
}

func fetchUrl(url string) error {
	_, err := http.Get(url)
	if err != nil {
		return err
	}
	
	return nil
}