package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	value := make([]int, 0)
	for head != nil {
		value = append(value, head.Val)
		head = head.Next
	}
	resultNode := ListNode{}
	for i := len(value); i >= 0; i-- {
		resultNode.Val = value[i]
		nextNode := ListNode{}
		resultNode.Next = &nextNode
		resultNode = *resultNode.Next
	}
	return &resultNode
}

func main() {
	node2 := ListNode{Val: 3, Next: nil}
	node1 := ListNode{Val: 2, Next: &node2}
	node := reverseList2(&node1)
	fmt.Print(node)
}

// 大佬的思路1:双指针
func reverseList2(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next

	}
	return pre
}

// 递归法：
func reverseList3(head *ListNode) *ListNode {
	return help(nil, head)
}
func help(pre, head *ListNode) *ListNode {
	if head == nil {
		return pre
	}
	next := head.Next
	head.Next = pre
	return help(head, next)
}
