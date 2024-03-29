package main

func countSubarrays(nums []int, k int) int64 {
	n := len(nums)

	left, right := 0, 0
	cnt := 0
	maxNum := 0
	var res int64 = 0

	for _, num := range nums {
		maxNum = max(maxNum, num)
	}

	for right < n {
		if nums[right] == maxNum {
			cnt++
		}

		for left <= right && cnt >= k {
			res += int64(n - right) // left to right is the minimum subarray that satisfies the condition, so all subarrays that end at right also satisfy the condition

			if nums[left] == maxNum {
				cnt--
			}

			left++
		}

		right++
	}

	return res
}

func countSubarraysTest() {
	println(countSubarrays([]int{1, 3, 2, 3, 3}, 2))
	println(countSubarrays([]int{1, 4, 2, 1}, 3))
	println(countSubarrays([]int{37, 20, 38, 66, 34, 38, 9, 41, 1, 14, 25, 63, 8, 12, 66, 66, 60, 12, 35, 27, 16, 38, 12, 66, 38, 36, 59, 54, 66, 54, 66, 48, 59, 66, 34, 11, 50, 66, 42, 51, 53, 66, 31, 24, 66, 44, 66, 1, 66, 66, 29, 54}, 5))
}

// func main() {
// 	countSubarraysTest()
// }
