package main

import (
	"io"
	"os"
	"sync"
)

func main() {
	inputFile, err := os.OpenFile("input", os.O_RDONLY, 0o644)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	outputFile, err := os.OpenFile("output", os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0o644)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	pool := sync.Pool{New: func() any {return make([]byte, 32 * 1024)}}

	buf := pool.Get().([]byte)
	defer pool.Put(buf)


	_, _ = io.CopyBuffer(outputFile, inputFile, buf)
}
