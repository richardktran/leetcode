package main

import (
	"fmt"
	"sort"
)

func findRelativeRanks(score []int) []string {
	scoreOrigin := make([]int, len(score))
	copy(scoreOrigin, score)

	sort.Ints(score)
	scoreMap := make(map[int]string)

	for i := 0; i < len(score); i++ {
		switch i {
		case len(score) - 1:
			scoreMap[score[i]] = "Gold Medal"
		case len(score) - 2:
			scoreMap[score[i]] = "Silver Medal"
		case len(score) - 3:
			scoreMap[score[i]] = "Bronze Medal"
		default:
			scoreMap[score[i]] = fmt.Sprintf("%d", len(score)-i)
		}
	}

	var result []string
	for _, s := range scoreOrigin {
		result = append(result, scoreMap[s])
	}

	return result
}

func findRelativeRanksTest() {
	fmt.Println(findRelativeRanks([]int{5, 4, 3, 2, 1}))  // ["Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"]
	fmt.Println(findRelativeRanks([]int{10, 3, 8, 9, 4})) // ["Gold Medal", "5", "Bronze Medal", "Silver Medal", "4"]
}

// func main() {
// 	findRelativeRanksTest()
// }
