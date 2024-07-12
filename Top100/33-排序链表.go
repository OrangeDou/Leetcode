package main

import "sort"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	resultArr := make([]int, 0)
	for head != nil {
		resultArr = append(resultArr, head.Val)
		head = head.Next
	}
	sort.Ints(resultArr)
	dummyHead := &ListNode{Next: nil}
	current := dummyHead
	for _, values := range resultArr {
		//创建新节点
		newNode := &ListNode{Val: values}
		current.Next = newNode
		current = newNode
	}
	return dummyHead.Next

}
