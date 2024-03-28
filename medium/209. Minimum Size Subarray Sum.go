package main

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)

	left, right := 0, 0
	result := 1e9 + 1
	sum := 0

	for right < n {
		sum += nums[right]

		for left <= right && sum >= target {
			result = min(result, float64(right-left+1))
			sum -= nums[left]
			left++
		}

		right++
	}

	if result == 1e9+1 {
		return 0
	}

	return int(result)
}

func minSubArrayLenTest() {
	println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))        // 2
	println(minSubArrayLen(4, []int{1, 4, 4}))                 // 1
	println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1})) // 0
}

func main() {
	minSubArrayLenTest()
}
