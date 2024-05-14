package main

import "fmt"

func getMaximumGold(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	var dfs func(int, int) int

	dfs = func(x, y int) int {
		if x < 0 || x >= n || y < 0 || y >= m || grid[x][y] == 0 {
			return 0
		}

		tmp := grid[x][y]
		grid[x][y] = 0
		up := dfs(x-1, y)
		down := dfs(x+1, y)
		left := dfs(x, y-1)
		right := dfs(x, y+1)
		grid[x][y] = tmp
		return tmp + max(up, down, left, right)
	}

	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] != 0 {
				res = max(res, dfs(i, j))
			}
		}
	}

	return res
}

func getMaximumGoldTest() {
	fmt.Println(getMaximumGold([][]int{{0, 6, 0}, {5, 8, 7}, {0, 9, 0}}))                        // 24
	fmt.Println(getMaximumGold([][]int{{1, 0, 7}, {2, 0, 6}, {3, 4, 5}, {0, 3, 0}, {9, 0, 20}})) // 28
}

// func main() {
// 	getMaximumGoldTest()
// }
