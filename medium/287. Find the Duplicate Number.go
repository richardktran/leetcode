package main

import "fmt"

func findDuplicate(nums []int) int {
	hashMap := map[int]bool{}
	for _, s := range nums {
		if hashMap[s] == true {
			return s
		}

		hashMap[s] = true
	}

	return nums[0]
}

func FindDuplicateTest() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}))
	fmt.Println(findDuplicate([]int{3, 1, 3, 4, 2}))
	fmt.Println(findDuplicate([]int{1, 1}))
	fmt.Println(findDuplicate([]int{1, 1, 2}))
	fmt.Println(findDuplicate([]int{3, 3, 3, 3, 3}))
}

func main() {
	FindDuplicateTest()
}
