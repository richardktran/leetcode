package main

import (
	"fmt"
)

func minFallingPathSum(grid [][]int) int {
	n, m := len(grid), len(grid[0])

	var dp = make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}

	if n == 1 {
		return grid[0][0]
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			dp[i][j] = 20000
		}
	}

	for j := 0; j < m; j++ {
		dp[0][j] = grid[0][j]
	}

	for i := 1; i < n; i++ {
		for j := 0; j < m; j++ {
			for k := 0; k < m; k++ {
				if k == j {
					continue
				}

				dp[i][j] = min(dp[i][j], dp[i-1][k]+grid[i][j])
			}
		}
	}

	result := 20000

	for j := 0; j < m; j++ {
		result = min(result, dp[n-1][j])
	}

	return result
}

func minFallingPathSumTest() {
	fmt.Println(minFallingPathSum([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))                                                                                   // 13
	fmt.Println(minFallingPathSum([][]int{{7}}))                                                                                                               // 7
	fmt.Println(minFallingPathSum([][]int{{-73, 61, 43, -48, -36}, {3, 30, 27, 57, 10}, {96, -76, 84, 59, -15}, {5, -49, 76, 31, -7}, {97, 91, 61, -46, 67}})) // -192
}

// func main() {
// 	minFallingPathSumTest()
// }
