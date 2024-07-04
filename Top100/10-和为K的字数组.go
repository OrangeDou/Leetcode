package main

import "fmt"

func subarraySum(nums []int, k int) int {
	nLen := len(nums)
	result := 0
	for i := 0; i < nLen; i++ {
		sum := 0
		for win := i; win < nLen; win++ {
			sum += nums[win]
			if sum == k {
				result += 1
			}
		}
	}
	return result
}

func main() {
	nums := []int{1, -1, 0}
	fmt.Print(subarraySum(nums, 0))
}

// 其他快速解法
func subarraySum2(nums []int, k int) int {
	n := len(nums)
	mp := map[int]int{0: 1}
	pre, ans := 0, 0

	for i := 0; i < n; i++ {
		pre += nums[i]
		if c, ok := mp[pre-k]; ok {
			ans += c
		}
		mp[pre]++
	}
	return ans
}
