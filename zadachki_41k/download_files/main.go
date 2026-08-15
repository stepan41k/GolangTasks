package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

type DiskDownloader struct {
    client *http.Client
}

func (d *DiskDownloader) GetDownloadFile(ctx context.Context, url string) ([]byte, error) {
    ctx, cancel := context.WithTimeout(ctx, 30 * time.Second)
    defer cancel()
    
    req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
    if err != nil {
        return nil, err
    }

    resp, err := d.client.Do(req)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode != 200 {
        return nil, fmt.Errorf("bad status code: %d", resp.StatusCode)
    }

    data, err := io.ReadAll(resp.Body) // Проблема 4?
    if err != nil {
        return nil, err
    }

    return data, nil
}