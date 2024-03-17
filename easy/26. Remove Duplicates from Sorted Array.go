package main

import "fmt"

func removeDuplicates2(nums []int) int {
	hashMap := make(map[int]bool)
	orderedNums := []int{}

	for _, x := range nums {
		if _, ok := hashMap[x]; !ok {
			hashMap[x] = true
			orderedNums = append(orderedNums, x)
		}
	}

	for i := 0; i < len(nums); i++ {
		if i < len(orderedNums) {
			nums[i] = orderedNums[i]
		} else {
			nums[i] = 0
		}
	}

	return len(orderedNums)
}

func removeDuplicates(nums []int) int {
	cnt := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] && cnt < len(nums) {
			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums, -101)
			cnt++
			// fmt.Println(nums, i, cnt)
			i--
		} else if nums[i-1] > nums[i] {
			return len(nums) - cnt
		}
	}

	return len(nums)
}

func RemoveDuplicatesTest() {
	fmt.Println(removeDuplicates2([]int{1, 1, 2}))
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

// func main() {
// 	RemoveDuplicatesTest()
// }
