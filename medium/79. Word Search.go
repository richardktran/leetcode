package main

import "fmt"

var dirs = [][]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}

func dfs(board [][]byte, word string, i, j, iWord int, seen []bool) bool {
	if iWord == len(word) {
		return true
	}

	n := len(board)
	m := len(board[0])

	for _, dir := range dirs {
		ii := i + dir[0]
		jj := j + dir[1]
		val := ii*m + jj

		if ii >= 0 && ii < n && jj >= 0 && jj < m && !seen[val] && board[ii][jj] == word[iWord] {
			seen[val] = true
			if dfs(board, word, ii, jj, iWord+1, seen) {
				return true
			}
			seen[val] = false
		}
	}

	return false
}

func exist(board [][]byte, word string) bool {
	n := len(board)
	m := len(board[0])

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == word[0] {
				seen := make([]bool, n*m)
				seen[i*m+j] = true
				if dfs(board, word, i, j, 1, seen) {
					return true
				}
			}
		}
	}

	return false
}

func existTest() {
	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED")) // true
	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "SEE"))    // true
	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCB"))   // false
}

// func main() {
// 	existTest()
// }
