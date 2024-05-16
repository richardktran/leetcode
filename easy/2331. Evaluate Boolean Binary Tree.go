package main

import "fmt"

func evaluateTree(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return root.Val == 1
	}

	if root.Val == 2 {
		return evaluateTree(root.Left) || evaluateTree(root.Right)
	}

	return evaluateTree(root.Left) && evaluateTree(root.Right)
}

func evaluateTreeTest() {
	root := &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 1}}}
	fmt.Println(evaluateTree(root)) // true

	root = &TreeNode{Val: 0, Left: nil, Right: nil}
	fmt.Println(evaluateTree(root)) // false
}

// func main() {
// 	evaluateTreeTest()
// }
