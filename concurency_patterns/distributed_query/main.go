package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

var ErrNotFound = errors.New("not found")

type queryResult struct {
	res string
	err error
}

type DatabaseHost interface {
	DoQuery(ctx context.Context, query string) (string, error)
}

type DatabaseInstance struct {
	Number int
}

func (di *DatabaseInstance) DoQuery(ctx context.Context, query string) (string, error) {
	select {
	case <-time.After(500 * time.Millisecond):
		if query == "SELECT 1" {
			return "", ErrNotFound
		}
		return fmt.Sprintf("Data from replica %d", di.Number), nil

	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func DistributedQuery(query string, replicas []DatabaseHost) (string, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	resChan := make(chan queryResult, len(replicas))
	wg := &sync.WaitGroup{}

	for _, replica := range replicas {
		wg.Add(1)
		go func(r DatabaseHost) {
			defer wg.Done()

			res, err := r.DoQuery(ctx, query)
			if err != nil && errors.Is(err, ErrNotFound) {
					return 
			}
			
			resChan <- queryResult{res: res, err: err}

			cancel()
		}(replica)
	}

	go func() {
		wg.Wait()
		close(resChan)
	}()


	result, ok := <-resChan
	if !ok {
		return "", ErrNotFound
	}

	if result.err != nil {
		return "", result.err
	}

	return result.res, result.err
}

func main() {

	dbs2 := []DatabaseHost{&DatabaseInstance{Number: 1}, &DatabaseInstance{Number: 2}, &DatabaseInstance{Number: 3}, &DatabaseInstance{Number: 4}}

	// fmt.Println(DistributedQuery("SELECT 1", dbs2))

	fmt.Println(DistributedQuery("SELECT 2", dbs2))
}	