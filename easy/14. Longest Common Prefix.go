package main

import (
	"sort"
)

// brute force
func longestCommonPrefix1(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	minLenth := len(strs[0])

	for s := range strs {
		minLenth = min(minLenth, len(strs[s]))
	}

	for i := 0; i < minLenth; i++ {
		for _, s := range strs {
			if s[:i+1] != strs[0][:i+1] {
				return s[0:i]
			}
		}
	}

	return strs[0][0:minLenth]
}

func longestCommonPrefix(strs []string) string {
	sort.Strings(strs)

	for i := 0; i < len(strs[0]); i++ {
		if strs[0][i] != strs[len(strs)-1][i] {
			return strs[0][:i]
		}
	}

	return strs[0]
}

func LongestCommonPrefixTest() {
	println(longestCommonPrefix1([]string{"flower", "flow", "flight"}))
	println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}

// func main() {
// 	LongestCommonPrefixTest()
// }
