package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	node2 := TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	node1 := TreeNode{
		Val:   3,
		Left:  &node2,
		Right: nil,
	}
	root := TreeNode{
		Val:   1,
		Left:  nil,
		Right: &node1,
	}
	fmt.Print(child(&root))

}

func child(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := child(root.Left)
	right := child(root.Right)
	fmt.Printf("节点 %v 的左子树有 %d 个节点，右子树有 %d 个节点\n",
		root, left, right)
	return left + right + 1
}
