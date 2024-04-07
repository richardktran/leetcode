package main

import "fmt"

func checkValidString(s string) bool {
	var stack []int
	var asterisk []int

	for i, c := range s {
		if c == '(' {
			stack = append(stack, i)
		} else if c == '*' {
			asterisk = append(asterisk, i)
		} else {
			if len(stack) == 0 && len(asterisk) == 0 {
				return false
			}

			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			} else {
				asterisk = asterisk[:len(asterisk)-1]
			}
		}
	}

	for len(stack) > 0 {
		if len(asterisk) == 0 || stack[len(stack)-1] > asterisk[len(asterisk)-1] {
			return false
		}

		stack = stack[:len(stack)-1]
		asterisk = asterisk[:len(asterisk)-1]
	}

	return true
}

func checkValidStringTest() {
	fmt.Println(checkValidString("()"))    // true
	fmt.Println(checkValidString("(*)"))   // true
	fmt.Println(checkValidString("(*))"))  // true
	fmt.Println(checkValidString("(*))"))  // true
	fmt.Println(checkValidString("(*)))")) // false
}

// func main() {
// 	checkValidStringTest()
// }
