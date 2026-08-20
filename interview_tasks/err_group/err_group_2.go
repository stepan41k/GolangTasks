package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

// Graceful Shutdown / Ошибки: Если один из запросов вернул критическую ошибку, нужно отменить все остальные текущие запросы и вернуть ошибку немедленно.

type User struct {
	ID   int
	Name string
}

func fetchUser(ctx context.Context, id int) (*User, error) {
	select {
	case <-time.After(time.Millisecond * 500):
		return &User{ID: id, Name: fmt.Sprintf("User-%d", id)}, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func ProcessUsers(ctx context.Context, ids []int, concurrency int) ([]*User, error) {
	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(concurrency)

	res := make([]*User, len(ids))

	for idx, id := range ids {
		id := id

		g.Go(func() error {
			user, err := fetchUser(ctx, id)
			if err != nil {
				return err
			}

			res[idx] = user
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return res, nil
}
