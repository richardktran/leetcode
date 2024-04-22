package main

import (
	"fmt"
)

type openLockNode struct {
	value string
	depth int
}

func plusOne(s string, j int) string {
	var bytes = []byte(s)
	if bytes[j] == '9' {
		bytes[j] = '0'
	} else {
		bytes[j]++
	}
	return string(bytes)
}

func minusOne(s string, j int) string {
	var bytes = []byte(s)
	if bytes[j] == '0' {
		bytes[j] = '9'
	} else {
		bytes[j]--
	}
	return string(bytes)
}

func openLock(deadends []string, target string) int {
	sInit := "0000"
	var queue []openLockNode
	var visited = make(map[string]bool)

	for _, deadend := range deadends {
		if deadend == sInit || deadend == target {
			return -1
		}
		visited[deadend] = true
	}

	queue = append(queue, openLockNode{sInit, 0})

	for len(queue) > 0 {
		u := queue[0].value
		depth := queue[0].depth
		queue = queue[1:]

		visited[u] = true

		if u == target {
			return depth
		}

		for i := 0; i < 4; i++ {
			x := plusOne(u, i)
			y := minusOne(u, i)

			if !visited[x] {
				queue = append(queue, openLockNode{x, depth + 1})
				visited[x] = true
			}

			if !visited[y] {
				queue = append(queue, openLockNode{y, depth + 1})
				visited[y] = true
			}
		}
	}

	return -1
}

func openLockTest() {
	fmt.Println(openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202"))                         // 6
	fmt.Println(openLock([]string{"8888"}, "0009"))                                                         // 1
	fmt.Println(openLock([]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888")) // -1
	fmt.Println(openLock([]string{"0000"}, "8888"))                                                         // -1
}

// func main() {
// 	openLockTest()
// }
