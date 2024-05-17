package main

import "fmt"

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left = removeLeafNodes(root.Left, target)
	root.Right = removeLeafNodes(root.Right, target)

	if root.Left == nil && root.Right == nil && root.Val == target {
		return nil
	}

	return root
}

func removeLeafNodesTest() {
	root := &TreeNode{1, &TreeNode{2, &TreeNode{2, nil, nil}, nil}, &TreeNode{3, &TreeNode{2, nil, nil}, &TreeNode{4, nil, nil}}}
	target := 2

	result := removeLeafNodes(root, target)

	// print result with lnr order
	printLNR(result)
}

func printLNR(root *TreeNode) {
	if root == nil {
		return
	}

	printLNR(root.Left)
	fmt.Println(root.Val)
	printLNR(root.Right)
}

// func main() {
// 	removeLeafNodesTest()
// }
