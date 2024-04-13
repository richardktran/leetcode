package main

func largestRectangleAreaHistogram(heights []int) int {
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

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	n, m := len(matrix), len(matrix[0])

	heights := make([]int, m)
	area := 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == '1' {
				heights[j]++
			} else {
				heights[j] = 0
			}
		}
		area = max(area, largestRectangleAreaHistogram(heights))
	}

	return area
}

func maximalRectangleTest() {
	println(maximalRectangle([][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	})) // 6
	println(maximalRectangle([][]byte{
		{'0'},
	})) // 0
	println(maximalRectangle([][]byte{
		{'1'},
	})) // 1
}

// func main() {
// 	maximalRectangleTest()
// }
