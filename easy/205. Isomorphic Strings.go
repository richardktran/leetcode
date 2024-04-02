package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	m1 := make(map[byte]byte)
	m2 := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		if _, ok := m2[t[i]]; !ok {
			m2[t[i]] = s[i]
		} else {
			if m2[t[i]] != s[i] {
				return false
			}
		}

		if _, ok := m1[s[i]]; !ok {
			m1[s[i]] = t[i]
		} else {
			if m1[s[i]] != t[i] {
				return false
			}
		}
	}

	return true
}

func isIsomorphicTest() {
	fmt.Println(isIsomorphic("egg", "add"))     // true
	fmt.Println(isIsomorphic("foo", "bar"))     // false
	fmt.Println(isIsomorphic("paper", "title")) // true
	fmt.Println(isIsomorphic("badc", "baba"))   // false
}

// func main() {
// 	isIsomorphicTest()
// }
