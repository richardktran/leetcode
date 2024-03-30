package main

// Sliding window advanced
func subarraysWithAtLeastKDistinct(nums []int, k int) int {
	n := len(nums)
	hashSet := map[int]int{}

	left, right := 0, 0
	res := 0
	numDiff := 0

	for right < n {
		if _, ok := hashSet[nums[right]]; !ok || hashSet[nums[right]] == 0 {
			numDiff++
			hashSet[nums[right]] = 1
		} else {
			hashSet[nums[right]]++
		}

		for left <= right && numDiff > k {
			hashSet[nums[left]]--
			if hashSet[nums[left]] == 0 {
				numDiff--
				delete(hashSet, nums[left])
			}
			left++
		}

		res += right - left + 1

		right++
	}
	return res
}

func subarraysWithKDistinct(nums []int, k int) int {
	return subarraysWithAtLeastKDistinct(nums, k) - subarraysWithAtLeastKDistinct(nums, k-1)
}

func subarraysWithKDistinctTest() {
	println(subarraysWithKDistinct([]int{1, 2, 1, 2, 3}, 2)) // 7
	println(subarraysWithKDistinct([]int{1, 2, 1, 3, 4}, 3)) // 3
}

// func main() {
// 	subarraysWithKDistinctTest()
// }
