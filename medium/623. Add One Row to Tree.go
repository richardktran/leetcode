package main

import "fmt"

func addRow(node *TreeNode, val int, depth int, currentDepth int) {
	if currentDepth == depth-1 {
		left := node.Left
		right := node.Right

		node.Left = &TreeNode{val, left, nil}
		node.Right = &TreeNode{val, nil, right}
	}

	if node.Left != nil {
		addRow(node.Left, val, depth, currentDepth+1)
	}

	if node.Right != nil {
		addRow(node.Right, val, depth, currentDepth+1)
	}
}

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		return &TreeNode{val, root, nil}
	}

	addRow(root, val, depth, 1)

	return root
}

// Draw tree as graph tree structure for better visualization of tree structure with depth level
func drawTreeNode(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Println(node.Val)
	if node.Left != nil {
		fmt.Printf("L:%d", node.Left.Val)
	}
	if node.Right != nil {
		fmt.Printf(" R:%d", node.Right.Val)
	}
	fmt.Println()
	drawTreeNode(node.Left)
	drawTreeNode(node.Right)
}

func addOneRowTest() {
	root := &TreeNode{4, &TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{1, nil, nil}}, &TreeNode{6, &TreeNode{5, nil, nil}, nil}}
	val := 1
	depth := 2
	result := addOneRow(root, val, depth)
	drawTreeNode(result)

	fmt.Println("=====================================")

	root = &TreeNode{4, &TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{1, nil, nil}}, nil}
	val = 1
	depth = 3
	result = addOneRow(root, val, depth)
	drawTreeNode(result)
}

// func main() {
// 	addOneRowTest()
// }
