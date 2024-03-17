package main

import (
	"fmt"
	"sort"
)

func removeElement(nums []int, val int) int {
	cnt := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums[i] = 101
			cnt++
		}
	}

	sort.Ints(nums)

	return len(nums) - cnt
}

func RemoveElementTest() {
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}

// func main() {
// 	RemoveElementTest()
// }
