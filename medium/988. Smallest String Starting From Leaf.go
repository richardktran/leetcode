package main

import "fmt"

func toChar(val int) byte {
	return byte(val + 'a')
}

func reverseString(s string) string {
	result := []byte(s)

	for i := 0; i < len(s)/2; i++ {
		result[i], result[len(s)-1-i] = result[len(s)-1-i], result[i]
	}

	return string(result)
}

func smallest(node *TreeNode, currentString string, result *string) {
	leafString := currentString + string(toChar(node.Val))
	if node.Left == nil && node.Right == nil {
		if result == nil || *result == "" {
			*result = reverseString(leafString)
		} else {
			if reverseString(leafString) < *result {
				*result = reverseString(leafString)
			}
		}
	}

	if node.Left != nil {
		smallest(node.Left, leafString, result)
	}

	if node.Right != nil {
		smallest(node.Right, leafString, result)
	}
}

func smallestFromLeaf(root *TreeNode) string {
	if root == nil {
		return ""
	}

	var result string

	currentString := string(toChar(root.Val))

	smallest(root, currentString, &result)

	return result[:len(result)-1]
}

func smallestFromLeafTest() {
	root := &TreeNode{0, &TreeNode{1, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}}, &TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}}}
	result := smallestFromLeaf(root)
	fmt.Println(result) // "dba"

	fmt.Println("=====================================")
	root = &TreeNode{25, &TreeNode{1, &TreeNode{3, nil, nil}, &TreeNode{1, nil, nil}}, &TreeNode{3, &TreeNode{0, nil, nil}, &TreeNode{2, nil, nil}}}
	result = smallestFromLeaf(root)
	fmt.Println(result) // "adz"

	fmt.Println("=====================================")
	root = &TreeNode{2, &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{1, nil, nil}}, &TreeNode{1, &TreeNode{1, nil, nil}, &TreeNode{2, nil, nil}}}
	result = smallestFromLeaf(root)
	fmt.Println(result) // "abc"
}

// func main() {
// 	smallestFromLeafTest()
// }
