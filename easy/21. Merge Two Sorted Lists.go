package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Normal solution
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	current := &ListNode{
		Val:  0,
		Next: nil,
	}

	results := current

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	} else if list2 != nil {
		current.Next = list2
	}

	return results.Next
}

// Recursive solution
func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists2(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists2(list1, list2.Next)
		return list2
	}
}

func MergeTwoSortedListsTest() {
	l1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	l2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	fmt.Println(mergeTwoLists(l1, l2))
	fmt.Println(mergeTwoLists2(l1, l2))
}

func main() {
	MergeTwoSortedListsTest()
}
