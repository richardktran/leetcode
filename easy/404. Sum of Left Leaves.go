package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := 0
	stack := [][]any{{root, false}}

	for len(stack) > 0 {
		node := stack[len(stack)-1][0].(*TreeNode)
		isLeft := stack[len(stack)-1][1].(bool)

		stack = stack[:len(stack)-1]

		if node.Left == nil && node.Right == nil && isLeft {
			sum += node.Val
		} else {
			if node.Left != nil {
				stack = append(stack, []any{node.Left, true})
			}
			if node.Right != nil {
				stack = append(stack, []any{node.Right, false})
			}
		}
	}

	return sum
}

func sumOfLeftLeavesTest() {
	// [3,9,20,null,null,15,7]
	// 9 + 15 = 24
	// 3
	// / \
	// 9  20
	// /  \
	// 15   7
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println(sumOfLeftLeaves(root)) // 24
}

// func main() {
// 	sumOfLeftLeavesTest()
// }
