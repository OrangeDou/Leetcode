package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node2 := ListNode{Val: 3, Next: nil}
	node1 := ListNode{Val: 2, Next: &node2}
	node := ListNode{Val: 1, Next: &node1}

	boolre := isPalindrome(&node)
	fmt.Print(boolre)
	//print_values_in_reverse(&node)
}

func isPalindrome(head *ListNode) bool {
	valSet := make([]int, 0)
	for head != nil {
		valSet = append(valSet, head.Val)
		head = head.Next
	}
	length := len(valSet)
	if length == 1 {
		return true
	}
	if length == 0 {
		return false
	}
	i := 0
	j := length - 1
	for i < j {
		if valSet[i] != valSet[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func print_values_in_reverse(head *ListNode) {
	if head != nil {
		print_values_in_reverse(head.Next)
	}
	fmt.Print(head.Val)
}
