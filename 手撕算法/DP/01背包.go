package main

import "fmt"

func main() {
	var w, n int
	fmt.Println("请输入总重量m，以及个数n")
	fmt.Scanf("%d %d", &w, &n)
	var weight []int
	var value []int
	fmt.Println("请输入n个物品各自的重量：")
	for i := 0; i < n; i++ {
		curW := 0
		fmt.Scanf("%d", &curW)
		weight = append(weight, curW)
	}
	fmt.Println("请输入n个物品各自的价值：")
	for j := 0; j < n; j++ {
		curV := 0
		fmt.Scanf("%d", &curV)
		value = append(value, curV)
	}
	fmt.Println(knapsack(w, n, weight, value))

}

func knapsack(W, N int, wt, val []int) int {
	// base case 已初始化
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, W+1)
	}
	for i := 1; i <= N; i++ {
		for w := 1; w <= W; w++ {
			if w-wt[i-1] < 0 {
				// 这种情况下只能选择不装入背包
				dp[i][w] = dp[i-1][w]
			} else {
				// 装入或者不装入背包，择优
				dp[i][w] = max(
					dp[i-1][w-wt[i-1]]+val[i-1],
					dp[i-1][w],
				)
			}
		}
	}

	return dp[N][W]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
