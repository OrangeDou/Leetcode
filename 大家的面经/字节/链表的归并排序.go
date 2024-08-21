package main

import "fmt"

// 数组归并排序的写法
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

func main() {
	// arr := []int{38, 27, 43, 3, 9, 82, 10}
	// fmt.Println("Given array is:", arr)

	// sortedArr := mergeSort(arr)
	// fmt.Println("Sorted array is: ", sortedArr)
	node3 := ListNode{Val: 1, Next: nil}
	node2 := ListNode{Val: 2, Next: &node3}
	node1 := ListNode{Val: 3, Next: &node2}
	node := ListNode{Val: 4, Next: &node1}
	listHead := mergeSort2(&node)
	for listHead != nil {
		fmt.Print(listHead.Val)
		listHead = listHead.Next
	}

}

// 链表的归并排序

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeLinkList(headA, headB *ListNode) *ListNode {
	dummy := &ListNode{}
	result := dummy
	i, j := headA, headB
	for i != nil && j != nil {
		if i.Val < j.Val {
			dummy.Next = i
			i = i.Next
			dummy = dummy.Next
		} else {
			dummy.Next = j
			j = j.Next
			dummy = dummy.Next
		}
	}
	if i != nil {
		dummy.Next = i
	}
	if j != nil {
		dummy.Next = j
	}
	return result.Next
}

func mergeSort2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := findMiddle(head)
	rightList := mid.Next
	mid.Next = nil
	// 递归左右两边
	left := mergeSort2(head)
	right := mergeSort2(rightList)

	return mergeLinkList(left, right)

}

// 寻找链表的中间节点
func findMiddle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
