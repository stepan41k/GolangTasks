// Организаторы мероприятия вводят новый класс строк, используемых в олимпиадных задачах.
// Строка называется TOI-строкой, если выполняются следующие условия:
// она состоит только из символов T, O, I
// она является палиндромом
// Дана произвольная строка из больших букв латинского алфавита, требуется преобразовать её в TOI-строку.
// За одну операцию разрешается выполнить одно из следующих действий:
// заменить один символ строки на другой
// удалить один символ
// вставить один символ в произвольную позицию.
// Требуется определить минимальное количество операций, необходимых для преобразования данной строки в TOI-строку.
// Входные данные:
// В первой строке - целое число n (1 <= n <= 1000) -  длина строки
// Во второй строке задана строка s длины n из больших букв латинского алфавита.
// Выход:
// Одно целое число - минимальное количество операций

package main

import (
	"fmt"
)

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		return
	}

	var s string
	if _, err := fmt.Scan(&s); err != nil {
		return
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	toi := "TOI"

	for length := 1; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1

			if length == 1 {
				cost := 1
				for k := 0; k < 3; k++ {
					if s[i] == toi[k] {
						cost = 0
						break
					}
				}
				dp[i][j] = cost
				continue
			}

			res := dp[i+1][j] + 1
			if dp[i][j-1]+1 < res {
				res = dp[i][j-1] + 1
			}

			for k := 0; k < 3; k++ {
				char := toi[k]
				currentCost := 0
				
				if s[i] != char {
					currentCost++
				}
				if s[j] != char {
					currentCost++
				}
				
				inner := 0
				if i+1 <= j-1 {
					inner = dp[i+1][j-1]
				}
				
				if currentCost+inner < res {
					res = currentCost + inner
				}
			}
			dp[i][j] = res
		}
	}

	fmt.Println(dp[0][n-1])
}