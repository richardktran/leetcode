package main

import "fmt"

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func findRotateSteps(ring string, key string) int {
	n, m := len(ring), len(key)
	var dp = make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = 200000
		}
	}

	for i := 0; i < n; i++ {
		if ring[i] == key[0] {
			dp[0][i] = min(i, n-i) + 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if ring[j] != key[i] {
				continue
			}

			for k := 0; k < n; k++ {
				if ring[k] != key[i-1] {
					continue
				}

				dp[i][j] = min(dp[i][j], dp[i-1][k]+min(abs(j-k), n-abs(j-k))+1)
			}
		}
	}

	result := 200000

	for j := 0; j < n; j++ {
		result = min(result, dp[m-1][j])
	}

	return result
}

func findRotateStepsTest() {
	fmt.Println(findRotateSteps("godding", "gd"))      // 4
	fmt.Println(findRotateSteps("godding", "godding")) // 13
}

// func main() {
// 	findRotateStepsTest()
// }
