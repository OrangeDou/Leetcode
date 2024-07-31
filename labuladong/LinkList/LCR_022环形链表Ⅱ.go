package main

func detectCycle(head *ListNode) *ListNode {
	hash := make(map[*ListNode]int, 0)
	for head != nil {
		if hash[head] == 1 {
			return head
		}
		hash[head] = 1
		head = head.Next
	}
	return nil
}
