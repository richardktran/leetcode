package main

import "fmt"

// 10, 5, 2, 6
// left = 0 , right = 0 , mul = 10 , result = 1 (10)
// left = 0 , right = 1 , mul = 50 , result = 1+2 = 3 (10, 5) (5)
// left = 0 , right = 2 , mul = 100 >= 100, left++
// left = 1 , right = 2 , mul = 10 , result = 3+2 = 5 (5, 2) (2)
// left = 1 , right = 3 , mul = 60 , result = 5 + 3 = 8 (5, 2, 6) (2, 6) (6)
func numSubarrayProductLessThanK(nums []int, k int) int {
	i := 0
	mul := 1
	result := 0
	for j := 0; j < len(nums); j++ {
		mul *= nums[j]
		for mul >= k && i <= j {
			mul /= nums[i]
			i++
			fmt.Println("left =", i, ", right =", j, ", mul =", mul)
		}
		if i <= j {
			result += (j - i + 1)
			fmt.Println("left =", i, ", right =", j, ", mul =", mul, ", result =", result)
		}
	}

	return result
}

func NumSubarrayProductLessThanKTest() {
	fmt.Println(numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100)) // 8
	fmt.Println(numSubarrayProductLessThanK([]int{1, 2, 3}, 0))       // 0
}

// func main() {
// 	NumSubarrayProductLessThanKTest()
// }
