package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	listLength := 0
	result := head
	for head != nil {
		listLength += 1
		head = head.Next
	}
	if listLength == 1 {
		return nil
	}
	target := listLength - n
	if target == 0 {
		return result.Next
	}
	headNum := 0
	head = result
	for head != nil {
		if headNum == target-1 {
			head.Next = head.Next.Next
			return result
		}
		headNum += 1
		head = head.Next
	}
	return nil
}
