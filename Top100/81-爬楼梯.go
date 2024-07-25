package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入阶梯数：")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input) // 去除输入字符串的前后空格
	n, err := strconv.Atoi(input)    // 将字符串转换为整数
	if err != nil {
		fmt.Println("输入错误，请输入一个整数。")
		return
	}
	fmt.Printf("爬到第%d阶有%d种方法", n, climbStairs2(n))
}

// 动态规划
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 2
	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n-1]
}
