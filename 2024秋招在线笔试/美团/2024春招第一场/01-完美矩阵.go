/*
1 0 1
0 1 0
1 0 0
*/

package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	square := make([][]int, n)
	for i := 0; i < n; i++ {
		square[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&square[i][j])
		}
	}

	fmt.Println(countPerfectRectangles(square, n))

}

func countPerfectRectangles(matrix [][]int, n int) []int {
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	// 初始化 dp 数组
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = matrix[i][j] == 1
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] && matrix[i][j] == 1
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] && matrix[i][j] == 1
			} else {
				dp[i][j] = (dp[i-1][j] && dp[i][j-1] && (dp[i-1][j-1] == (matrix[i][j] == 1)))
			}
		}
	}

	// 计算每个 i*i 完美矩形区域的数量
	counts := make([]int, n+1)
	for i := 1; i <= n; i++ {
		perfectCount := 0
		for x := 0; x <= n-i; x++ {
			for y := 0; y <= n-i; y++ {
				if dp[x+i-1][y+i-1] {
					perfectCount++
				}
			}
		}
		counts[i] = perfectCount
	}

	return counts[1:] // 去掉第一个无意义的0
}
