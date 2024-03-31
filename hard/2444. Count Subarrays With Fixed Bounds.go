package main

import "fmt"

func countSubarrays(nums []int, minK int, maxK int) int64 {
	left := 0
	minIndex := -1
	maxIndex := -1
	var res int64 = 0

	for right, num := range nums {
		if num < minK || num > maxK {
			left = right + 1
			minIndex = -1
			maxIndex = -1
			continue
		}

		if num == minK {
			minIndex = right
		}

		if num == maxK {
			maxIndex = right
		}

		res += int64(max(0, min(minIndex, maxIndex)-left+1))
	}

	return res
}

func countSubarraysTest() {
	fmt.Println(countSubarrays([]int{1, 3, 5, 2, 7, 5}, 1, 5)) // 2
	fmt.Println(countSubarrays([]int{1, 1, 1, 1}, 1, 1))       // 10
}

// func main() {
// 	countSubarraysTest()
// }
