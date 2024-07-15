package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// // Solution 结构体包含 Max 变量，用于存储最大直径
// type Solution struct {
//     Max int
// }

// // depth 辅助递归函数，计算以 node 为根的子树的深度，并更新 Max 为最大直径
// func (s *Solution) depth(node *TreeNode) int {
//     if node == nil {
//         return 0 // 空节点返回0
//     }
//     left := s.depth(node.Left)  // 左子树的深度
//     right := s.depth(node.Right) // 右子树的深度
//     s.Max = max(s.Max, left+right) // 更新最大直径
//     return max(left, right) + 1 // 返回以当前节点为根的子树的深度
// }

// // diameterOfBinaryTree 计算二叉树的直径
// func  diameterOfBinaryTree(root *TreeNode) int {
//     s := Solution{}
//     s.Max = 0 // 初始化 Max 为 0
//     s.depth(root) // 调用 depth 函数
//     return s.Max // 返回最大直径
// }

// // max 辅助函数，比较两个整数并返回较大的那个
// func max(a, b int) int {
//     if a > b {
//         return a
//     }
//     return b
// }

func diameterOfBinaryTree(root *TreeNode) (ans int) {
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return -1 // 下面 +1 后，对于叶子节点就刚好是 0
		}
		lLen := dfs(node.Left) + 1  // 左子树最大链长+1
		rLen := dfs(node.Right) + 1 // 右子树最大链长+1
		ans = max(ans, lLen+rLen)   // 两条链拼成路径
		return max(lLen, rLen)      // 当前子树最大链长
	}
	dfs(root)
	return
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
