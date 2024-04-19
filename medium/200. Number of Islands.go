package main

import "fmt"

var di = []int{0, -1, 0, 1}
var dj = []int{-1, 0, 1, 0}

type pair struct {
	first, second int
}

func numIsLandsBFS(grid *[][]byte, visited *[][]bool, i, j, n, m int) {
	g := *grid

	var queue []pair

	(*visited)[i][j] = true
	queue = append(queue, pair{i, j})

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		for k := 0; k < 4; k++ {
			ii := u.first + di[k]
			jj := u.second + dj[k]

			if ii < 0 || ii >= n || jj < 0 || jj >= m {
				continue
			}

			if (*visited)[ii][jj] == false && g[ii][jj] == '1' {
				(*visited)[ii][jj] = true
				queue = append(queue, pair{ii, jj})
			}
		}
	}
}

func numIslands(grid [][]byte) int {
	n, m := len(grid), len(grid[0])

	visited := make([][]bool, n)
	for v := range visited {
		visited[v] = make([]bool, m)
	}

	cnt := 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if visited[i][j] == false && grid[i][j] == '1' {
				cnt++
				numIsLandsBFS(&grid, &visited, i, j, n, m)
			}
		}
	}

	return cnt
}

func numIslandsTest() {
	grid := [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}
	result := numIslands(grid)
	fmt.Println(result) // 1

	fmt.Println("=====================================")
	grid = [][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}}
	result = numIslands(grid)
	fmt.Println(result) // 3
}

// func main() {
// 	numIslandsTest()
// }
