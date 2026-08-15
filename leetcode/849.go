package main

import (
	"fmt"
)

func maxDistToClosest(seats []int) int {
	// 1, 0, 0, 0, 1, 0, 1
	// 1, 0, 0, 0, 0, 1, 1
	// 0, 1, 0, 0, 0, 1, 1

	// Если встретил 0 — увеличивай zeros.
	// Если встретил 1:
	// Если это первая единица, расстояние = zeros.
	// Если не первая, расстояние = (zeros + 1) // 2.
	// Обнови max_dist и сбрось zeros = 0.
	
	if len(seats) == 2{
		return 1
	}

	zeros, maxDist := 0, 0
	curDif := 0
	countOne := 0

	for i := 0; i < len(seats); i++ {
		if seats[i] == 0 {
			zeros++
		} else {
			if countOne == 0 {
				curDif = zeros
				countOne++
			} else {
				curDif = (zeros + 1) / 2
				countOne = 0
			}

			if maxDist < curDif {
				maxDist = curDif
			}

			zeros = 0
		}
	}

	return maxDist
}

func main() {
	var seats []int

	seats = []int{1, 0, 0, 0, 1, 0, 1} // 2
	fmt.Println(maxDistToClosest(seats))
	seats = []int{1, 0, 0, 0} // 3
	fmt.Println(maxDistToClosest(seats))
	seats = []int{0, 1} // 1
	fmt.Println(maxDistToClosest(seats))
}
