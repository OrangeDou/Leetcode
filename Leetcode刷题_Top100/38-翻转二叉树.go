package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}

// 深度优先搜索

func main() {

	node6 := TreeNode{
		Val:   9,
		Left:  nil,
		Right: nil,
	}
	node5 := TreeNode{
		Val:   6,
		Left:  nil,
		Right: nil,
	}
	node4 := TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	node3 := TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	node2 := TreeNode{
		Val:   2,
		Left:  &node3,
		Right: &node4,
	}
	node1 := TreeNode{
		Val:   7,
		Left:  &node5,
		Right: &node6,
	}
	root := TreeNode{
		Val:   1,
		Left:  &node2,
		Right: &node1,
	}
	invertTree(&root)
	fmt.Print(&root)
}
