package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Введите город: ")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	input = strings.TrimSpace(input)

	fmt.Printf("Ваш город: %s\n", input)
}