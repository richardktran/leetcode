package main

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func distributeCoins(root *TreeNode) int {
	var result int
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		result += abs(left) + abs(right)
		return node.Val + left + right - 1
	}

	dfs(root)
	return result
}

func distributeCoinsTest() {
	root := &TreeNode{3, &TreeNode{0, nil, nil}, &TreeNode{0, nil, nil}}
	result := distributeCoins(root)
	println(result)
}

// func main() {
// 	distributeCoinsTest()
// }
