package main

import "fmt"

// nums = [3, 4, -1, 1]
// Add 0 to nums
// Set true for each element in nums
// p1 = {0: true, 3: true, 4: true, -1: true, 1: true}
// Check if v+1 exists in p1, if yes set true in p2
// p2 = {0: true, 3: true, 4: false, -1: true, 1: false}
// Find the minimum value in nums that does not exist in p2 (1,4)
// Result is the minimum value + 1 => 2
func firstMissingPositive(nums []int) int {
	p1 := map[int]bool{0: true}
	p2 := map[int]bool{}

	nums = append(nums, 0)

	for _, v := range nums {
		p1[v] = true
	}

	for _, v := range nums {
		if _, ok := p1[v+1]; v+1 >= 0 && ok {
			p2[v] = true
		}
	}

	minn := 2147483648

	for _, v := range nums {
		if _, ok := p2[v]; ok {
			continue
		}
		if v+1 <= 0 {
			continue
		}

		if v < minn {
			minn = v
		}
	}

	return minn + 1

}

func FirstMissingPositiveTest() {
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))         // 3
	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))     // 2
	fmt.Println(firstMissingPositive([]int{7, 8, 9, 11, 12})) // 1
	fmt.Println(firstMissingPositive([]int{1, 2, 3, 4, 5}))   // 6
	fmt.Println(firstMissingPositive([]int{1, 1, 1, 1, 1}))   // 2
}

// func main() {
// 	FirstMissingPositiveTest()
// }
