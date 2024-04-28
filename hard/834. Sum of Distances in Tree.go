package main

import "fmt"

func sumOfDistancesInTree(n int, edges [][]int) []int {
	var (
		dfs1 func(node, parent int)
		dfs2 func(node, parent int)
	)
	var graph = make([][]int, n)
	var count = make([]int, n)
	var result = make([]int, n)

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	dfs1 = func(node, parant int) {
		for _, child := range graph[node] {
			if child == parant {
				continue
			}
			dfs1(child, node)
			count[node] += count[child]
			result[node] += result[child] + count[child]
		}
		count[node]++
	}

	dfs2 = func(node, parent int) {
		for _, child := range graph[node] {
			if child == parent {
				continue
			}
			result[child] = result[node] - count[child] + n - count[child]
			dfs2(child, node)
		}
	}

	dfs1(0, -1)
	dfs2(0, -1)

	return result
}

func sumOfDistancesInTreeTest() {
	fmt.Println(sumOfDistancesInTree(6, [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}})) // [8,12,6,10,10,10]
	fmt.Println(sumOfDistancesInTree(1, [][]int{}))                                       // [0]
	fmt.Println(sumOfDistancesInTree(2, [][]int{{1, 0}}))                                 // [1,1]
}

// func main() {
// 	sumOfDistancesInTreeTest()
// }
