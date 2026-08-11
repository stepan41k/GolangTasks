// На уроке русского языка Саша получил строку S, состоящую из строчных  английских букв. Учитель попросил его превратить эту строку в палиндром, но разрешил выполнять только одно действие : дописывать некоторое количество букв "a" в начало строки. Количество добавленных букв может быть и нулевым.
// Помогите Саше определить, можно ли таким способом получить палиндром.
// Входные данные: В единственной строку входных данных дана строка S(1 <= | S |  <= 10^6) из строчных латинских букв

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const maxCapacity = 1000005
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	if !scanner.Scan() {
		return
	}
	s := strings.TrimSpace(scanner.Text())
	n := len(s)

	rightA := 0
	for i := n - 1; i >= 0 && s[i] == 'a'; i-- {
		rightA++
	}

	if rightA == n {
		fmt.Println("Yes")
		return
	}

	leftA := 0
	for i := 0; i < n && s[i] == 'a'; i++ {
		leftA++
	}

	if leftA > rightA {
		fmt.Println("No")
		return
	}

	target := s[leftA : n-rightA]
	
	if isPalindrome(target) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}