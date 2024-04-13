package main

func largestRectangleArea(heights []int) int {
	n := len(heights)

	stack := []int{}
	leftMin := make([]int, n)
	rightMin := make([]int, n)

	for i := 0; i < n; i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			leftMin[i] = 0
		} else {
			leftMin[i] = stack[len(stack)-1] + 1
		}

		stack = append(stack, i)
	}
	stack = []int{}

	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			rightMin[i] = n - 1
		} else {
			rightMin[i] = stack[len(stack)-1] - 1
		}

		stack = append(stack, i)
	}

	area := 0

	for i := 0; i < n; i++ {
		area = max(area, (rightMin[i]-leftMin[i]+1)*heights[i])
	}

	return area
}

func largestRectangleAreaTest() {
	println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3})) // 10
	println(largestRectangleArea([]int{2, 4}))             // 4
}

// func main() {
// 	largestRectangleAreaTest()
// }
