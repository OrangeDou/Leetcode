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
	fmt.Print(subarraySum2(nums, -1))
}

// 其他快速解法: 前缀和+哈希表
func subarraySum2(nums []int, k int) int {
	n := len(nums)
	mp := map[int]int{0: 1}
	pre, ans := 0, 0

	for i := 0; i < n; i++ {
		pre += nums[i]
		if c, ok := mp[pre-k]; ok { //从数组的开始到当前索引的累积和是 pre。
			//从数组的开始到某个之前的索引（假设是 i-1）的累积和是 pre-k。
			//因此，从索引 i-1 到当前索引 i 的元素之和就是 k。换句话说，数组中从 i-1 到 i 的子数组的和等于 k。
			ans += c
		}
		mp[pre]++
	}
	return ans
}

/*
前缀和加哈希表的解法通常用于解决涉及累积和、频率统计或子数组/子序列问题的场景。这种解法特别适用于以下类型的算法题：
子数组和问题：当需要找出数组中和等于特定值的子数组数量时，例如上述的 subarraySum 函数。
区间求和：快速计算数组中任意两个索引之间的元素和，例如，找出数组中从索引 i 到 j 的元素和。
滑动窗口问题：当问题涉及到在滑动窗口中维持特定条件的统计，如最大和、最小值或其他累积量。
频率统计：需要统计数组中元素出现次数的场景，尤其是在需要快速查找元素出现次数时。
前缀数量：计算小于等于某个值的元素数量，或在数组中找到第 k 大（或小）的元素。
子序列问题：当问题涉及到找出满足特定条件的连续子序列，如连续子序列的最大和。
区间更新问题：在数组中快速更新某个区间的值，并能够快速计算更新后的累积量。
分组问题：将数组分成若干组，每组的累积和满足特定条件。
模式匹配：在大数据集中查找特定模式或子字符串的出现次数。
优化搜索效率：在搜索问题中，通过前缀和快速排除不可能的选项，结合哈希表来存储中间状态。
*/
