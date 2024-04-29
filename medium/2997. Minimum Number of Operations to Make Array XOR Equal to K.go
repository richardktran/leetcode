package main

import "fmt"

// Bitwise XOR
func countOne(num int) int {
	cnt := 0

	for num > 0 {
		num = num & (num - 1)
		cnt++
	}

	return cnt
}

func minOperations(nums []int, k int) int {
	original := 0

	for _, num := range nums {
		original ^= num
	}

	original ^= k

	return countOne(original)
}

func minOperationsTest() {
	fmt.Println(minOperations([]int{2, 1, 3, 4}, 1)) // 2
	fmt.Println(minOperations([]int{2, 0, 2, 0}, 0)) // 0
}

// func main() {
// 	minOperationsTest()
// }
