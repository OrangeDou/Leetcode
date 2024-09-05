package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 反转链表2
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// 把待反转链表部分断开单独拿出来
	forward, start, behind := head, head, head
	dummy := &ListNode{Next: head}
	for i := 1; forward != nil && i < left-1; i++ {
		forward = forward.Next
	}

	start = forward.Next

	for i := 1; start != nil && i < right; i++ {
		behind = behind.Next
	}
	end := behind
	behind = behind.Next
	// [head..][start:end][behind]
	forward.Next = nil
	end.Next = nil
	// 进行反转

	start = reverseList(start)
	// 拼接新的链表
	dummy.Next = head
	for head.Next != nil {
		head = head.Next
	}
	head.Next = start
	for head.Next != nil {
		head = head.Next
	}
	head.Next = behind
	return dummy.Next
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func main() {
	node1 := ListNode{Val: 5, Next: nil}
	head := ListNode{Val: 3, Next: &node1}
	fmt.Print(reverseBetween(&head, 1, 2))
}

// 反转链表1
func reverseList1(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
