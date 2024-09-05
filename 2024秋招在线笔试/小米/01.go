package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, 0)
	b := make([]int, 0)
	for i := 0; i < n; i++ {
		curNum := 0
		fmt.Scan(&curNum)
		a = append(a, curNum)
	}
	for i := 0; i < n; i++ {
		curNum := 0
		fmt.Scan(&curNum)
		b = append(b, curNum)
	}
	if n < 1 || n > 100000 {
		return
	}
	if n == 1 {
		if a[0] == b[0] {
			fmt.Print(a[0] + b[0])
			return
		} else {
			fmt.Print(max(a[0], b[0]))
			return
		}
	}
	result := help(a, b)
	fmt.Print(result)
}
func help(a, b []int) int {
	n := len(a)
	if n == 0 {
		return 0
	}
	minTime := math.MaxInt32
	for i := 0; i < n; i++ {
		minTime = min(minTime, a[i]+b[i])
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j {
				minTime = min(minTime, max(a[i], b[j]))
			}
		}
	}
	return minTime
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
