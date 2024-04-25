package main

import "fmt"

func toInt(c byte) int {
	return int(c - 'a')
}

func longestIdealString(s string, k int) int {
	n := len(s)
	var dp = make([]int, 26)

	for i := 0; i < n; i++ {
		c := toInt(s[i])
		lower := max(0, c-k)
		upper := min(25, c+k)

		for j := lower; j <= upper; j++ {
			dp[c] = max(dp[c], dp[j])
		}

		dp[c]++
	}

	result := 0

	for i := 0; i < 26; i++ {
		result = max(result, dp[i])
	}

	return result
}

func longestIdealStringTest() {
	fmt.Println(longestIdealString("acfgbd", 2)) // 4
	fmt.Println(longestIdealString("abcd", 3))   // 4
}

// func main() {
// 	longestIdealStringTest()
// }
