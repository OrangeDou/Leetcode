package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&b[i])
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
