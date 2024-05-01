package main

import (
	"fmt"
	"strings"
)

func reversePrefix(word string, ch byte) string {
	var reverse func(s string) string

	reverse = func(s string) string {
		var res string

		for _, v := range s {
			res = string(v) + res
		}

		return res
	}

	index := strings.IndexRune(word, rune(ch))
	if index == -1 {
		return word
	}

	return reverse(word[:index+1]) + word[index+1:]

}

func reversePrefixTest() {
	fmt.Println(reversePrefix("abcdefd", 'd')) // "dcbaefd"
	fmt.Println(reversePrefix("xyxzxe", 'z'))  // "zxyxxe"
	fmt.Println(reversePrefix("abcd", 'z'))    // "abcd"
}

// func main() {
// 	reversePrefixTest()
// }
