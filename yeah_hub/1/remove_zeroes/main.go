package main

import "fmt"

func RemoveZeros(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	find, sind := 0, len(nums)-1
	for find <= sind {
		if nums[find] == 0 && nums[sind] != 0 {
			nums[find], nums[sind] = nums[sind], nums[find]
			find++
			sind--
		} else if nums[find] != 0 && nums[sind] == 0 {
			find++
			sind--
		} else if nums[find] == 0 && nums[sind] == 0 {
			sind--
		} else {
			find++
		}
	}

	return nums[:sind+1]
}

func main() {
	fmt.Println(RemoveZeros([]int{1, 0, 0, 2}))
	fmt.Println(RemoveZeros([]int{1, 2, 0, 0}))
	fmt.Println(RemoveZeros([]int{1, 1, 0, 2}))
	fmt.Println(RemoveZeros([]int{1, 1, 1, 2}))
	fmt.Println(RemoveZeros([]int{0, 0, 0, 0}))
	fmt.Println(RemoveZeros([]int{}))
}