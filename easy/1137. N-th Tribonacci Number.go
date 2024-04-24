package main

import "fmt"

func tribonacci(n int) int {
	var f = make([]int, n+1)

	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}

	f[0] = 0
	f[1], f[2] = 1, 1

	for i := 3; i <= n; i++ {
		f[i] = f[i-1] + f[i-2] + f[i-3]
	}

	return f[n]
}

func tribonacciTest() {
	fmt.Println(tribonacci(4))  // 4
	fmt.Println(tribonacci(25)) // 1389537
	fmt.Println(tribonacci(37)) // 1389537
}

// func main() {
// 	tribonacciTest()
// }
