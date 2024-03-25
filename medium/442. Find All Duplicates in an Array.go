package main

import "fmt"

func findDuplicates(nums []int) []int {
	hashSet := map[int]bool{}
	result := []int{}

	for _, num := range nums {
		if hashSet[num] == true {
			result = append(result, num)
			continue
		}
		hashSet[num] = true
	}

	return result
}

func FindDuplicatesTest() {
	fmt.Println(findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	fmt.Println(findDuplicates([]int{1, 1, 2}))
	fmt.Println(findDuplicates([]int{1}))
}

// func main() {
// 	FindDuplicatesTest()
// }
