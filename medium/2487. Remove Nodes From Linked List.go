package main

import "fmt"

func removeNodes(head *ListNode) *ListNode {
	stack := []*ListNode{}

	for head != nil {
		for len(stack) > 0 && stack[len(stack)-1].Val < head.Val {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			stack = append(stack, head)
			head = head.Next
			continue
		}

		stack[len(stack)-1].Next = head
		stack = append(stack, head)
		head = head.Next
	}

	if len(stack) > 0 {
		stack[len(stack)-1].Next = nil
	}

	return stack[0]
}

func removeNodesTest() {
	node := &ListNode{Val: 5}
	node.Next = &ListNode{Val: 2}
	node.Next.Next = &ListNode{Val: 13}
	node.Next.Next.Next = &ListNode{Val: 3}
	node.Next.Next.Next.Next = &ListNode{Val: 8}

	resultNode := removeNodes(node)

	for resultNode != nil {
		fmt.Println(resultNode.Val)
		resultNode = resultNode.Next
	}
}

// func main() {
// 	removeNodesTest()
// }
