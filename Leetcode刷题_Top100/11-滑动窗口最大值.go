package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Print(maxSlidingWindow(nums, k))
}

// 正确但是超出时间限制
func maxSlidingWindow(nums []int, k int) []int {
	nLen := len(nums)
	result := make([]int, 0)
	if nLen <= k {
		sort.Ints(nums)
		result = append(result, nums[nLen-1])
		return result
	}
	for i := 0; i < nLen-k+1; i++ {
		maxNum := getMaxNums(i, nums, k)
		result = append(result, maxNum)
	}

	return result
}

func getMaxNums(index int, nums []int, k int) int {
	currentNum := make([]int, 0)
	for i := index; i < index+k; i++ {
		currentNum = append(currentNum, nums[i])
	}
	sort.Ints(currentNum)
	return currentNum[k-1]
}

//官方：优先队列（大根堆）见题解吧
