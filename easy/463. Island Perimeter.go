package main

import "fmt"

var di = []int{0, -1, 0, 1}
var dj = []int{-1, 0, 1, 0}

func islandPerimeter(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	result := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 0 {
				continue
			}

			cnt := 0

			for k := 0; k < 4; k++ {
				if i+di[k] < 0 || i+di[k] >= n || j+dj[k] < 0 || j+dj[k] >= m {
					continue
				}
				if grid[i+di[k]][j+dj[k]] == 1 {
					cnt++
				}
			}

			result += (4 - cnt)
		}
	}

	return result
}

func islandPerimeterTest() {
	grid := [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}
	result := islandPerimeter(grid)
	fmt.Println(result) // 16

	fmt.Println("=====================================")
	grid = [][]int{{1}}
	result = islandPerimeter(grid)
	fmt.Println(result) // 4

	fmt.Println("=====================================")
	grid = [][]int{{1, 0}}
	result = islandPerimeter(grid)
	fmt.Println(result) // 4
}

// func main() {
// 	islandPerimeterTest()
// }
