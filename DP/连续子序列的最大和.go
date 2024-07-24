/*
输入：[3,-4,2,-1,2,6,-5,4]
输出：9
解释：[2,-1,2,6]
*/

package main

import "fmt"

/* 暴力
找到每个元素开头的子序列的最大值，保存起来
在保存起来的map中选最大的
*/

func main() {
	nums := []int{3, -4, 2, -1, 2, 6, -5, 4}
	fmt.Print(findCountinusSeq1(nums))
}

// 暴力
func findCountinusSeq1(nums []int) int {
	resultMap := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		curSum := nums[i]
		resultMap[i] = curSum
		for j := i + 1; j < len(nums); j++ {
			curSum += nums[j]
			resultMap[i] = max(resultMap[i], curSum)
		}
	}
	return findMaxInMap(resultMap)
}

func findMaxInMap(mp map[int]int) int {
	max := mp[0]
	for _, v := range mp {
		if v > max {
			max = v
		}
	}
	return max
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// 优化2 ---------------
func findCountinusSeq3(nums []int) int {
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
	}
	max := nums[0]
	for j := 1; j < len(nums); j++ {
		if nums[j] > max {
			max = nums[j]
		}
	}
	return max
}
