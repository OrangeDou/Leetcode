package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node3 := ListNode{Val: 3, Next: nil}
	node2 := ListNode{Val: 4, Next: &node3}
	head1 := ListNode{Val: 2, Next: &node2}

	node5 := ListNode{Val: 4, Next: nil}
	node4 := ListNode{Val: 6, Next: &node5}
	head2 := ListNode{Val: 5, Next: &node4}

	fmt.Print((addTwoNumbers(&head1, &head2)))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	list1 := make([]int, 0)
	list2 := make([]int, 0)
	for l1 != nil {
		list1 = append(list1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		list2 = append(list2, l2.Val)
		l2 = l2.Next
	}
	num1 := arrayToNum(list1)
	num2 := arrayToNum(list2)
	result := num1 + num2
	listnode := numToLinkList(result)
	return listnode

}

func arrayToNum(arr []int) int {
	k := 1
	result := 0
	length := len(arr)
	for i := length - 1; i >= 0; i-- {
		result += arr[i] * k
		k *= 10
	}
	return result
}

func numToLinkList(a int) *ListNode {
	strNum := strconv.Itoa(a)
	head := &ListNode{}
	result := head
	for i := len(strNum) - 1; i >= 0; i-- {
		digit, _ := strconv.Atoi(string(strNum[i]))
		result.Next = &ListNode{Val: digit}
		result = result.Next
	}

	result.Next = nil
	return head.Next
}
