package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	err := os.Truncate("output", 0)
	if err != nil {
		fmt.Println("failed to truncate file: %w", err)
		return
	}
	
	file, err := os.OpenFile("output", os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0644)
	defer func() {
		if err = file.Close(); err != nil {
			fmt.Println("failed to close file: %w", err)
		}
	}()

	if err != nil {
		fmt.Println("failed to open file: %w", err)
		return
	}
	
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(file)
	
	defer func() { 
		writer.Flush()
		fmt.Println(writer.Available(), writer.Buffered(), writer.Size())
	}()
	
	for range 10 {
		line, err := reader.ReadSlice(byte('\n'))
		if err != nil {
			fmt.Println("failed to read string")
		}
		
		_, err = writer.WriteString("write:" + string(line))
		if err != nil {
			fmt.Println("failed to write string")
		}
		fmt.Println(writer.Available(), writer.Buffered(), writer.Size())
	}

	writer.Flush()
}