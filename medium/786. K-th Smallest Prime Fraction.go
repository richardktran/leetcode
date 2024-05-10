package main

import "fmt"

// TODO: Need to understand the solution
func kthSmallestPrimeFraction(arr []int, k int) []int {
	n := len(arr)
	l, r := 0.0, 1.0
	for l < r {
		mid := (l + r) / 2
		count, maxP, up, bot := 0, 0.0, 0, 0
		j := 1
		for i := 0; i < n; i++ {
			for j < n && float64(arr[i]) >= mid*float64(arr[j]) {
				j++
			}
			count += n - j
			if j < n && maxP < float64(arr[i])/float64(arr[j]) {
				maxP = float64(arr[i]) / float64(arr[j])
				up, bot = arr[i], arr[j]
			}
		}
		if count == k {
			return []int{up, bot}
		} else if count < k {
			l = mid
		} else {
			r = mid
		}
	}
	return []int{}
}

func kthSmallestPrimeFractionTest() {
	fmt.Println(kthSmallestPrimeFraction([]int{1, 2, 3, 5}, 3)) // [2, 5]
	fmt.Println(kthSmallestPrimeFraction([]int{1, 7}, 1))       // [1, 7]
}

// func main() {
// 	kthSmallestPrimeFractionTest()
// }
