package main

func deleteNode(node *ListNode) {
	for node.Next != nil {
		node.Val = node.Next.Val
		if node.Next.Next == nil {
			node.Next = nil
		} else {
			node = node.Next
		}
	}

}

func DeleteNodeTest() {
	list := &ListNode{Val: 4}
	list.Next = &ListNode{Val: 5}
	list.Next.Next = &ListNode{Val: 1}
	list.Next.Next.Next = &ListNode{Val: 9}

	deleteNode(list.Next)
	for list != nil {
		println(list.Val)
		list = list.Next
	}
}

// func main() {
// 	DeleteNodeTest()
// }
