package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func intervalArray(source []int) string {
	if len(source) == 0 {
		return ""
	}

	nums := append([]int(nil), source...)

	slices.Sort(nums)

	nums = slices.Compact(nums)

	var builder strings.Builder
	fPtr := 0

	for sPtr := 1; sPtr <= len(nums); sPtr++ {
		if sPtr == len(nums) || nums[sPtr]-nums[sPtr-1] != 1 {
			if builder.Len() > 0 {
				builder.WriteByte(',')
			}

			if (sPtr - 1) > fPtr {
				builder.WriteString(strconv.Itoa(nums[fPtr]))
				builder.WriteByte('-')
				builder.WriteString(strconv.Itoa(nums[sPtr-1]))
			} else {
				builder.WriteString(strconv.Itoa(nums[fPtr]))
			}

			fPtr = sPtr
		}
	}

	return builder.String()
}

func main() {
	fmt.Println(intervalArray([]int{1, 4, 5, 2, 3, 9, 8, 11, 0})) // 0, 1, 2, 3, 4, 5, 8, 9, 11   "0-5,8-9,11"
	fmt.Println(intervalArray([]int{}))
	fmt.Println(intervalArray([]int{1}))
	fmt.Println(intervalArray([]int{1, 2}))
	fmt.Println(intervalArray([]int{1, 3}))
	fmt.Println(intervalArray([]int{1, 3, 5}))
	fmt.Println(intervalArray([]int{1, 2, 3, 4, 5, 6}))
}
