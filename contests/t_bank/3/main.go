// Вы получили доступ к секретной информации о будущих ценах акций одной компании на ближайшие n дней. Для каждого дня заранее известна точная стоимость акций.
// Вы планируете воспользоваться этой информацией и совершить ровно две сделки с одной акцией. Каждая сделка состоит из покупки и последующей продажи одной акции. Формально:
// вы покупаете акцию в день a и продаёте её в день b, где a <= b;
// затем вы снова покупаете акцию в день c и продаёте её в день d, где b <= c <= d.
// Таких образов, вы не можете держать более одной акции одновременно а вторая сделка начинается не раньше завершенной первой.
// Требуется определить максимальную суммарную прибыль, которую можно получить при оптимальном использовании этой информации.
// Входные данные:
// Первая строка содержит целое число n (1 <= n <= 5 * 10^5) - количество дней
// Вторая строка содержит n целых чисел p1, p2, ..., pn (1 <= pi <= 10^9) - цену акций в каждый из дней.
// Выходные данные:Выведите одно целое число - максимальную возможную прибыль

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	if !scanner.Scan() {
		return
	}
	n, _ := strconv.Atoi(scanner.Text())

	prices := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		prices[i], _ = strconv.Atoi(scanner.Text())
	}

	if n < 2 {
		fmt.Println(0)
		return
	}

	leftProfit := make([]int64, n)
	minPrice := int64(prices[0])
	var maxP int64 = 0
	for i := 1; i < n; i++ {
		curPrice := int64(prices[i])
		if curPrice < minPrice {
			minPrice = curPrice
		}
		if curPrice-minPrice > maxP {
			maxP = curPrice - minPrice
		}
		leftProfit[i] = maxP
	}

	rightProfit := make([]int64, n)
	maxPrice := int64(prices[n-1])
	maxP = 0
	for i := n - 2; i >= 0; i-- {
		curPrice := int64(prices[i])
		if curPrice > maxPrice {
			maxPrice = curPrice
		}
		if maxPrice-curPrice > maxP {
			maxP = maxPrice - curPrice
		}
		rightProfit[i] = maxP
	}

	var totalMaxProfit int64 = 0
	for i := 0; i < n; i++ {
		sum := leftProfit[i] + rightProfit[i]
		if sum > totalMaxProfit {
			totalMaxProfit = sum
		}
	}

	fmt.Println(totalMaxProfit)
}