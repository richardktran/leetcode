package main

import "math"

// Sliding window
func maxSubarrayLength(nums []int, k int) int {
	n := len(nums)

	left, right := 0, 0
	result := 0

	hashSet := map[int]int{}

	for right < n {
		if _, ok := hashSet[nums[right]]; !ok {
			hashSet[nums[right]] = 1
		} else {
			hashSet[nums[right]]++
		}

		for left < right && hashSet[nums[right]] > k {
			hashSet[nums[left]] = int(math.Max(0.0, float64(hashSet[nums[left]])-1.0))
			left++
		}

		result = int(math.Max(float64(result), float64(right-left+1)))
		right++
	}

	return result
}

func maxSubarrayLengthTest() {
	println(maxSubarrayLength([]int{1, 2, 3, 1, 2, 3, 1, 2}, 2)) // 6
	println(maxSubarrayLength([]int{1, 2, 1, 2, 1, 2, 1, 2}, 1)) // 2
	println(maxSubarrayLength([]int{5, 5, 5, 5, 5, 5, 5}, 4))    // 4
}

// func main() {
// 	maxSubarrayLengthTest()
// }
