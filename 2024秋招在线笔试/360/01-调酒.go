package main

import (
	"fmt"
)

func main() {

	var n, m int
	fmt.Scan(&n, &m)
	buy := make([][]int, 0)
	for i := 0; i < n; i++ {
		var num1, num2 int
		fmt.Scan(&num1, &num2)
		cur := make([]int, 2)
		cur[0] = num1
		cur[1] = num2
		buy = append(buy, cur)
	}
	//fmt.Print(buy)
	wire := make([]int, m)
	for i := 1; i <= m; i++ {
		wire[i-1] = i
	}

	// 创建一个映射来记录每个数字是否已被拿走
	taken := make(map[int]bool)

	// 记录满足的配对
	satisfiedPairs := make([]int, 0)
	var count int
	// 读取每一行的两个数字
	for i := 0; i < n; i++ {
		num1 := buy[i][0]
		num2 := buy[i][1]
		// 检查两个数字是否都未被拿走
		if !taken[num1] && !taken[num2] {
			// 标记这两个数字为已拿走
			taken[num1] = true
			taken[num2] = true
			// 记录这一对数字
			satisfiedPairs = append(satisfiedPairs, num1)
			satisfiedPairs = append(satisfiedPairs, num2)
			count++
		} else {
			continue
		}
	}

	fmt.Print(len(satisfiedPairs) / 2)
}
