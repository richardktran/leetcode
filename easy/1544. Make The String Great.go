package main

import (
	"fmt"
)

func isBad(a, b byte) bool {
	return a != b && (a == b-32 || a == b+32)
}

func makeGood(s string) string {
	if len(s) < 2 {
		return s
	}
	stack := []byte{s[0]}

	for i := 1; i < len(s); i++ {
		if len(stack) > 0 && isBad(stack[len(stack)-1], s[i]) {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}

	return string(stack)
}

func makeGoodTest() {
	fmt.Println(makeGood("leEeetcode")) // "leetcode"
	fmt.Println(makeGood("abBAcC"))     // ""
	fmt.Println(makeGood("s"))          // "s"
	fmt.Println(makeGood("Pp"))         // ""
	fmt.Println(makeGood("aA"))         // ""
	fmt.Println(makeGood("ab"))         // "ab"
	fmt.Println(makeGood("Ab"))         // "Ab"
	fmt.Println(makeGood("aB"))         // "aB"
	fmt.Println(makeGood("AB"))         // "AB"
	fmt.Println(makeGood("a"))          // "a"
}

// func main() {
// 	makeGoodTest()
// }
