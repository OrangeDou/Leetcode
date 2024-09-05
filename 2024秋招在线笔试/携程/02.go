package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var n int
	fmt.Scan(&n)
	var s string
	fmt.Scan(&s)
	res := calculate(s, n)
	fmt.Print(res * 2)
}

func findWeight(s string, k int) int {
	dp := make([]int, k+1)
	dp[0] = 0
	for i := 1; i <= k; i++ {
		if s[i-1] == '1' {
			dp[i] = dp[i-1]
		} else {
			dp[i] = min(dp[i-1], dp[0]+1)
		}
	}

	return dp[k]
}

func calculate(s string, n int) int {
	if len(s) == 0 {
		return 0
	}
	count := 0
	// 从1开始到小于n的奇数
	for i := 1; i <= n; i += 2 {
		for j := 0; j < len(s)-i; j++ {
			curStr := s[j : j+i]
			if findWeight(curStr, i)%2 != 0 {
				count++
			}
		}
	}
	return count
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func randu(n int) int {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(n+1) + 1
	return randNum
}
