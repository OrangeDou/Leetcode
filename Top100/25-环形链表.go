package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node2 := ListNode{Val: 3}
	node1 := ListNode{Val: 2, Next: &node2}
	node := ListNode{Val: 1, Next: &node1}
	node2.Next = &node

	boolre := hasCycle(&node)
	fmt.Print(boolre)
}

func hasCycle(head *ListNode) bool {
	hashMap := make(map[ListNode]int, 0)
	for head != nil {
		if _, ok := hashMap[*head]; ok {
			return true
		}
		hashMap[*head] = head.Val
		head = head.Next
	}
	return false

}
