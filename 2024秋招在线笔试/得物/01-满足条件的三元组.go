/*
给出两个非负整数n和m，请你在[0,n]区间内找到三个数x,y,z，满足它们的和等于m。一共有多少种合法的(x,y,z)三元组？
*/
package main

import "fmt"

func main() {
	var n, m, count int
	fmt.Scan(&n, &m)
	nums := make([]int, n+1)
	for i := 0; i <= n; i++ {
		nums[i] = i
	}
	for i := 0; i <= n; i++ {
		for j := 0; j <= n; j++ {
			curNum2 := m - i - j
			if curNum2 >= 0 && curNum2 <= n {
				count++
			}
			if i+j >= m {
				break
			}
		}
	}
	fmt.Print(count)
}
