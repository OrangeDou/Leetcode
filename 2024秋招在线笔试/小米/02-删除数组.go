package main

import (
	"fmt"
)

func main() {
	var n, x int
	fmt.Scan(&n, &x)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	if n == 1 {
		fmt.Print(1)
		return
	}
	result := help(nums, x)
	fmt.Print(result)

}

func help(nums []int, x int) int {
	n := len(nums)
	sum := 0
	for _, v := range nums {
		sum += v
	}
	r := sum % x
	if r == 0 {
		return 0
	}
	minDelete := n + 1
	for i := 0; i < n; i++ {
		if nums[i]%x <= r {
			minDelete = min(minDelete, i)
		}
	}
	minAdd := n + 1
	for i := 0; i < n; i++ {
		if nums[i]%x >= x-r {
			minAdd = min(minAdd, i)
		}
	}
	return min(minDelete, minAdd)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
