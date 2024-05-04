package main

import (
	"fmt"
	"sort"
)

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)

	left, right := 0, len(people)-1

	res := 0

	for left <= right {
		if people[left]+people[right] <= limit {
			left++
		}
		right--
		res++
	}

	return res
}

func NumRescueBoatsTest() {
	fmt.Println(numRescueBoats([]int{1, 2}, 3))       // 1
	fmt.Println(numRescueBoats([]int{3, 2, 2, 1}, 3)) // 3
	fmt.Println(numRescueBoats([]int{3, 5, 3, 4}, 5)) // 4
}

// func main() {
// 	NumRescueBoatsTest()
// }
