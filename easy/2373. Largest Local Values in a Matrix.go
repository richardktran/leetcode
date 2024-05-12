package main

import "fmt"

func largestLocal(grid [][]int) [][]int {
	var dirs = [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

	var n = len(grid)

	var res = make([][]int, n-2)
	for i := range res {
		res[i] = make([]int, n-2)
	}

	for i := 1; i < n-1; i++ {
		for j := 1; j < n-1; j++ {
			localMax := grid[i][j]

			for _, dir := range dirs {
				x, y := i+dir[0], j+dir[1]

				if grid[x][y] > localMax {
					localMax = grid[x][y]
				}
			}

			res[i-1][j-1] = localMax
		}
	}

	return res
}

func largestLocalTest() {
	fmt.Println(largestLocal([][]int{
		{9, 9, 8, 1},
		{5, 6, 2, 6},
		{8, 2, 6, 4},
		{6, 2, 2, 2},
	})) // [[9, 9], [8, 6]]
}

// func main() {
// 	largestLocalTest()
// }
