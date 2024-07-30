/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] 分隔链表
 */

// @lc code=start

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// error
// func partition(head *ListNode, x int) *ListNode {
// 	if head == nil {
// 		return nil
// 	}
// 	if head.Next == nil {
// 		return head
// 	}
// 	dummy := &ListNode{-1, nil}
// 	result := dummy
// 	newList := make([]*ListNode, 0)
// 	upperList := make([]*ListNode, 0)

// 	for head != nil {
// 		if head.Val < x {
// 			newList = append(newList, head)
// 		}
// 		if head.Val >= x {
// 			upperList = append(upperList, head)
// 		}
// 		head = head.Next
// 	}

// 	for i := 0; i < len(newList)+len(upperList); i++ {
// 		if i < len(newList) {
// 			dummy.Next = newList[i]
// 		} else {
// 			dummy.Next = upperList[i-3]
// 		}
// 		dummy = dummy.Next
// 	}
// 	return result
// }

func main() {
	node2 := ListNode{Val: 1, Next: nil}
	node1 := ListNode{Val: 2, Next: &node2}
	fmt.Print(partition(&node1, 2))
}

func partition(head *ListNode, x int) *ListNode {
	dummy1 := &ListNode{-1, nil}
	dummy2 := &ListNode{-1, nil}

	p1, p2 := dummy1, dummy2
	p := head
	for p != nil {
		if p.Val >= x {
			p2.Next = p
			p2 = p2.Next
		} else {
			p1.Next = p
			p1 = p1.Next
		}
		temp := p.Next
		p.Next = nil
		p = temp
	}
	p1.Next = dummy2.Next
	return dummy1.Next
}

// @lc code=end
