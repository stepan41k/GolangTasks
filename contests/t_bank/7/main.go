// Компания планирует нанять a бекенд-разработчиков и b ml-инженеров. Всего на позиции претендуют n кандидатов, и каждый из них может быть принят либо на позицию backend, либо на позицию ml.
// Для каждого кандидата известны две оценки: x - уровень backend и y - уровень ml
// Каждого кандидата можно назначить только на одну из двух ролей.
// Требуется выбрать сотрудников так, чтобы:
// было выбрано ровно a backend разработчиков
// было выбрано ровно b ml инженеров
// каждый кандидат выбран не более одного раза
// При этом нужно максимизировать сумму соответствующих навыков
// Вход:
// В первой строке три целых числа a, b, n (1 <= n <= 2 * 10^5, 0 <= a, b <= n, a + b <= n) -  количество  backend разработчиков, ml - инженеров и общее число кандидатов.
// Далее следует n строк, каждая из которых содержит два целых числа x и y (0 <= x, y <= 10^9) - соответсвующие навыки кандидата.
// Выход:
// Одно целое число - максимальную суммарную сумму навыков

package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

type Candidate struct {
	x, y int64
	d    int64
}

type Int64Heap []int64

func (h Int64Heap) Len() int           { return len(h) }
func (h Int64Heap) Less(i, j int) bool { return h[i] < h[j] }
func (h Int64Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *Int64Heap) Push(x interface{}) {
	*h = append(*h, x.(int64))
}
func (h *Int64Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var a, b, n int
	fmt.Fscan(reader, &a, &b, &n)

	candidates := make([]Candidate, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &candidates[i].x, &candidates[i].y)
		candidates[i].d = candidates[i].x - candidates[i].y
	}

	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].d > candidates[j].d
	})

	pref := make([]int64, n+1)
	if a > 0 {
		h := &Int64Heap{}
		heap.Init(h)
		var currentSum int64
		for i := 0; i < n; i++ {
			val := candidates[i].x
			heap.Push(h, val)
			currentSum += val
			if h.Len() > a {
				smallest := heap.Pop(h).(int64)
				currentSum -= smallest
			}
			if h.Len() == a {
				pref[i+1] = currentSum
			}
		}
	}

	suff := make([]int64, n+1)
	if b > 0 {
		h := &Int64Heap{}
		heap.Init(h)
		var currentSum int64
		for i := n - 1; i >= 0; i-- {
			val := candidates[i].y
			heap.Push(h, val)
			currentSum += val
			if h.Len() > b {
				smallest := heap.Pop(h).(int64)
				currentSum -= smallest
			}
			if h.Len() == b {
				suff[i] = currentSum
			}
		}
	}

	var maxTotal int64 = 0

	if a == 0 {
		maxTotal = suff[0]
	} else if b == 0 {
		maxTotal = pref[n]
	} else {
		for k := a; k <= n-b; k++ {
			maxTotal = max(maxTotal, pref[k]+suff[k])
		}
	}

	fmt.Println(maxTotal)
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
