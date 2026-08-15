package main

import (
	"bytes"
	"fmt"
)

func main() {
	var b bytes.Buffer

	b.Write([]byte("Hello, "))
	b.WriteString("world!")

	fmt.Println(b.String())

	newData := make([]byte, 12)
	b.Read(newData)

	b.Reset()

	fmt.Println(newData, string(newData))
}