package main

import (
	"bufio"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)

	writer.WriteString("fastest output through the bufio\n")
	writer.WriteString("one more string\n")

	// ОЧЕНЬ ВАЖНО: вызвать Flush, чтобы данные из буфера реально отправились в консоль
	writer.Flush()
}