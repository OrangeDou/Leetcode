package main

func main() {

}

func detectCycle(head *ListNode) *ListNode {
    hashMap := make(map[ListNode]int, 0)
	for head != nil {
		if _, ok := hashMap[*head]; ok {
			return head
		}
		hashMap[*head] = head.Val
		head = head.Next
	}
	return nil

}