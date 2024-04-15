package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumLeaf(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	sum = sum*10 + root.Val

	if root.Left == nil && root.Right == nil {
		return sum
	}

	return sumLeaf(root.Left, sum) + sumLeaf(root.Right, sum)
}

func sumNumbers(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return root.Val
	}

	return sumLeaf(root.Left, root.Val) + sumLeaf(root.Right, root.Val)
}

func sumNumbersTest() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	println(sumNumbers(root)) // 25

	root = &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   0,
			Left:  nil,
			Right: nil,
		},
	}
	println(sumNumbers(root)) // 1026
}

// func main() {
// 	sumNumbersTest()
// }
