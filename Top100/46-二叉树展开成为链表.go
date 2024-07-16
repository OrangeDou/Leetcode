package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 构建二叉树
func flatten(root *TreeNode) {
	list := buildTree(root)
	for i := 1; i < len(list); i++ {
		prev, curr := list[i-1], list[i]
		prev.Left, prev.Right = nil, curr
	}
}

// 前序遍历
func buildTree(root *TreeNode) []*TreeNode {
	list := []*TreeNode{}
	if root != nil {
		list = append(list, root)
		list = append(list, buildTree(root.Left)...)
		list = append(list, buildTree(root.Right)...)
	}
	return list
}

// 前序遍历
func preorderTraversal(root *TreeNode) (ans []int) {
	if root != nil {
		ans = append(ans, root.Val)
		ans = append(ans, preorderTraversal(root.Left)...)
		ans = append(ans, preorderTraversal(root.Right)...)
	}
	return ans
}

// 中序遍历
func inorderTraversal(root *TreeNode) (ans []int) {
	if root == nil {
		return nil
	}
	if root.Left != nil {
		ans = append(ans, inorderTraversal(root.Left)...)
	}
	ans = append(ans, root.Val)
	if root.Right != nil {
		ans = append(ans, inorderTraversal(root.Right)...)
	}
	return ans
}

// 后序遍历
func postorderTraversal(root *TreeNode) (ans []int) {
	if root == nil {
		return nil
	}
	if root.Left != nil {
		ans = append(ans, postorderTraversal(root.Left)...)
	}

	if root.Right != nil {
		ans = append(ans, postorderTraversal(root.Right)...)
	}
	ans = append(ans, root.Val)
	return ans
}
