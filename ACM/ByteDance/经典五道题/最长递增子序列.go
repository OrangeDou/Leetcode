package main

import "fmt"

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxLength := 1
	dp := make([]int, len(nums))
	// 初始化dp数组，否则报错
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		maxLength = max(maxLength, dp[i])
	}
	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	arr := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Print(lengthOfLIS(arr))
}
