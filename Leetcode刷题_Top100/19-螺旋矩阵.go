package main

import "fmt"

func main() {
	matrix := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	fmt.Print(spiralOrder2(matrix))
}

func spiralOrder(matrix [][]int) []int {
	ans := make([]int, 0)
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return ans
	}

	for {
		var (
			row       = 0
			column    = 0
			rowEnd    = len(matrix) - 1
			columnEnd = len(matrix[0]) - 1
		)
		//左 -> 右
		for i := column; i <= columnEnd; i++ {
			ans = append(ans, matrix[row][i])
		}
		row++
		if row > rowEnd {
			break
		}

		//上 -> 下
		for j := row; j < rowEnd; j++ {
			ans = append(ans, matrix[j][columnEnd])
		}
		columnEnd--
		if columnEnd < column {
			break
		}

		//右 -> 左
		for k := columnEnd; k >= column; k-- {
			ans = append(ans, matrix[rowEnd][k])
		}
		rowEnd--
		if rowEnd < row {
			break
		}

		//下 -> 上
		for o := rowEnd; o >= row; o-- {
			ans = append(ans, matrix[o][column])
		}
		column++
		if column > columnEnd {
			break
		}
	}
	return ans
}

func spiralOrder2(matrix [][]int) []int {
	ans := make([]int, 0)
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return ans
	}

	var (
		u = 0                  // 上边界
		d = len(matrix) - 1    // 下边界
		l = 0                  // 左边界
		r = len(matrix[0]) - 1 // 右边界
	)

	for u <= d && l <= r { // 确保至少有一行和一列可以遍历
		// 左 -> 右
		for i := l; i <= r; i++ {
			ans = append(ans, matrix[u][i])
		}
		u++

		// 上 -> 下
		for i := u; i <= d; i++ {
			ans = append(ans, matrix[i][r])
		}
		r--

		if u <= d { // 检查是否还有行可以遍历
			// 右 -> 左
			for i := r; i >= l; i-- {
				ans = append(ans, matrix[d][i])
			}
			d--
		}

		if l <= r { // 检查是否还有列可以遍历
			// 下 -> 上
			for i := d; i >= u; i-- {
				ans = append(ans, matrix[i][l])
			}
			l++
		}
	}
	return ans
}
