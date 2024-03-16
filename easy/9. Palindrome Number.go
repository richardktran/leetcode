package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	if x < 10 {
		return true
	}

	n := 0
	for x >= n {
		n = n*10 + x%10
		if x == n || x/10 == n {
			return true
		}
		x = x / 10
	}

	return false
}

func IsPalindromeTest() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
	fmt.Println(isPalindrome(-101))
	fmt.Println(isPalindrome(0))
}

// func main() {
// 	IsPalindromeTest()
// }
