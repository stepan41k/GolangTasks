package main

import (
	"fmt"
	"sync"
	"time"
)

func downloadFile(filename string) {
	fmt.Printf("downloading %s\n", filename)
	time.Sleep(1 * time.Second)
	fmt.Println("downloaded")
}

func main() {
	files := []string{"file1", "file2", "file3", "file4", "file5", "file6", "file7"}
	const goroutinesLimit = 3
	wg := sync.WaitGroup{}
	semaphore := make(chan struct{}, goroutinesLimit)

	wg.Add(len(files))
	for _, file := range files {
		semaphore <- struct{}{}
		go func() {
			defer func() {
				<-semaphore
				wg.Done()
			}()

			downloadFile(file)
		}()
	}

	wg.Wait()
}
