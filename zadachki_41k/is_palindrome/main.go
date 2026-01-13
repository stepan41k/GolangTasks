package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(str string) bool {
	if len(str) == 0 {
		 return true
	}

	str = strings.ToLower(str)

	runeStr := []rune(str)

	left := 0
	right := len(runeStr)-1

	for left <= right {
		if unicode.IsLetter(runeStr[left]) && unicode.IsLetter(runeStr[right]) {
			if runeStr[left] == runeStr[right] {
				left++
				right--
				continue
			} else {
				return false
			}
		}

		if !unicode.IsLetter(runeStr[left]) {
			left++
		}

		if !unicode.IsLetter(runeStr[right]) {
			right--
		}
	}

	return true
}


func main() {
	str := "А роза упала на лапу Азора!"
	
	fmt.Println(isPalindrome(str))

	str = ""
	
	fmt.Println(isPalindrome(str))

	str = "!!"
	
	fmt.Println(isPalindrome(str))

	str = "АбА"
	
	fmt.Println(isPalindrome(str))

	str = ""
	
	fmt.Println(isPalindrome(str))
}