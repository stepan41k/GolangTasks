package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func processSingleFile(ctx context.Context, fileID string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		time.Sleep(1 * time.Second)
		fmt.Println("fetching:", fileID)
	}

	return nil
}

func ProcessFiles(ctx context.Context, fileIDs []string, maxConcurrent int) error {
	errGroup, ctx := errgroup.WithContext(ctx)
	
	filesChan := make(chan string, len(fileIDs))
	for _, v := range fileIDs {
		filesChan <- v
	}
	close(filesChan)

	for range maxConcurrent {
		errGroup.Go(func() error {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				for file := range filesChan {
					err := processSingleFile(ctx, file)
					if err != nil {
						return err
					}
				}
			}
			return nil
		})
	}

	if err := errGroup.Wait(); err != nil {
		return err
	}

	return nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	files := []string{"file1", "file2", "file3", "file4", "file5", "file6", "file7", "file8", "file9", "file10", "file11", "file12"}

	fmt.Println(ProcessFiles(ctx, files, 3))
}
