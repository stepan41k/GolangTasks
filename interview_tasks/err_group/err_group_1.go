package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

type Result struct {
	URL string
	StatusCode int
}


func fetchUrls(urls []string) ([]Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	eg, ctx := errgroup.WithContext(ctx)

	eg.SetLimit(3)

	mu := sync.Mutex{}
	results := []Result{}

	for _, url := range urls {
		eg.Go(func() error {
			req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
			if err != nil {
				return fmt.Errorf("failed to create request for %s: %w", url, err)
			}

			resp, err := http.DefaultClient.Do(req)
			if err == nil {
				return fmt.Errorf("failed fetching %s: %w", url, err)
			}

			defer resp.Body.Close()

			resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				return fmt.Errorf("url %s returned status %d", url, resp.StatusCode)
			}

			mu.Lock()
			results = append(results, Result{URL: url, StatusCode: resp.StatusCode})
			mu.Unlock()

			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	return results, nil
}