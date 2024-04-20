package main

import "fmt"

type findFarmlandPair struct {
	first, second int
}

func findFarmlandDFS(land [][]int, groups *[][]int, i, j, cnt int) (int, int, int, int) {
	var di = []int{0, -1, 0, 1}
	var dj = []int{-1, 0, 1, 0}
	g := *groups
	n, m := len(land), len(land[0])

	var stack []findFarmlandPair

	stack = append(stack, findFarmlandPair{i, j})

	r1, c1, r2, c2 := i, j, i, j

	for len(stack) > 0 {
		u := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for k := 0; k < 4; k++ {
			ii := u.first + di[k]
			jj := u.second + dj[k]

			if ii < 0 || ii >= n || jj < 0 || jj >= m {
				continue
			}

			if g[ii][jj] == 0 && land[ii][jj] == 1 {
				g[ii][jj] = cnt
				r1 = min(r1, ii)
				c1 = min(c1, jj)
				r2 = max(r2, ii)
				c2 = max(c2, jj)
				stack = append(stack, findFarmlandPair{ii, jj})
			}
		}
	}
	return r1, c1, r2, c2
}

func findFarmland(land [][]int) [][]int {
	n, m := len(land), len(land[0])

	groups := make([][]int, n)
	for c := range groups {
		groups[c] = make([]int, m)
	}

	cnt := 0
	result := [][]int{}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if groups[i][j] == 0 && land[i][j] == 1 {
				cnt++
				r1, c1, r2, c2 := findFarmlandDFS(land, &groups, i, j, cnt)
				result = append(result, []int{r1, c1, r2, c2})
			}
		}
	}

	return result
}

func findFarmLandTest() {
	land := [][]int{{1, 0, 0}, {0, 1, 1}, {0, 1, 1}}
	fmt.Println(findFarmland(land)) // [[0, 0, 0, 0], [1, 1, 2, 2]]

	land = [][]int{{1, 1}, {1, 1}}
	fmt.Println(findFarmland(land)) // [[0, 0, 1, 1]]

	land = [][]int{{0}}
	fmt.Println(findFarmland(land)) // []
}

// func main() {
// 	findFarmLandTest()
// }
