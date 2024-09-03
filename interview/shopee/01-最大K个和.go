package main

import (
	"fmt"
	"sort"
)

func main() {
	test := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Print(maxK(test, 2))
}

func maxK(nums []int, k int) int {
	// write code here
	sort.Ints(nums)
	count := 0
	if k > len(nums) {
		return 0
	}
	for i := len(nums) - 1; i >= len(nums)-k; i-- {
		count += nums[i]
	}
	return count
}
