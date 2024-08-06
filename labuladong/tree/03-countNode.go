package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNode(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftNode := countNode(root.Left)
	rightNode := countNode(root.Right)
	return leftNode + rightNode + 1
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
	traverse(&root)

}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("从节点 %v 进入节点 %v\n", root, root.Left)
	traverse(root.Left)
	fmt.Printf("从节点 %v 回到节点 %v\n", root.Left, root)

	fmt.Printf("从节点 %v 进入节点 %v\n", root, root.Right)
	traverse(root.Right)
	fmt.Printf("从节点 %v 回到节点 %v\n", root.Right, root)
}
