package main

import (
	"fmt"
)

func wonderfulSubstrings(word string) int64 {
	n := len(word)
	cnt := make([]int64, 1024)

	cnt[0] = 1
	res := int64(0)

	mask := int64(0)
	for i := 0; i < n; i++ {
		mask ^= 1 << (word[i] - 'a')
		res += cnt[mask]

		for j := 0; j < 10; j++ {
			res += cnt[mask^(1<<j)]
		}

		cnt[mask]++
	}

	return res
}

func wonderfulSubstringsTest() {
	fmt.Println(wonderfulSubstrings("aba"))  // 4
	fmt.Println(wonderfulSubstrings("aabb")) // 9
	fmt.Println(wonderfulSubstrings("he"))   // 2
}

// func main() {
// 	wonderfulSubstringsTest()
// }
