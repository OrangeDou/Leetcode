package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: -1, Next: nil}
	p, q := dummy, dummy
	for i := 0; i < n+1; i++ {
		p = p.Next
	}
	if p == nil {
		q.Next = q.Next.Next
		return q.Next
	}
	for p != nil {
		p = p.Next
		q = q.Next
	}
	q.Next = q.Next.Next
	return head
}

func main() {
	node1 := ListNode{Val: 1, Next: nil}
	fmt.Print(removeNthFromEnd(&node1, 1))
}
