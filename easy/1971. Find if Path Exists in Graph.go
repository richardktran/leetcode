package main

import "fmt"

// DFS with Edge List
func validPath(n int, edges [][]int, source int, destination int) bool {
	var graph = make([][]int, n)

	for _, edge := range edges {
		u := edge[0]
		v := edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	var visited = make([]bool, n)
	var stack = []int{source}

	for len(stack) > 0 {
		u := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		visited[u] = true

		if u == destination {
			return true
		}

		for _, v := range graph[u] {
			if !visited[v] {
				stack = append(stack, v)
			}
		}

	}

	return false
}

func validPathTest() {
	fmt.Println(validPath(3, [][]int{{0, 1}, {1, 2}, {2, 0}}, 0, 2))                 // true
	fmt.Println(validPath(6, [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}, 0, 5)) // false
}

// func main() {
// 	validPathTest()
// }
