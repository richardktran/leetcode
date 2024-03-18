package main

import (
	"fmt"
	"sort"
)

func findMinArrowShots(points [][]int) int {
	sort.SliceStable(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	cnt := 1

	xEnd := points[0][1]

	for i := 1; i < len(points); i++ {
		if points[i][0] > xEnd {
			cnt++
			xEnd = points[i][1]
		}
	}

	return cnt
}

func FindMinArrowShotsTest() {
	fmt.Println(findMinArrowShots([][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}))
	fmt.Println(findMinArrowShots([][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}, {11, 18}}))
	fmt.Println(findMinArrowShots([][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}))
	fmt.Println(findMinArrowShots([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}))
}

// func main() {
// 	FindMinArrowShotsTest()
// }
