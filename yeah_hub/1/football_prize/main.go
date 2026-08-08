package main

import "fmt"

func getPrize(guessScore string, realScore string) int {
	var bX, bY, rX, rY int

	fmt.Sscanf(guessScore, "%d:%d", &bX, &bY)
	fmt.Sscanf(realScore, "%d:%d", &rX, &rY)
	
    if rX == bX && rY == bY {
    	return 2
    }

    if (rX > rY && bX > bY) || (rX < rY && bX < bY) || (rX == rY && bX == bY) {
    	return 1
    }
    
    return 0
}

func main() {
	println(getPrize("1:2", "1:2"))  // -> 2 (точный счёт)
	println(getPrize("2:1", "5:0"))  // -> 1 (исход: победа А)
		println(getPrize("3:0", "2:2"))  // -> 0 (не угадал)
}
