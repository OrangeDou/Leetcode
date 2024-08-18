package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	q := []*TreeNode{root}        //广度优先搜索队列
	for i := 0; len(q) > 0; i++ { //轮训每一层，i为层数
		ret = append(ret, []int{}) //存放每个level的结果
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ { //轮寻当前层的每一个节点
			node := q[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p //更新广度优先搜索队列
	}
	return ret
}

func isSymmetric4(root *TreeNode) [][]*TreeNode {
	levelSet := make([][]*TreeNode, 0)
	q := []*TreeNode{root}
	for len(q) > 0 {
		node := q[0]
		levelSet = append(levelSet, q)
		q = q[1:]
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
	}
	return levelSet
}

func main() {
	node3 := TreeNode{
		Val:   6,
		Left:  nil,
		Right: nil,
	}
	node2 := TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	node1 := TreeNode{
		Val:   3,
		Left:  &node2,
		Right: &node3,
	}
	root := TreeNode{
		Val:   1,
		Left:  nil,
		Right: &node1,
	}
	fmt.Print(isSymmetric4(&root))
}
