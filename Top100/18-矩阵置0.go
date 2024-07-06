package main

import "fmt"

func main() {
	matrix := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroees(matrix)
	fmt.Print(matrix)
}

// a pices of shit
func setZeroes(matrix [][]int) {
	// 记录0出现的行和列

	lineAndCol := []int{}
	length := len(matrix)
	col := len(matrix[0])
	replace := []int{}
	result := make([][]int, length)
	for i := 0; i < col; i++ {
		replace = append(replace, 0)
	}

	for i := 0; i < length; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] == 0 {
				lineAndCol = append(lineAndCol, i, j)
			}
		}
	}
	// 归零操作
	for k := 0; k < len(lineAndCol); k++ {
		if k%2 == 0 {
			//对行操作
			matrix[k] = replace
		} else {
			for i := 0; i < length; i++ {
				matrix[i][k] = 0
			}
		}
	}
	for l := 0; l < length; l++ {
		result[l] = matrix[l]
	}
}

// 官方题解：
func setZeroees(matrix [][]int) {
	row := make([]bool, len(matrix))
	col := make([]bool, len(matrix[0]))
	for i, r := range matrix {
		for j, v := range r {
			if v == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}
	for i, r := range matrix {
		for j := range r {
			if row[i] || col[j] {
				r[j] = 0
			}
		}
	}
}
