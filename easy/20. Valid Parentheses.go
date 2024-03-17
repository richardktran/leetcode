package main

func isValid(s string) bool {
	pMap := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}

	stack := make([]string, 0)

	for _, c := range s {
		if len(stack) == 0 {
			stack = append(stack, string(c))
		} else {
			if pMap[string(c)] == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, string(c))
			}
		}
	}

	return len(stack) == 0
}

func ValidParenthesesTest() {
	println(isValid("()"))
	println(isValid("()[]{}"))
	println(isValid("(]"))
	println(isValid("([)]"))
	println(isValid("([)]"))
}

// func main() {
// 	ValidParenthesesTest()
// }
