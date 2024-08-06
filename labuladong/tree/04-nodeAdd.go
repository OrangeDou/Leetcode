package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Add(root *TreeNode) {
	if root == nil {
		return
	}
	root.Val++
	Add(root.Left)
	Add(root.Right)
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
	Add(&root)
	fmt.Print(root)
}
