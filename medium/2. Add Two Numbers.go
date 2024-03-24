package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := &ListNode{}
	result := sum // save the head of the sum
	remain := 0
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			sum.Val = l1.Val + l2.Val + remain

			l1 = l1.Next
			l2 = l2.Next
		} else if l1 == nil {
			sum.Val = l2.Val + remain
			l2 = l2.Next
		} else if l2 == nil {
			sum.Val = l1.Val + remain
			l1 = l1.Next
		}

		if sum.Val >= 10 {
			sum.Val -= 10
			remain = 1
		} else {
			remain = 0
		}

		if l1 == nil && l2 == nil {
			break
		}

		sum.Next = &ListNode{
			Val:  -1,
			Next: nil,
		}
		sum = sum.Next
	}

	if remain == 1 {
		sum.Next = &ListNode{Val: 1}
	}

	return result
}

func addTwoNumbersTest() {
	l1 := &ListNode{Val: 9}
	l1.Next = &ListNode{Val: 9}
	l1.Next.Next = &ListNode{Val: 9}
	l1.Next.Next.Next = &ListNode{Val: 9}

	l2 := &ListNode{Val: 9}
	l2.Next = &ListNode{Val: 9}
	l2.Next.Next = &ListNode{Val: 9}

	result := addTwoNumbers(l1, l2)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}

// func main() {
// 	addTwoNumbersTest()
// }
