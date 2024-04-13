package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func trap(height []int) int {
	n := len(height)
	i, j := 0, n-1

	lMax, rMax := height[i], height[j]
	result := 0

	for i < j {
		if lMax <= rMax {
			result += lMax - height[i]
			i++
			lMax = max(lMax, height[i])
		} else {
			result += rMax - height[j]
			j--
			rMax = max(rMax, height[j])
		}
	}

	return result
}

func trapTest() {
	println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})) // 6
	println(trap([]int{4, 2, 0, 3, 2, 5}))                   // 9
}

// func main() {
// 	trapTest()
// }
