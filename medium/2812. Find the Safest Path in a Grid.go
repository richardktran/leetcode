package main

import (
	"container/heap"
	"fmt"
	"math"
)

// TODO: Need to understand this code
type cell struct {
	x      int
	y      int
	factor int
}

type PriorityQueue []*cell

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].factor > pq[j].factor }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*cell)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[:n-1]
	return item
}

func maximumSafenessFactor(grid [][]int) int {
	n := len(grid)

	thiefs := make([][]int, n)
	for i := 0; i < len(thiefs); i++ {
		thiefs[i] = make([]int, n)
		for j := 0; j < len(thiefs[i]); j++ {
			thiefs[i][j] = -1
		}
	}

	queue := make([][2]int, 0)
	directions := [4][2]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				thiefs[i][j] = 0
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	for len(queue) > 0 {
		newQueue := make([][2]int, 0)
		for _, curr := range queue {
			row, col := curr[0], curr[1]
			for _, dir := range directions {
				x, y := row+dir[0], col+dir[1]
				if x >= 0 && x < n && y >= 0 && y < n && thiefs[x][y] == -1 {
					thiefs[x][y] = 1 + thiefs[row][col]
					newQueue = append(newQueue, [2]int{x, y})
				}
			}
		}
		queue = newQueue
	}

	visited := make([][]bool, n)
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, n)
	}

	visited[0][0] = true
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &cell{0, 0, thiefs[0][0]})

	res := math.MaxInt32
	for pq.Len() > 0 {
		top := heap.Pop(&pq).(*cell)
		row, col, factor := top.x, top.y, top.factor
		res = min(res, factor)
		if row == n-1 && col == n-1 {
			return res
		}
		for _, dir := range directions {
			x, y := row+dir[0], col+dir[1]
			if x >= 0 && x < n && y >= 0 && y < n && !visited[x][y] {
				visited[x][y] = true
				heap.Push(&pq, &cell{x, y, thiefs[x][y]})
			}
		}
	}

	return res
}

func maximumSafenessFactorTest() {

	fmt.Println(maximumSafenessFactor([][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 1}}))                        // 0
	fmt.Println(maximumSafenessFactor([][]int{{0, 0, 1}, {0, 0, 0}, {0, 0, 0}}))                        // 2
	fmt.Println(maximumSafenessFactor([][]int{{0, 0, 0, 1}, {0, 0, 0, 0}, {0, 0, 0, 0}, {1, 0, 0, 0}})) // 2
}

// func main() {
// 	maximumSafenessFactorTest()
// }
