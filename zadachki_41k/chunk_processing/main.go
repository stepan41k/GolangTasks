// При отмене ctx или первой возникшей ошибке в processSingleChunk:
// Прекратить обработку новых чанков.
// Дождаться завершения уже работающих горутин.
// Вернуть ошибку наверх.
// Без Race Condition и без утечек горутин/каналов.
package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

type Chunk struct {
	ID   int
	Data []byte
}

type Result struct {
	ID   int
	Hash string
}

func processSingleChunk(ctx context.Context, chunk Chunk) (Result, error) {
	fmt.Println(chunk.ID)
	time.Sleep(5 * time.Second)
	if chunk.ID == 4 {
		return Result{}, fmt.Errorf("FAILED TO PRCESS CHUNK")
	}
	return Result{ID: chunk.ID, Hash: "hash"}, nil
}

func ProcessChunks(ctx context.Context, chunks []Chunk, maxWorkers int) ([]Result, error) {
	results := make([]Result, len(chunks))

	errGroup, ctx := errgroup.WithContext(ctx)
	errGroup.SetLimit(maxWorkers)

	for index, chunk := range chunks {
		index := index
		chunk := chunk
		if err := ctx.Err(); err != nil {
			break
		}

		errGroup.Go(func() error {
			res, err := processSingleChunk(ctx, chunk)
			if err != nil {
				return err
			}

			results[index] = res

			return nil
		})

	}

	if err := errGroup.Wait(); err != nil {
		return results, err
	}

	return results, nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	chunk1 := Chunk{ID: 1, Data: []byte("first chunk")}
	chunk2 := Chunk{ID: 2, Data: []byte("second chunk")}
	chunk3 := Chunk{ID: 3, Data: []byte("third chunk")}
	chunk4 := Chunk{ID: 4, Data: []byte("fourth chunk")}
	chunk5 := Chunk{ID: 5, Data: []byte("fifth chunk")}
	chunk6 := Chunk{ID: 6, Data: []byte("sixth chunk")}
	chunk7 := Chunk{ID: 7, Data: []byte("seventh chunk")}

	chunks := []Chunk{chunk1, chunk2, chunk3, chunk4, chunk5, chunk6, chunk7}

	result, err := ProcessChunks(ctx, chunks, 3)
	if err != nil {
		fmt.Println("ERROR:", err.Error())
	}

	fmt.Print(result, err)
}
	