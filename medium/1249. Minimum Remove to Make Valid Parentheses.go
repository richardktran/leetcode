package main

import "fmt"

func minRemoveToMakeValid(s string) string {
	var stack []int
	var res []byte
	var removeIds []int

	for i, c := range s {
		if c == '(' {
			stack = append(stack, i)
		} else if c == ')' {
			if len(stack) == 0 {
				removeIds = append(removeIds, i)
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}

	removeIds = append(removeIds, stack...)

	for i, c := range s {
		if len(removeIds) > 0 && i == removeIds[0] {
			removeIds = removeIds[1:]
			continue
		}
		res = append(res, byte(c))
	}

	return string(res)
}

func minRemoveToMakeValidTest() {
	fmt.Println(minRemoveToMakeValid("lee(t(c)o)de)") == "lee(t(c)o)de") // "lee(t(c)o)de"
	fmt.Println(minRemoveToMakeValid("a)b(c)d") == "ab(c)d")             // "ab(c)d"
	fmt.Println(minRemoveToMakeValid("))((") == "")                      // ""
	fmt.Println(minRemoveToMakeValid("(a(b(c)d)") == "a(b(c)d)")         // "a(b(c)d)"
}

// func main() {
// 	minRemoveToMakeValidTest()
// }
