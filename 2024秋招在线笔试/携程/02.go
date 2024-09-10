/*
游游有一个长度为 k且只包含字符 0和 1的字符串，下标为 1~ k。
定义其权值为：每次操作可以选择一个下标 $i,(1\leq i\leq k)$，将[1,i]的字符全部取反（0变1,1变0）。将字符串变为全1的最少操作次数。
例如字符串 00110，第一次操作选择 $i=5$，字符串变成 11001，第二次操作选择 $i=4$，字符串变成 00111，第三次操作选择 $i=2$，字符串变成 11111。至少需要 3次操作，字符串权值为 3。
给出一个长度为n的01字符串，问：有多少个长度为奇数的非空子串的权值是奇数（子串是该字符串中任意连续的字符组成的字符串）？
定义子串为，从原字符串中，连续的选择一段字符（可以全选、可以不选）组成的新字符串。
*/
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
