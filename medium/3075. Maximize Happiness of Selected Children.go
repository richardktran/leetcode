package main

import (
	"fmt"
	"sort"
)

func maximumHappinessSum(happiness []int, k int) int64 {
	sort.Ints(happiness)

	sum := 0
	n := len(happiness)

	for i := 0; i < k; i++ {
		if n-1-i < 0 {
			break
		}
		sum += max(0, happiness[n-1-i]-i)
	}

	return int64(sum)
}

func maximumHappinessSumTest() {
	fmt.Println(maximumHappinessSum([]int{1, 2, 3}, 2))    // 4
	fmt.Println(maximumHappinessSum([]int{1, 1, 1, 1}, 2)) // 1
	fmt.Println(maximumHappinessSum([]int{2, 3, 4, 5}, 1)) // 5
}

// func main() {
// 	maximumHappinessSumTest()
// }
