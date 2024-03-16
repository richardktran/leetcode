package main

// brute force
func longestCommonPrefix(strs []string) string {
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

func LongestCommonPrefixTest() {
	println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}

// func main() {
// 	LongestCommonPrefixTest()
// }
