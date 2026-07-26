package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func FormatWithApostrophe(price int) string {
	if price <= 999 {
		fmt.Println(price)
	}
	
	res := strings.Builder{}
	
	for price > 0 {
		res.WriteString(strconv.FormatInt(int64(price % 1000), 10))
		res.WriteRune(' ')
		price /= 1000
	}

	res2 := strings.Split(res.String(), " ")
	slices.Reverse(res2)
	res3 := strings.Join(res2[1:], " ")
	
	return res3
}

func main() {
	fmt.Println(FormatWithApostrophe(12345678))
	// fmt.Println(12345678 % 1000, 12345678 / 1000)
}