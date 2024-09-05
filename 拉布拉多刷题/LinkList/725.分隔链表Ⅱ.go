package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	var length int
	lengthArr := make([]int, 0)
	result := make([]*ListNode, 0)
	head2 := head
	for head != nil {
		length++
		head = head.Next
	}
	p := length / k //p=3
	q := length % k //q=1
	if p > 0 {
		for i := 0; i < p; i++ {
			lengthArr = append(lengthArr, p) //[3,3,3]
		}
		for j := 0; j < q; j++ {
			lengthArr[j]++ //[4,3,3]
		}
		for k := 0; k < len(lengthArr); k++ {
			curLength := lengthArr[k] //4
			curList := &ListNode{}
			curHead := curList
			// result = append(result, newList)
			for o := 0; o < curLength; o++ {
				curList.Next = head2
				head2 = head2.Next
				curList = curList.Next
			}
			curList.Next = nil
			result = append(result, curHead.Next)
		}
	} else {
		for r := 0; r < k; r++ {
			if head2 != nil {
				head2Node := &ListNode{Next: head2}
				head2 = head2.Next
				head2Node.Next.Next = nil
				result = append(result, head2Node.Next)

			} else {
				result = append(result, nil)
			}
		}
	}
	return result
}
func main() {
	// 初始化链表：1 -> 2 -> 3 -> ... -> 10
	var head *ListNode = &ListNode{Val: 1} // 创建第一个节点并初始化头指针
	current := head                        // 初始化当前节点为头节点

	// 循环创建剩余的节点并连接到链表
	for i := 2; i <= 3; i++ {
		current.Next = &ListNode{Val: i} // 为当前节点设置下一个节点
		current = current.Next           // 移动到下一个节点
	}

	fmt.Print(splitListToParts(head, 5))

}

//-------------------------------------
//官解
func splitListToParts(head *ListNode, k int) []*ListNode {
	n := 0
	for node := head; node != nil; node = node.Next {
		n++
	}
	quotient, remainder := n/k, n%k
	parts := make([]*ListNode, k)
	for i, curr := 0, head; i < k && curr != nil; i++ {
		parts[i] = curr
		partSize := quotient
		if i < remainder {
			partSize++
		}
		for j := 1; j < partSize; j++ {
			curr = curr.Next
		}
		curr, curr.Next = curr.Next, nil
	}
	return parts

}
