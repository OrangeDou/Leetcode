package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var level int = 1

func PrintLevel(head *TreeNode, level int, res map[int]int) {
	if head == nil {
		return
	}
	res[head.Val] = level
	fmt.Printf("Node %v at level %v\n", head.Val, level)
	PrintLevel(head.Left, level+1, res)
	PrintLevel(head.Right, level+1, res)

	return

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
	mapu := make(map[int]int, 0)
	PrintLevel(&root, 1, mapu)
	fmt.Print(mapu)
}
