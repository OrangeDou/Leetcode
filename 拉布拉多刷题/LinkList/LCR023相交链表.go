package main

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 如果不相交，p,q会最终同时走到nil
	p, q := headA, headB
	for p != q {
		if p == nil {
			p = headB
		} else {
			p = p.Next
		}
		if q == nil {
			q = headA
		} else {
			q = q.Next
		}
	}
	return p
}
