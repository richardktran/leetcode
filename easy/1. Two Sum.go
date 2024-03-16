package main

import "fmt"

func twoSum(nums []int, target int) []int {
	a := make(map[int]int)

	for i, num := range nums {
		if j, ok := a[target-num]; ok {
			return []int{i, j}
		}

		a[num] = i
	}

	return []int{}
}

func TwoSumTest() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}

// func main() {
// 	TwoSumTest()
// }
