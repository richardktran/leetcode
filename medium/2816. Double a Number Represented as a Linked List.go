package main

import "fmt"

func doubleIt(head *ListNode) *ListNode {
	result := head

	for head != nil {
		head.Val *= 2
		head = head.Next
	}
	head = result

	if head.Val >= 10 {
		newHead := &ListNode{Val: 1, Next: head}
		head.Val -= 10
		head = newHead
		result = head
	}

	for head.Next != nil {
		if head.Next.Val >= 10 {
			head.Next.Val -= 10
			head.Val += 1
		}
		head = head.Next
	}

	return result
}

func doubleItTest() {
	node := &ListNode{Val: 9}
	node.Next = &ListNode{Val: 9}
	node.Next.Next = &ListNode{Val: 9}

	resultNode := doubleIt(node)

	for resultNode != nil {
		fmt.Println(resultNode.Val)
		resultNode = resultNode.Next
	}
}

// func main() {
// 	doubleItTest()
// }
