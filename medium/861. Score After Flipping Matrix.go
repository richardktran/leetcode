package main

import "fmt"

// TODO: Need to understand the logic
func matrixScore(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	for i := 0; i < n; i++ {
		if grid[i][0] == 0 {
			for j := 0; j < m; j++ {
				grid[i][j] = 1 - grid[i][j]
			}
		}
	}

	for j := 1; j < m; j++ {
		count := 0
		for i := 0; i < n; i++ {
			if grid[i][j] == 0 {
				count++
			}
		}
		if count > n/2 {
			for i := 0; i < n; i++ {
				grid[i][j] = 1 - grid[i][j]
			}
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			res += grid[i][j] << (m - j - 1)
		}
	}
	return res
}

func matrixScoreTest() {
	fmt.Println(matrixScore([][]int{{0, 0, 1, 1}, {1, 0, 1, 0}, {1, 1, 0, 0}})) // 39
	fmt.Println(matrixScore([][]int{{0}}))                                      // 1
}

// func main() {
// 	matrixScoreTest()
// }
