package main

import "fmt"

func findMaxK(nums []int) int {

	negative := make(map[int]bool)

	for _, num := range nums {
		if num < 0 {
			negative[num] = true
		}
	}

	maxV := -1

	for _, num := range nums {
		if num > 0 && negative[-num] {
			maxV = max(maxV, num)
		}
	}

	return maxV
}

func findMaxKTest() {
	fmt.Println(findMaxK([]int{-1, 2, -3, 3}))         // 3
	fmt.Println(findMaxK([]int{-1, 10, 6, 7, -7, 1}))  // 7
	fmt.Println(findMaxK([]int{-10, 8, 6, 7, -2, -3})) // -1
}

// func main() {
// 	findMaxKTest()
// }
