package main

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 合并两个有序链表
func mergeTwoLists2(head1, head2 *ListNode) *ListNode {
	dummy := &ListNode{Val: -1, Next: nil}
	result := dummy
	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			dummy.Next = head1
			head1 = head1.Next
			dummy = dummy.Next
		} else {
			dummy.Next = head2
			head2 = head2.Next
			dummy = dummy.Next
		}
	}
	if head1 != nil {
		dummy.Next = head1
	}

	if head2 != nil {
		dummy.Next = head2
	}
	return result
}

// 合并两个有序链表k-1次:如果存在空列表则通不过
func mergeKLists1(lists []*ListNode) *ListNode {
	firstList := lists[0]
	for i := 1; i < len(lists); i++ {
		secondList := lists[i]
		firstList = mergeTwoLists2(firstList, secondList)
	}
	return firstList
}

func main() {
	// 初始化第一个有序链表: 1 -> 3 -> 5
	list1 := &ListNode{Val: 1}
	list1.Next = &ListNode{Val: 3}
	list1.Next.Next = &ListNode{Val: 5}

	// 初始化第二个有序链表: 2 -> 4 -> 6
	list2 := &ListNode{Val: 2}
	list2.Next = &ListNode{Val: 4}
	list2.Next.Next = &ListNode{Val: 6}
	res := make([]*ListNode, 0)
	res = append(res, list1)
	res = append(res, list2)

	mergeKLists1(res)
}

// 优先队列方法-------------------------------
type hp []*ListNode

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(*ListNode)) }
func (h *hp) Pop() any {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}
func mergeKLists2(lists []*ListNode) *ListNode {
	h := hp{}
	for _, head := range lists {
		if head != nil {
			h = append(h, head)
		}
	}
	heap.Init(&h) // 建堆

	dummy := &ListNode{}
	cur := dummy
	for len(h) > 0 {
		node := heap.Pop(&h).(*ListNode)
		if node.Next != nil {
			heap.Push(&h, node.Next)
		}
		cur.Next = node
		cur = cur.Next
	}
	return dummy.Next
}
