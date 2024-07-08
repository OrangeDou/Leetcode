package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := new(ListNode)
	result := head
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			head.Next = list1
			list1 = list1.Next
			head = head.Next
		} else {
			head.Next = list2
			list2 = list2.Next
			head = head.Next
		}

	}
	if list1 == nil && list2 != nil {
		for list2 != nil {
			head.Next = list2
			list2 = list2.Next
			head = head.Next
		}
	}
	if list2 == nil && list1 != nil {
		for list1 != nil {
			head.Next = list1
			list1 = list1.Next
			head = head.Next
		}
	}
	return result.Next
}
