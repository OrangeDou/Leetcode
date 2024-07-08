package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	node := ListNode{Val: 2, Next: nil}
	boolre := isPalindrome(&node)
	fmt.Print(boolre)
}

func isPalindrome(head *ListNode) bool {
	valSet := make([]int, 0)
	for head != nil {
		valSet = append(valSet, head.Val)
		head = head.Next
	}
	length := len(valSet)
	if length%2 != 0 || length == 0 {
		return false
	}
	if length == 1 {
		return true
	}
	i := 0
	j := length - 1
	for i < j {
		if valSet[i] != valSet[j] {
			return false
			break
		}
		i++
		j--
	}
	return true
}
