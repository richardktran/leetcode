package main

import "fmt"

type findMinHeightTreesNode struct {
	val   int
	depth int
}

func MHT(root int, edgeList [][]int, n, minDepth int) int {
	var stack []findMinHeightTreesNode
	var visited = make([]bool, n)
	stack = append(stack, findMinHeightTreesNode{root, 0})

	maxDepth := 0

	for len(stack) > 0 {
		u := stack[len(stack)-1].val
		depth := stack[len(stack)-1].depth
		stack = stack[:len(stack)-1]

		if depth > maxDepth {
			maxDepth = depth
		}

		if depth > minDepth {
			return n + 1
		}

		visited[u] = true

		for _, v := range edgeList[u] {
			if !visited[v] {
				visited[v] = true
				stack = append(stack, findMinHeightTreesNode{v, depth + 1})
			}
		}
	}

	return maxDepth
}

func findMinHeightTrees(n int, edges [][]int) []int {
	var edgeList = make([][]int, n*2)

	for _, edge := range edges {
		edgeList[edge[0]] = append(edgeList[edge[0]], edge[1])
		edgeList[edge[1]] = append(edgeList[edge[1]], edge[0])
	}

	minDepth := n + 1
	var result []int

	for i := 0; i < n; i++ {
		mht := MHT(i, edgeList, n, minDepth)
		if mht < minDepth {
			minDepth = mht
			result = []int{i}
		} else if mht == minDepth {
			result = append(result, i)
		}
	}

	return result
}

func findMinHeightTreesTest() {
	fmt.Println(findMinHeightTrees(4, [][]int{{1, 0}, {1, 2}, {1, 3}}))                 // [1]
	fmt.Println(findMinHeightTrees(6, [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}})) // [3, 4]
}

func main() {
	findMinHeightTreesTest()
}
