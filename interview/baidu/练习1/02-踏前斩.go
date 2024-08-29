package main

import "fmt"

func minMP(a []int) int {
	n := len(a)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + 1 // 初始状态下，假设只使用强力攻击
		if i >= 3 && i+1 < n && i+2 < n {
			cur1, cur2 := 0, 0
			if a[i+1] >= 2 {
				cur1 = 1
			}
			if a[i+2] >= 3 {
				cur2 = 1
			}
			// 计算使用踏前斩的mp
			mp := 5 + cur1 + cur2
			if a[i-3] > 0 {
				mp += 1 // 因为踏前斩可以对前一个怪物也造成伤害
			}
			dp[i] = min(dp[i], dp[i-3]+mp)
		}
	}
	return dp[n]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	monsters := []int{2, 3, 4, 2, 3} // 示例怪物血量数组
	fmt.Println("Minimum MP required:", minMP(monsters))
}
