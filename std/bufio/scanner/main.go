package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Введите ваше имя: ")

	if scanner.Scan() {
		name := scanner.Text()
		fmt.Printf("Привет, %s!\n", name)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка чтения:", err)
	}
}