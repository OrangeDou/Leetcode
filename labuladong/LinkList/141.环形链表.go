/*
 * @lc app=leetcode.cn id=141 lang=golang
 *
 * [141] 环形链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	hash := make(map[*ListNode]int, 0)
	for head != nil {
		if _, exist := hash[head]; exist {
			return true
		}
		hash[head] = 1
		head = head.Next
	}
	return false
}

// @lc code=end

