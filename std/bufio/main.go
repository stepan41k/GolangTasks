package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func main() {
	inputFile, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	outputFile, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	reader := bufio.NewReader(inputFile)
	writer := bufio.NewWriter(outputFile)

	for {
		line, err := reader.ReadString('\n')

		if len(line) > 0 {
			processedLine := "LOG: " + line
			_, writeErr := writer.WriteString(processedLine)
			if writeErr != nil {
				log.Fatal(writeErr)
			}
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}
	}

	err = writer.Flush()
	if err != nil {
		log.Fatal(err)
	}
}
