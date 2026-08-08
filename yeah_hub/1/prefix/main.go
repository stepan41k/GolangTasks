package main

import (
	"fmt"
	"strings"
)

func WordsWithPrefix(s string, prefix string) []string {
	if len(s) == 0 {
		return []string{}
	}

	words := strings.Fields(s)
	res := []string{}

	for _, word := range words {
		if strings.HasPrefix(word, prefix) {
			res = append(res, word)
		}
	}

	return res
}

func main() {
	s := "ab abc def abc xyz ace ab cab"
	prefix := "ab"
	fmt.Println(WordsWithPrefix(s, prefix))

	s = "cat cat cat"
	prefix = ""
	fmt.Println(WordsWithPrefix(s, prefix))
}