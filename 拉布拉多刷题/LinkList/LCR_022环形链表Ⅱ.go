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

func detectCycle2(head *ListNode) *ListNode {
	var fast, slow *ListNode
	fast, slow = head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}

	if fast == nil || fast.Next == nil {
		return nil
	}

	// 重新跑, 会在环的入口相遇
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
