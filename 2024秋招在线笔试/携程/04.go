package main

import (
	"fmt"
	"math"
)

func main() {
	var n, k, sum int
	fmt.Scan(&n, &k, &sum)
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		curNum := 0
		fmt.Scan(&curNum)
		nums = append(nums, curNum)
	}
	result := calculate(nums, n, k, sum)
	fmt.Print(result)

}

func calculate(a []int, n, k, sum int) int {
	windows := 0
	for i := 0; i < k; i++ {
		windows += a[i] // 窗口内的和
	}

	mod := 0
	for i := k; i <= n; i++ {
		if windows > sum {
			cha := windows - sum
			for j := i - 1; j >= i-k && cha > 0; j-- {
				if a[j] > 0 {
					de := int(math.Min(float64(a[j]), float64(cha)))
					a[j] -= de
					cha -= de
					mod += de
				}
			}
			windows = 0
			for j := i - k; j < i; j++ {
				windows += a[j]
			}
		}
		if i < n {
			windows += a[i] - a[i-k]
		}
	}
	return mod
}
