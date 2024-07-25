package main

func generate(numRows int) [][]int {
	dp := make([][]int, numRows)
	for i := 2; i < numRows; i++ {
		for j := 1; j < i; j++ {
			dp[i][j] = dp[i-1][j] + dp[i-1][j-1]
		}
	}
	return dp
}
