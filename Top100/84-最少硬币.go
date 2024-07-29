package main

import "math"

func coinChange(coins []int, amount int) int {
	length := len(coins)
	if amount < coins[0] {
		return 0
	}
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt32
	}

	for i := 0; i < length; i++ {
		for j := coins[i]; j <= amount; j++ {
			if dp[j-coins[i]] != math.MaxInt32 {
				dp[j] = min(dp[j], dp[j-coins[i]+1])
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]

}
