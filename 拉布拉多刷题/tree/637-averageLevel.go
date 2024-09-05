package main

func averageOfLevels(root *TreeNode) []float64 {
	res := make([][]float64, 0)
	if root == nil {
		return nil
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		res = append(res, []float64{})
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			res[i] = append(res[i], float64(node.Val))
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	result := make([]float64, 0)
	for k := 0; k < len(res); k++ {
		sum := 0.0
		for m := 0; m < len(res[k]); m++ {
			sum += res[k][m]
		}
		result = append(result, sum/float64(len(res[k])))
	}
	return result

}

// athoner manner
func averageOfLevels2(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	q := []*TreeNode{root}

	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			curNode := q[0]
			q = q[1:]
			if curNode.Left != nil {
				q = append(q, curNode.Left)
			}
			if curNode.Right != nil {
				q = append(q, curNode.Right)
			}
		}
	}

}
