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
//我的思路
// func isSymmetric1(root *TreeNode, result []int) bool {
//     //中序遍历二叉树
//     if root == nil {
//         return true
//     }
//     isSymmetric(root.Left, result)
//     result = append(result, root.Val)
//     isSymmetric(root.Right, result)

//	    length := len(result)
//	    i := 0
//	    j := length -1
//	    for i < j {
//	        if result[i] != result[j] {
//	            return false
//	        }
//	        i++
//	        j--
//	    }
//	    return true
//	}
//
// 递归
func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}

func check(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
}

//迭代

func isSymmetric3(root *TreeNode) bool {
	u, v := root, root
	q := []*TreeNode{}
	q = append(q, u)
	q = append(q, v)
	for len(q) > 0 {
		u, v = q[0], q[1]
		q = q[2:]
		if u == nil && v == nil {
			continue
		}
		if u == nil || v == nil {
			return false
		}
		if u.Val != v.Val {
			return false
		}

		q = append(q, u.Left)
		q = append(q, v.Right)

		q = append(q, u.Right)
		q = append(q, v.Left)

	}
	return true
}
