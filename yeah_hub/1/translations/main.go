package main

import "fmt"

var translations = map[string]map[string]string{
	"en": {
		"hello":   "Hello",
		"welcome": "Welcome",
	},
	"ru": {
		"hello":   "Privet",
		"welcome": "Dobro pozhalovat",
	},
}

func getTranslationSafe(lang, key string) string {
	dict, ok := translations[lang]
	if !ok {
		return "[MISSING]"
	}

	trans, ok := dict[key]
	if !ok {
		return "[MISSING]"
	}

	return trans
}

func main() {
	fmt.Println(getTranslationSafe("en", "hello"))
}
