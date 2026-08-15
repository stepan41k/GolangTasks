package main

import (
	"fmt"
	"sync"
)

// Тебе нужно написать функцию, которая параллельно вычисляет квадраты чисел и возвращает их общую сумму.

// Условие:
// Реализуй функцию SumSquares(tasks []int, workerCount int) int.
// Воркеры: Нужно запустить ровно workerCount горутин.
// Каналы:
	// jobs — канал для передачи чисел из слайса tasks воркерам.
	// results — канал, куда воркеры будут складывать результаты (квадраты чисел).
// Синхронизация: Используй sync.WaitGroup, чтобы дождаться окончания работы всех воркеров.
// Сбор данных: Основной поток должен просуммировать все значения из канала results и вернуть итоговое число.
// Важный нюанс:
// Middle+ разработчик должен следить за тем, чтобы не возникло Deadlock (взаимной блокировки). Подумай, в какой момент нужно закрывать каналы и где именно делать wg.Wait().

// На что я буду смотреть:
// Как ты распределяешь задачи.
// Где закрываешь каналы (это самое частое место ошибок).
// Не блокируется ли выполнение, если буфер канала закончится.
// Корректность использования WaitGroup.

func SumSquares(tasks []int, workerCount int) int {
	results := make(chan int, workerCount)
	jobs := make(chan int, workerCount)
	wg := sync.WaitGroup{}

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range jobs {
				results <- v * v
			}
			
		}()
	}

	go func() {
		for _, v := range tasks {
			jobs <- v
		}

		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	total := 0
	for v := range results {
		total += v
	}

    return total
}

func main() {
	tasks := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	workerCount := 3

	result := SumSquares(tasks, workerCount)
	fmt.Printf("Total Sum: %d\n", result) // Ожидаем 385
}



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
